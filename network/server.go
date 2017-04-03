package main

import (
	"ceno/common"
	"ceno/network/protocol"
	"net"
)

func main() {
	netListen, err := net.Listen("tcp", "localhost:6060")
	common.CheckError(err)

	defer netListen.Close()

	common.Log("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		//timeouSec :=10
		//conn.
		common.Log(conn.RemoteAddr().String(), " tcp connect success")
		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {

	// 缓冲区，存储被截断的数据
	tmpBuffer := make([]byte, 0)

	//接收解包
	readerChannel := make(chan []byte, 16)
	go reader(readerChannel)

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			common.Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}

		tmpBuffer = protocol.Depack(append(tmpBuffer, buffer[:n]...), readerChannel)
	}
	defer conn.Close()
}

func reader(readerChannel chan []byte) {
	for {
		select {
		case data := <-readerChannel:
			common.Log(string(data))
		}
	}
}
