package bTree

import "errors"

type BST struct {
	root *Node
}

func NewBST(node *Node) *BST {
	return &BST{root: node}
}

func (b *BST) Insert(d int) error {

	if node := b.root; node == nil {
		b.root = NewNode(d)
		return nil
	} else {
		if d < node.D {
			if node.Left == nil {
				node.Left = NewNode(d)
			}
			node.Left.D = d
		} else if d > node.D {
			if node.Right == nil {
				node.Right = NewNode(d)
			}
			node.Right.D = d
		} else {
			errors.New("node is exist")
		}
		return nil
	}
}
