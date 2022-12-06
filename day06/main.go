package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	rawBody, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	marker1, err := GetMarkerFromSignal(ioutil.NopCloser(bytes.NewBuffer(rawBody)), 4)
	if err != nil {
		panic(err)
	}

	fmt.Printf("part 1: %d\n", marker1)

	os.Stdin.Seek(0, 0)
	marker2, err := GetMarkerFromSignal(ioutil.NopCloser(bytes.NewBuffer(rawBody)), 14)
	if err != nil {
		panic(err)
	}

	fmt.Printf("part 2: %d\n", marker2)
}

func GetMarkerFromSignal(in io.Reader, seqSize int) (marker int, err error) {
	err = IterateLines(in, func(x int, s string) bool {
		for i := range s {
			if i < seqSize {
				continue
			}

			sequence := s[i-seqSize : i]
			uniquenessSum := 0

			for _, c2 := range sequence {
				uniquenessSum += strings.Count(sequence, string(c2))
			}

			if uniquenessSum == seqSize {
				marker = i
				return false
			}
		}

		return true
	})

	return
}
