package main

import (
	"ceno/common/crypto"
	"ceno/common/rand"
	"ceno/network/protocol"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func send(conn net.Conn) {

	for i := 0; i < 100; i++ {
		session := GetSession()
		randstr := rand.MakeRand(16, rand.RAND_KIND_ALL)
		ss := crypto.MD5(string(randstr))
		words := "{\"ID\":" + strconv.Itoa(i) + ",\"Session\":\"" + session + string(rand.MakeRand(16, rand.RAND_KIND_NUM)) + "\",\"test\":\"golang\",\"Hash\":" + ss + "}"
		conn.Write(protocol.Enpacket([]byte(words)))
	}
	fmt.Println("send over")
	defer conn.Close()
}

func GetSession() string {
	gs1 := time.Now().Unix()
	gs2 := strconv.FormatInt(gs1, 10)
	return gs2
}

func main() {
	server := "localhost:6060"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")
	send(conn)

}
