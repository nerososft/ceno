package main

import (
	"ceno/cluster/mem/db"
	"ceno/cluster/types"
	"math/rand"

	"fmt"
	"os"

	uuid "github.com/nu7hatch/gouuid"
)

func main() {

	memdb := db.NewMemoryDB()
	for i := 0; i < 10; i++ {
		j := rand.Intn(1000)
		uuid, _ := uuid.NewV4()
		nodeR := types.NewNode(j, uuid.String(), 1, "33", int32(j), "ww", "ww")
		memdb.MInsert(nodeR)
	}
	fmt.Fprintln(os.Stderr, memdb.Tree.String())
	fmt.Fprintln(os.Stderr, "--------------------------")
	fmt.Fprintln(os.Stderr, memdb.MSelect(0).String())
	memdb.MDelete(0)
	fmt.Fprintln(os.Stderr, memdb.MSelect(0).String())
}
