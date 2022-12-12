package main

import (
	"fmt"
	"strings"

	"github.com/dwethmar/adventofcode2022/pkg/color"
)

const (
	TopLeftCorner     = '╔'
	TopRightCorner    = '╗'
	BottomLeftCorner  = '╚'
	BottomRightCorner = '╝'
	VerticalLine      = '║'
	HorizontalLine    = '═'
)

func PrintGrid(grid Grid, highlight ...*Point) {
	fmt.Print("  ")
	for i := 0; i < len(grid[0]); i++ {
		fmt.Printf("%d", i%10)
	}
	fmt.Println()

	fmt.Printf(" %s", string(TopLeftCorner))
	fmt.Print(strings.Repeat(string(HorizontalLine), len(grid[0])))
	fmt.Print(string(TopRightCorner))
	fmt.Println()

	for y, row := range grid {
		fmt.Printf("%d%s", y%10, string(VerticalLine))
		for x, r := range row {
			// Highlight
			h := false
			for _, p := range highlight {
				if p.X == x && p.Y == y {
					h = true
					break
				}
			}

			if h {
				fmt.Printf(color.Yellow+"%c"+color.Reset, r)
			} else {
				fmt.Printf("%c", r)
			}
		}

		fmt.Print(string(VerticalLine))
		fmt.Println()
	}

	fmt.Printf(" %s", string(BottomLeftCorner))
	fmt.Print(strings.Repeat(string(HorizontalLine), len(grid[0])))
	fmt.Print(string(BottomRightCorner))
	fmt.Println()
}
