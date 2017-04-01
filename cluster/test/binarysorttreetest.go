package main

import (
	"ceno/cluster/mem"
	"ceno/cluster/types"
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, "--------------------------")
	fmt.Fprintln(os.Stderr, "----BinarySortTreeTest----")
	fmt.Fprintln(os.Stderr, "--------------------------")
	var node *types.Node
	tree := mem.NewBinarySortTree(node)
	fmt.Fprintln(os.Stderr, tree)
}
