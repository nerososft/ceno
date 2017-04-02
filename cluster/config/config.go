package config

import "io/ioutil"
import "fmt"
import "os"

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
	return &Config{
		ConfPath: confpath,
	}
}

/*
LoadConfig 读取配置文件
*/
func (config *Config) LoadConfig() string {
	dat, err := ioutil.ReadFile(config.ConfPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	return string(dat)
}

/*
GetIPTable 获取注册中心节点ip表
*/
func (config *Config) GetIPTable() []*NodeIPTable {
	return nil
}
