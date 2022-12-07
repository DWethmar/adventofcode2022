package main

import (
	"bufio"
	"io"
)

// IterateLines iterates over lines in a reader
// and calls the provided function for each line.
// if the function returns false, the iteration stops.
func IterateLines(r io.Reader, f func(s string) bool) (err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if c := f(scanner.Text()); !c {
			break
		}
	}
	return scanner.Err()
}
