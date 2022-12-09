package main

import (
	"bufio"
	"io"
)

type Step int

const (
	Continue Step = iota
	Break
)

// IterateLines iterates over lines in a reader
// and calls the provided function for each line.
// if the function returns false, the iteration stops.
func IterateLines(r io.Reader, f func(s string) Step) (err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if c := f(scanner.Text()); c == Break {
			break
		}
	}
	return scanner.Err()
}
