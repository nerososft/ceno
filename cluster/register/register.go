package register

import (
	"ceno/cluster/types"
	"ceno/common"
	"fmt"
	"net"
)

/*
Register 注册中心
*/
type Register struct {
	Rid       string
	IP        string
	Port      int16
	Timestamp string
}

func NewRegister(rid string, ip string, port int16, timestamp string) *Register {
	return &Register{
		Rid:       rid,
		IP:        ip,
		Port:      port,
		Timestamp: timestamp,
	}
}

/*
LoadRegisterList 读取注册中心节点列表
*/
func (reg *Register) LoadRegisterList() {

}

/*
MemSaveNodeToNodeList 节点信息存入内存
*/
func (reg *Register) MemSaveNodeToNodeList(node *types.Node) (int, error) {
	return 1, nil
}

/*
DeleteNodeFromMemNodeList 从内存中删除节点信息
*/
func (reg *Register) DeleteNodeFromMemNodeList() {

}

/*
NodeReg 节点上线并注册
*/
func (reg *Register) NodeReg() {

}

/*
NodeLostConnection 节点下线
*/
func (reg *Register) NodeLostConnection() {

}

/*
ChooseLeader 选择注册中心leader节点
*/
func (reg *Register) ChooseLeader() {

}

/*
Vote 投票
*/
func (reg *Register) Vote() {

}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	common.CheckError(err)
	fmt.Println(string(buffer))
}

/*
Listen 监听客户端连接
*/
func (reg *Register) Listen() {
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
