package main

import "ceno/cluster/config"
import "fmt"
import "os"

func main() {
	conf := config.NewConfig("../register.conf")
	fmt.Fprintln(os.Stderr, conf.LoadConfig())
}
