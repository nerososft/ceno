package main

import (
	"ceno/cluster/config"
	"fmt"
)

func main() {
	conf := config.NewConfig("../register.json")

	ipt := conf.GetCurrentConfig().GetIPTable()
	fmt.Println(ipt)
}
