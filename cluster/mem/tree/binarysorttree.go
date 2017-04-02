package tree

import (
	"ceno/cluster/types"
	"encoding/json"
	"fmt"
	"os"
)

/*
BinarySortTree 节点
*/
type BinarySortTree struct {
	Left  *BinarySortTree `json:"left"`  //左子树
	Node  *types.Node     `json:"node"`  //节点信息
	Right *BinarySortTree `json:"right"` //右子树
}

/*
Insert 插入节点
*/
func Insert(tree *BinarySortTree, node *types.Node) *BinarySortTree {
	if tree == nil {
		return &BinarySortTree{nil, node, nil}
	}
	if node.GetHash() < tree.Node.GetHash() {
		tree.Left = Insert(tree.Left, node)
		return tree
	}
	tree.Right = Insert(tree.Right, node)
	return tree
}

/*
GetMin 获取hashid最小的节点
*/
func GetMin(tree *BinarySortTree) (*types.Node, bool) {
	if tree == nil {
		return nil, false
	}

	for {
		if tree.Left != nil {
			tree = tree.Left
		} else {
			return tree.Node, true
		}
	}
}

/*GetMax 获取hashid最大的节点*/
func GetMax(tree *BinarySortTree) (*types.Node, bool) {
	if tree == nil {
		return nil, false
	}
	for {
		if tree.Right != nil {
			tree = tree.Right
		} else {
			return tree.Node, true
		}
	}
}

/*
NewBinarySortTree 创建二叉排序树
*/
func NewBinarySortTree(node *types.Node) *BinarySortTree {
	var tree *BinarySortTree
	tree = Insert(tree, node)
	return tree
}

/*
Search 根据hash搜索节点
*/
func Search(tree *BinarySortTree, hash int) bool {
	if SearchNode(tree, hash) == nil {
		return false
	}
	return true
}

/*SearchNode ...*/
func SearchNode(tree *BinarySortTree, hash int) *types.Node {
	if tree == nil {
		return nil
	}
	switch {
	case hash == tree.Node.GetHash():
		return tree.Node
	case hash < tree.Node.GetHash():
		return SearchNode(tree.Left, hash)
	case hash > tree.Node.GetHash():
		return SearchNode(tree.Right, hash)
	}
	return nil
}

/*
Delete 删除节点
*/
func Delete(tree *BinarySortTree, hash int) bool {
	if tree == nil {
		return false
	}

	parent := tree
	found := false
	for {
		if tree == nil {
			break
		}
		if hash == tree.Node.GetHash() {
			found = true
			break
		}

		parent = tree
		if hash < tree.Node.GetHash() { //left
			tree = tree.Left
		} else {
			tree = tree.Right
		}
	}

	if found == false {
		return false
	}
	return deleteNode(parent, tree)
}

func deleteNode(parent, tree *BinarySortTree) bool {
	if tree.Left == nil && tree.Right == nil {
		//fmt.Println("delete() 左右树都为空 ")
		if parent.Left == tree {
			parent.Left = nil
		} else if parent.Right == tree {
			parent.Right = nil
		}
		tree = nil
		return true
	}

	if tree.Right == nil { //右树为空
		//fmt.Println("delete() 右树为空 ")
		parent.Left = tree.Left.Left
		parent.Node = tree.Left.Node
		parent.Right = tree.Left.Right
		tree.Left = nil
		tree = nil
		return true
	}

	if tree.Left == nil { //左树为空
		//fmt.Println("delete() 左树为空 ")
		parent.Left = tree.Right.Left
		parent.Node = tree.Right.Node
		parent.Right = tree.Right.Right
		tree.Right = nil
		tree = nil
		return true
	}

	//fmt.Println("delete() 左右树都不为空 ")
	previous := tree
	//找到左子节点的最右叶节点，将其值替换至被删除节点
	//然后将这个最右叶节点清除，所以说，为了维持树，
	//这种情况下，这个最右叶节点才是真正被删除的节点
	next := tree.Left
	for {
		if next.Right == nil {
			break
		}
		previous = next
		next = next.Right
	}

	tree.Node = next.Node
	if previous.Left == next {
		previous.Left = next.Left
	} else {
		previous.Right = next.Right
	}
	next.Left = nil
	next.Right = nil
	next = nil
	return true
}

// Walk traverses a tree depth-first,
// sending each Value on a channel.
func Walk(tree *BinarySortTree, node chan *types.Node) {
	if tree == nil {
		return
	}
	Walk(tree.Left, node)
	node <- tree.Node
	Walk(tree.Right, node)
}

// Walker launches Walk in a new goroutine,
// and returns a read-only channel of values.
func Walker(tree *BinarySortTree) <-chan *types.Node {
	ch := make(chan *types.Node)
	go func() {
		Walk(tree, ch)
		close(ch)
	}()
	return ch
}

/*
Compare ...
*/
func Compare(tree1, tree2 *BinarySortTree) bool {
	c1, c2 := Walker(tree1), Walker(tree2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			break
		}
	}
	return false
}

func (tree *BinarySortTree) String() string {
	jsons, err := json.Marshal(tree)
	if err != nil {
		fmt.Println("fatal error:", err.Error())
		os.Exit(1)
	}
	return string(jsons)
}
