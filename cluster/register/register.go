package register

import "ceno/cluster/types"

/*
Register 注册中心
*/
type Register struct {
	Rid       string
	IP        string
	Port      int16
	Timestamp string
}

/*
LoadRegisterList 读取注册中心节点列表
*/
func (reg *Register) LoadRegisterList() {

}

/*
MemSaveNodeToNodeList 节点信息存入内存
*/
func MemSaveNodeToNodeList(node *types.Node) (int, error) {
	return 1, nil
}

/*
DeleteNodeFromMemNodeList 从内存中删除节点信息
*/
func DeleteNodeFromMemNodeList() {

}

/*
NodeReg 节点上线并注册
*/
func NodeReg() {

}

/*
NodeLostConnection 节点下线
*/
func NodeLostConnection() {

}

/*
ChooseLeader 选择注册中心leader节点
*/
func ChooseLeader() {

}

/*
Vote 投票
*/
func Vote() {

}

/*
Listen 监听客户端连接
*/
func Listen() {

}
