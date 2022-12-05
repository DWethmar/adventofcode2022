package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var createRe = regexp.MustCompile(`\[([A-Z])\]|(\s{4})`)
var cmdRe = regexp.MustCompile(`[0-9]+`)

func main() {
	var filePath string
	if len(os.Args) < 1 {
		fmt.Println("input : " + os.Args[0] + " not found")
		os.Exit(1)
	} else {
		filePath = os.Args[1]
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	p1, err := RearrangeCreates(file, CrateMover9000)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %s\n", strings.Join(p1, ""))

	fmt.Println()

	file.Seek(0, io.SeekStart)
	p2, err := RearrangeCreates(file, CrateMover9001)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nPart 2: %s\n", strings.Join(p2, ""))
}


func ReverseSlice[T comparable](s []T) {
    sort.SliceStable(s, func(i, j int) bool {
        return i > j
    })
}

type CrateMover int
const (
	CrateMover9000 CrateMover = iota
	CrateMover9001
)

func RearrangeCreates(in io.Reader, crateMover CrateMover) (topCrates []string, err error) {
	var stackingCrates = true
	var crateStack [][]string

	err = IterateLines(in, func(s string) {
		if stackingCrates { // stack crates
			if s == "" {
				stackingCrates = false
				return
			}

			crates := createRe.FindAllStringSubmatch(s, -1)

			if crateStack == nil {
				crateStack = make([][]string, len(crates))
			}

			for i, crate := range crates {
				if crate[1] != "" {
					// prepend crate
					crateStack[i] = append([]string{crate[1]}, crateStack[i]...)
				}
			}
		} else { // move crates with commands
			command := cmdRe.FindAllString(s, -1)

			fmt.Printf("stack: %+v\n", crateStack)

			n, _ := strconv.Atoi(command[0]) // number of crates to move
			fromStack, _ := strconv.Atoi(command[1])
			toStack, _ := strconv.Atoi(command[2])

			fromStack -= 1
			toStack -= 1

			fmt.Printf("n: %d, fromStack: %d, toStack: %d \n", n, fromStack, toStack)

			crates := crateStack[fromStack][len(crateStack[fromStack])-n:]

			if crateMover == CrateMover9000 {
				ReverseSlice(crates)
			}

			// add crates to new stack
			crateStack[toStack] = append(crateStack[toStack], crates...)

			// remove crates from old stack
			crateStack[fromStack] = crateStack[fromStack][:len(crateStack[fromStack])-n]
		}
	})

	for _, stack := range crateStack {
		topCrates = append(topCrates, stack[len(stack)-1])
	}

	return
}
