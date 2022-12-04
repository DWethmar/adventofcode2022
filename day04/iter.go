package main

import (
	"bufio"
	"io"
)

func IterateLines(r io.Reader, f func(s string)) (err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	return scanner.Err()
}
