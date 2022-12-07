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

func main() {
	rawBody, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	marker1, err := SumFileSize(ioutil.NopCloser(bytes.NewBuffer(rawBody)))
	if err != nil {
		panic(err)
	}

	fmt.Printf("part 1: %d\n", marker1)
}

func SumFileSize(in io.Reader) (sum int, err error) {
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

	// PrintTree(root, 0)
	// PrintList(root)

	sum = SumFoldersOverN(root)

	return
}

func SumFoldersOverN(node *Node) int {
	var sum int

	IterateNodes(node, func(node *Node) bool {
		if node.Size == 0 {
			if size := NodeSize(node); size < 100000 {
				// fmt.Printf("node %s(%d) is smaller than 100000 %s \n", node.Name, NodeSize(node), NodePath(node))
				sum += size
			}
		}

		return true
	})

	return sum
}
