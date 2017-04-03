package main

import (
	"ceno/cluster/register"
)

func main() {
	reg := register.NewRegister("1", "localhost", 6060, "")
	reg.Listen()
}
