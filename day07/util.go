package main

import "fmt"

func PrintTree(node *Node, depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print("  ")
	}

	var name string
	if node.Size > 0 {
		name = fmt.Sprintf("file: %s (%d) path: %s", node.Name, node.Size, NodePath(node))
	} else {
		name = fmt.Sprintf("dir: %s (%d) path: %s", node.Name, NodeSize(node), NodePath(node))
	}

	fmt.Printf("%s\n", name)

	for _, child := range node.Children {
		PrintTree(child, depth+1)
	}
}

func PrintList(node *Node) {
	fmt.Printf("%s\n", fmt.Sprintf("%s %d", NodePath(node), NodeSize(node)))

	for _, child := range node.Children {
		PrintList(child)
	}
}

func NodePath(node *Node) string {
	var path string
	if node.Parent != nil {
		path = NodePath(node.Parent) + "/" + node.Name
	} else {
		path = node.Name
	}

	return path
}
