package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"github.com/dwethmar/adventofcode2022/day12/dijkstra"
	"github.com/dwethmar/adventofcode2022/pkg/iterate"
)

const Alphabet = "abcdefghijklmnopqrstuvwxyz"

type Point struct {
	X, Y int
}

func (p Point) String() string { return fmt.Sprintf("(%d, %d)", p.X, p.Y) }

func (p Point) Distance(p2 *Point) int {
	return int(math.Abs(float64(p.X-p2.X)) + math.Abs(float64(p.Y-p2.Y)))
}

func main() {
	grid, err := CreateGrid(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	length := GetPathToSignal(grid)

	fmt.Printf("Part 1: %d\n", length)
}

func GetPathToSignal(grid Grid) int {

	p := FindPoints(grid, 'S', 'E')
	if len(p) != 2 {
		log.Fatal("Could not find start and end point")
	}

	start := p[0]
	end := p[1]

	PrintGrid(grid, start, end)

	graph := dijkstra.NewGraph()

	grid.Iterate(func(x, y int, r rune) iterate.Step {
		c := &Point{x, y}
		height := GetHeight(grid, c)

		// Add edges
		for _, p := range GetNeighbors(grid, c) {
			if p.X < 0 || p.Y < 0 || p.X >= len(grid[0]) || p.Y >= len(grid) { // out of bounds
				continue
			}

			// the elevation of the destination square can be at most one higher than the elevation of your current square
			if h := GetHeight(grid, p); h > height+1 {
				continue
			}

			graph.AddEdge(c.String(), p.String(), 1)
		}

		return iterate.Continue
	})

	length, path := graph.Path(start.String(), end.String())

	fmt.Printf("Path: %v\n", path)

	return length
}

func GetHeight(grid Grid, p *Point) int {
	r := grid.Get(p.X, p.Y)
	if r == 'S' { // Your current position (S) has elevation a.
		return 0
	} else if r == 'E' { // location that should get the best signal (E) has elevation z.
		return 25
	}

	return strings.IndexRune(Alphabet, r)
}

func GetNeighbors(grid Grid, p *Point) (neighbors []*Point) {
	return []*Point{
		{p.X - 1, p.Y},
		{p.X + 1, p.Y},
		{p.X, p.Y - 1},
		{p.X, p.Y + 1},
	}
}
