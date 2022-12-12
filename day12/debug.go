package main

import (
	"fmt"
	"strings"
)

const (
	TopLeftCorner     = '╔'
	TopRightCorner    = '╗'
	BottomLeftCorner  = '╚'
	BottomRightCorner = '╝'
	VerticalLine      = '║'
	HorizontalLine    = '═'
)

func PrintGrid(grid Grid) {
	fmt.Print("  ")
	for i := 0; i < len(grid[0]); i++ {
		fmt.Printf("%d", i%10)
	}
	fmt.Println()

	fmt.Printf(" %s", string(TopLeftCorner))
	fmt.Print(strings.Repeat(string(HorizontalLine), len(grid[0])))
	fmt.Print(string(TopRightCorner))
	fmt.Println()

	for i, row := range grid {
		fmt.Printf("%d%s", i%10, string(VerticalLine))
		for _, r := range row {
			fmt.Printf("%c", r)
		}
		fmt.Print(string(VerticalLine))
		fmt.Println()
	}

	fmt.Printf(" %s", string(BottomLeftCorner))
	fmt.Print(strings.Repeat(string(HorizontalLine), len(grid[0])))
	fmt.Print(string(BottomRightCorner))
	fmt.Println()
}
