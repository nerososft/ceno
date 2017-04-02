package config

import (
	"ceno/common/json"
	"ceno/common/mmap"
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
	IPTables []NodeIPTable
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

type node struct {
	ip_addr string
	port    int
}

/*
GetCurrentConfig 获取当前配置
*/
func (config *Config) GetCurrentConfig() {
	conf := config.LoadConfig()
	confByte := []byte(conf)
	confMap := json.Json2Map(confByte)

	fmt.Println(os.Stderr, confMap)

	ipTable := confMap["ip_table"]
	if ipS, ok := ipTable.(map[string]interface{}); ok {
		for _, value := range ipS {
			// if node, nOk := value.(map[string]interface{}); nOk {
			// 	for k, v := range node {
			// 		fmt.Println(k, v)
			// 	}
			// } else {
			// 	fmt.Fprintln(os.Stderr, "fatal error: ip tables format error")
			// 	os.Exit(1)
			// }
			result := &node{}
			err := mmap.FillStruct(value.(map[string]interface{}), result)
			fmt.Println(err, fmt.Sprintf("%+v"), *result)
		}
	} else {
		fmt.Fprintln(os.Stderr, "fatal error: ip tables format error")
		os.Exit(1)
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
func (currentConfig *CurrentConfig) GetIPTable() []NodeIPTable {
	return currentConfig.IPTables
}
