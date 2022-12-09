package main

import "strings"

func DrawLinesNextToEachOther(i ...string) (out string) {
	lines := make([][]string, len(i))

	for i, x := range i {
		lines[i] = strings.Split(x, "\n")
	}

	for y := 0; y < len(lines[0]); y++ {
		for x := 0; x < len(lines); x++ {
			out += lines[x][y] + "\t"
		}
		out += "\n"
	}

	return
}
