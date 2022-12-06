package main

import (
	"bufio"
	"io"
)

// IterateLines iterates over lines in a reader
// and calls the provided function for each line.
// if the function returns false, the iteration stops.
func IterateLines(r io.Reader, f func(i int, s string) bool) (err error) {
	scanner := bufio.NewScanner(r)

	var i int
	for scanner.Scan() {
		if c := f(i, scanner.Text()); !c {
			break
		}
		i++
	}

	return scanner.Err()
}
