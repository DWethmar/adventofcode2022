package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

var commandRegex = regexp.MustCompile(`^\$\s([a-z]+)(?:\s(\S+))?`)

var listOutPutDirRegex = regexp.MustCompile(`^(^dir)\s(.+)$`)
var listOutputFileRegex = regexp.MustCompile(`^(\d+)\s(.+)`)

const (
	totalSize  = 70000000
	targetSize = 30000000
)

func main() {
	rawBody, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	tree, err := CreateTree(ioutil.NopCloser(bytes.NewBuffer(rawBody)))
	if err != nil {
		panic(err)
	}

	fmt.Printf("part 1: %d\n", SumNodesWithSizeLT(tree))
	fmt.Println("----")
	fmt.Printf("part 2: %d\n", SmallestFolderThatWouldFreeUpSpace(tree))
}

func CreateTree(in io.Reader) (node *Node, err error) {
	cmd := ""
	param := ""
	isCmd := false

	root := NewNode("")
	current := root

	err = IterateLines(in, func(s string) bool {
		// check if we have a command
		if isCmd = commandRegex.MatchString(s); isCmd {
			matches := commandRegex.FindStringSubmatch(s)
			cmd = matches[1]
			param = matches[2]
		}

		// fmt.Printf("isCmd: %v	cmd: %v	param: %v \n", isCmd, cmd, param)

		if root == nil {
			root = NewNode("")
			current = root
		}

		switch cmd {
		case "cd":
			// set current dir
			if param == ".." { // move up
				if current.Parent == nil {
					panic("cannot move up from root")
				} else {
					current = current.Parent
				}
			} else if param == "/" {
				if root == nil {
					panic("cannot move up from root")
				} else {
					current = root
				}
			} else { // move into sub dir
				// check if we have a child with the same name
				child := current.Child(param)

				if child == nil {
					child = NewNode(param)
					child.Parent = current
					current.Children = append(current.Children, child)
				}

				current = child
			}
		case "ls":
			// list files
			if !isCmd {
				// fmt.Printf("ls output: %s\n", s)

				// is dir?
				if listOutPutDirRegex.MatchString(s) { // add dir to current
					matches := listOutPutDirRegex.FindStringSubmatch(s)
					dirName := matches[2]

					// check if we have a child with the same name
					if child := current.Child(param); child == nil {
						child = NewNode(dirName)
						child.Parent = current
						current.Children = append(current.Children, child)
					}
				} else if listOutputFileRegex.MatchString(s) { // add file to current
					matches := listOutputFileRegex.FindStringSubmatch(s)
					fileName := matches[2]
					fileSize := matches[1]

					// check if we have a child with the same name
					// if not add it with the size
					if child := current.Child(param); child == nil {
						child = NewNode(fileName)
						if intVar, err := strconv.Atoi(fileSize); err == nil {
							child.Size = intVar
						} else {
							panic(err)
						}
						child.Parent = current
						current.Children = append(current.Children, child)
					}
				} else {
					panic("unknown ls output")
				}
			}
		}

		return true
	})

	return root, err
}

func SumNodesWithSizeLT(node *Node) int {
	var sum int

	IterateNodes(node, func(node *Node) bool {
		if node.Size == 0 {
			if size := NodeSize(node); size < 100000 {
				sum += size
			}
		}

		return true
	})

	return sum
}

// SmallestFolderThatWouldFreeUpSpace find a directory you can delete that will free up enough space
func SmallestFolderThatWouldFreeUpSpace(node *Node) int {
	smallestSize := 0

	unusedSpace := totalSize - NodeSize(node)

	IterateNodes(node, func(node *Node) bool {
		if node.Size == 0 {
			size := NodeSize(node)

			fmt.Printf("node: %s (%d) \n", NodePath(node), size)

			if unusedSpace+size >= targetSize {
				if smallestSize == 0 || size < smallestSize {
					smallestSize = size
				}
			}
		}

		return true
	})

	return smallestSize
}
