package main

type Node struct {
	Parent   *Node
	Children []*Node

	Name string
	Size int
}

func (n *Node) Child(name string) *Node {
	var child *Node
	for _, c := range n.Children {
		if c.Name == name {
			child = c
			break
		}
	}

	return child
}

func NewNode(name string) *Node {
	return &Node{
		Name:     name,
		Children: []*Node{},
	}
}

func NodeSize(node *Node) int {
	sum := 0
	// sum size of all files in this folder
	for _, child := range node.Children {
		if child.Size > 0 {
			sum += child.Size
		} else {
			sum += NodeSize(child)
		}
	}

	return sum
}

func IterateNodes(node *Node, f func(node *Node) bool) {
	if !f(node) {
		return
	}

	for _, child := range node.Children {
		IterateNodes(child, f)
	}
}
