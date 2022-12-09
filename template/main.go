package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	x, err := ReadAndParse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", x)
}

func ReadAndParse(in io.Reader) (x int, err error) {
	err = IterateLines(in, func(s string) Step {
		return Continue
	})

	return
}
