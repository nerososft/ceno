/*
二叉排序树数据库
*/
package mem

import "ceno/cluster/types"

/*
BinarySortTree 节点
*/
type BinarySortTree struct {
	Left  *BinarySortTree //左子树
	node  *types.Node     //节点信息
	Right *BinarySortTree //右子树
}

func Insert(tree *BinarySortTree, node *types.Node) *BinarySortTree {
	return nil
}

/*
NewBinarySortTree 创建二叉排序树
*/
func NewBinarySortTree(node *types.Node) *BinarySortTree {
	var t *BinarySortTree
	t = Insert(t, node)
	return t
}

func Search(tree *BinarySortTree, hash string) bool {
	if tree == nil {
		return false
	}
	switch {
	case hash == tree.node->hash:
		return true
	case hash < tree.node->hash:
		return Search(tree.Left, hash)
	case hash > tree.node->hash:
		return Search(tree.Right, hash)
	}
	return false
}
func Delete(tree *BinarySortTree,hash string) bool {
	if tree == nil {  
        return false  
    }  

    parent := tree  
    found := false  
    for {  
        if tree == nil {  
            break  
        }  
        if hash == t.node->hash {  
            found = true  
            break  
        }  
  
        parent = tree 
        if hash < t.node->hash { //left  
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
        fmt.Println("delete() 左右树都为空 ")  
        if parent.Left == tree {  
            parent.Left = nil  
        } else if parent.Right == tree {  
            parent.Right = nil  
        }  
        tree = nil  
        return true  
    }  
  
    if tree.Right == nil { //右树为空  
        fmt.Println("delete() 右树为空 ")  
        parent.Left = tree.Left.Left  
        parent.node = tree.Left.node  
        parent.Right = tree.Left.Right  
        tree.Left = nil  
        tree = nil  
        return true  
    }  
  
    if tree.Left == nil { //左树为空  
        fmt.Println("delete() 左树为空 ")  
        parent.Left = tree.Right.Left  
        parent.node = tree.Right.node  
        parent.Right = tree.Right.Right  
        tree.Right = nil  
        tree = nil  
        return true  
    }  
  
    fmt.Println("delete() 左右树都不为空 ")  
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
  
    tree.node = next.node  
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
func Walk(t *BinarySortTree, node chan *types.Node) {
	if t == nil {
		return
	}
	Walk(t.Left, node)
	node <- t.node
	Walk(t.Right, node)
}

// Walker launches Walk in a new goroutine,
// and returns a read-only channel of values.
func Walker(t *BinarySortTree) <-chan *types.Node {
	ch := make(chan *types.Node)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}

func Compare(t1, t2 *BinarySortTree) bool {
	c1, c2 := Walker(t1), Walker(t2)
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
