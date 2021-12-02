package bTree

import "testing"

func TestPreorderTraversal(t *testing.T) {
	node1 := NewNode(0) // 根节点
	node2 := NewNode(1)
	node3 := NewNode(2.0)
	node1.Left = node2
	node1.Right = node3
	PreOrderTraversal(node1)
}
