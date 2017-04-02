package main

import (
	"ceno/cluster/config"
)

func main() {
	conf := config.NewConfig("../register.conf")
	conf.GetCurrentConfig()

}
