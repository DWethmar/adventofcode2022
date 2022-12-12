package main

import (
	"io"

	"github.com/dwethmar/adventofcode2022/pkg/iterate"
)

type Grid [][]rune

func (g Grid) Get(x, y int) rune {
	return g[y][x]
}

func (g Grid) Iterate(f func(x, y int, r rune) iterate.Step) {
	for y, row := range g {
		for x, r := range row {
			if f(x, y, r) == iterate.Break {
				return
			}
		}
	}
}

func CreateGrid(in io.Reader) (grid Grid, err error) {
	i := 0
	err = iterate.Lines(in, func(s string) iterate.Step {
		if grid == nil {
			grid = make(Grid, 0)
		}

		grid = append(grid, make([]rune, len(s)))

		for j, r := range s {
			grid[i][j] = r
		}

		i++

		return iterate.Continue
	})

	return
}

func FindPoints(grid Grid, p ...rune) []*Point {
	points := make([]*Point, 0)

	for y, row := range grid {
		for x, r := range row {
			for _, c := range p {
				if r == c {
					points = append(points, &Point{x, y})
				}
			}
		}
	}

	return points
}
