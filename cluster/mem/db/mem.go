package db

import (
	"ceno/cluster/mem/tree"
	"ceno/cluster/types"
)

/*
MemoryDB 内存数据库
*/
type MemoryDB struct {
	Tree *tree.BinarySortTree
}

/*
NewMemoryDB 创建内存数据库
*/
func NewMemoryDB() *MemoryDB {
	node := types.NewNode(0, "Initial node", 0, "000.000.000.000", 0, "", "")
	tree := tree.NewBinarySortTree(node)
	return &MemoryDB{
		Tree: tree,
	}
}

/*
MInsert  将节点插入内存
*/
func (db *MemoryDB) MInsert(node *types.Node) {
	tree.Insert(db.Tree, node)
}

/*
MDelete 根据hash值从内存删除
*/
func (db *MemoryDB) MDelete(hashid int) bool {
	return tree.Delete(db.Tree, hashid)
}

/*
MSelect 根据hash从内存查找
*/
func (db *MemoryDB) MSelect(hashid int) *types.Node {
	return tree.SearchNode(db.Tree, hashid)
}
