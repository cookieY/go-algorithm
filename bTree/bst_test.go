package bTree

import (
	"fmt"
	"testing"
)

func TestNewBST(t *testing.T) {
	tree := NewBST(nil)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(1)
	tree.Insert(4)

	fmt.Println("中序遍历二叉排序树：")
	MidOrderTraversal(tree.root)
	fmt.Println()
}
