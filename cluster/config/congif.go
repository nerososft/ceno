package config

/*
Config 配置文件
*/
type Config struct {
	ConfPath string
}

/*
NodeIPTable 注册中心配置节点信息
*/
type NodeIPTable struct {
	IP   string
	port int16
}

/*
NewConfig 创建Config
*/
func NewConfig(confpath string) *Config {
	return nil
}

/*
LoadConfig 读取配置文件
*/
func (config *Config) LoadConfig() string {
	return ""
}

/*
GetIPTable 获取注册中心节点ip表
*/
func (config *Config) GetIPTable() []*NodeIPTable {
	return nil
}
