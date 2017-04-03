package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/*
Config 配置文件
*/
type Config struct {
	ConfPath string
}

//CurrentConfig 当前配置
type CurrentConfig struct {
	IPTable []struct {
		Node struct {
			IPAddr string `json:"ip_addr"`
			Port   int    `json:"port"`
		} `json:"node"`
	} `json:"ip_table"`
}

type IPTable struct {
	node Node
}

type Node struct {
	IPAddr string
	Port   int
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
GetCurrentConfig 获取当前配置
*/
func (config *Config) GetCurrentConfig() *CurrentConfig {
	conf := config.LoadConfig()

	var currentConfig *CurrentConfig
	json.Unmarshal([]byte(conf), &currentConfig)
	return currentConfig
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
此处应该使用动态数组
*/
func (currentConfig *CurrentConfig) GetIPTable() [6]IPTable {
	var table [6]IPTable
	for k, v := range currentConfig.IPTable {
		table[k] = IPTable{
			node: Node{
				IPAddr: v.Node.IPAddr,
				Port:   v.Node.Port,
			},
		}
	}
	return table
}
