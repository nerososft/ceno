package types

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
Node 的节点信息
*/
type Node struct {
	Hash       int    `json:"hash"`        //随机哈希
	Name       string `json:"name"`        //主机名
	Port       int16  `json:"port"`        //节点开放端口
	Ip         string `json:"ip"`          //节点ip，或许可以不用string
	ProccessId int32  `json:"proccess_id"` //进程号
	Token      string `json:"token"`       //节点连接授权码
	Timestamp  string `json:"timestamp"`   //时间戳
}

func NewNode(hash int,
	name string,
	port int16,
	ip string,
	proccessid int32,
	token string,
	timestamp string) *Node {
	return &Node{
		Hash:       hash,
		Name:       name,
		Port:       port,
		Ip:         ip,
		ProccessId: proccessid,
		Token:      token,
		Timestamp:  timestamp,
	}
}

//序列化
func (node *Node) String() string {
	jsons, err := json.Marshal(node)
	if err != nil {
		fmt.Println("fatal error:", err.Error())
		os.Exit(1)
	}
	return string(jsons)
}
func (node *Node) GetHash() int {
	return node.Hash
}
