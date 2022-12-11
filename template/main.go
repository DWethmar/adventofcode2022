package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/dwethmar/adventofcode2022/pkg/iterate"
)

func main() {
	rawBody, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	x, err := ReadAndParse(io.NopCloser(bytes.NewBuffer(rawBody)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", x)
}

func ReadAndParse(in io.Reader) (x int, err error) {
	err = iterate.Lines(in, func(s string) iterate.Step {
		return iterate.Continue
	})

	return
}
