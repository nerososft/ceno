package main

import (
	"ceno/cluster/mem"
	"ceno/cluster/types"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, "--------------------------")
	fmt.Fprintln(os.Stderr, "----BinarySortTreeTest----")
	fmt.Fprintln(os.Stderr, "--------------------------")
	//二叉排序树创建测试
	node := types.NewNode(11, "11", 12, "33", 12, "ww", "ww")
	fmt.Fprintln(os.Stderr, node.String())
	tree := mem.NewBinarySortTree(node)

	//二叉排序树插入测试
	node1 := types.NewNode(0, "0", 12, "33", 12, "ww", "ww")
	mem.Insert(tree, node1)
	node2 := types.NewNode(88, "88", 12, "33", 12, "ww", "ww")
	mem.Insert(tree, node2)
	for i := 0; i < 100; i++ {
		j := rand.Intn(1000)
		nodeR := types.NewNode(j, "88", 1, "33", int32(j), "ww", "ww")
		mem.Insert(tree, nodeR)
	}
	fmt.Fprintln(os.Stderr, tree.String())

	//二叉排序树查找节点测试
	nn := mem.SearchNode(tree, 88)
	fmt.Fprintln(os.Stderr, nn.String())

	//二叉排序树删除节点测试
	if mem.Delete(tree, 11) {
		fmt.Fprintln(os.Stderr, "删除成功")
		nn := mem.SearchNode(tree, 11)
		fmt.Fprintln(os.Stderr, nn.String())
	}

	fmt.Fprintln(os.Stderr, tree.String())
}
