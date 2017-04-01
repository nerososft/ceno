package types

/*
Node 的节点信息
*/
type Node struct {
	hash       string //随机哈希
	name       string //主机名
	port       int16  //节点开放端口
	ip         string //节点ip，或许可以不用string
	proccessid int32  //进程号
	token      string //节点连接授权码
	timestamp  string //时间戳
}
