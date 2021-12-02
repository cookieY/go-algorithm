package bTree

import "fmt"

func (n *Node) PrintD() {
	fmt.Printf("%v \n", n.D)
}

func PreOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	node.PrintD()
	PreOrderTraversal(node.Left)
	PreOrderTraversal(node.Right)
}

func MidOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	MidOrderTraversal(node.Left)
	node.PrintD()
	MidOrderTraversal(node.Right)
}

func PostOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	PostOrderTraversal(node.Left)
	node.PrintD()
	PostOrderTraversal(node.Right)
}
