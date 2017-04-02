package main

import (
	"ceno/cluster/mem/tree"
	"ceno/cluster/types"
	"fmt"
	"math/rand"
	"os"

	"github.com/nu7hatch/gouuid"
)

func main() {
	fmt.Fprintln(os.Stderr, "--------------------------")
	fmt.Fprintln(os.Stderr, "----BinarySortTreeTest----")
	fmt.Fprintln(os.Stderr, "--------------------------")
	//二叉排序树创建测试
	node := types.NewNode(11, "11", 12, "33", 12, "ww", "ww")
	fmt.Fprintln(os.Stderr, node.String())
	t := tree.NewBinarySortTree(node)

	//二叉排序树插入测试
	node1 := types.NewNode(0, "0", 12, "33", 12, "ww", "ww")
	tree.Insert(t, node1)
	node2 := types.NewNode(88, "88", 12, "33", 12, "ww", "ww")
	tree.Insert(t, node2)
	for i := 0; i < 10; i++ {
		j := rand.Intn(1000)
		uuid, _ := uuid.NewV4()
		nodeR := types.NewNode(j, uuid.String(), 1, "33", int32(j), "ww", "ww")
		tree.Insert(t, nodeR)
	}
	fmt.Fprintln(os.Stderr, t.String())

	//二叉排序树查找节点测试
	nn := tree.SearchNode(t, 88)
	fmt.Fprintln(os.Stderr, nn.String())

	//二叉排序树删除节点测试
	if tree.Delete(t, 11) {
		fmt.Fprintln(os.Stderr, "删除成功")
		nn := tree.SearchNode(t, 11)
		fmt.Fprintln(os.Stderr, nn.String())
	}

	fmt.Fprintln(os.Stderr, t.String())
}
