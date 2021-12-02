package bTree

type Node struct {
	D     int // node data
	Left  *Node
	Right *Node
}

func NewNode(d int) *Node {
	return &Node{
		D: d,
	}
}
