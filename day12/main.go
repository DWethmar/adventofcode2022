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

	path := GetPathToSignal(grid)

	fmt.Printf("Part 1: %d: %+v\n", len(path)-1, path)
}

func GetPathToSignal(grid Grid) (path []*Point) {
	PrintGrid(grid)

	graph := dijkstra.NewGraph()

	cells := []string{}
	grid.Iterate(func(x, y int, r rune) iterate.Step {
		p := &Point{x, y}
		cells = append(cells, p.String())

		return iterate.Continue
	})

	nodes := dijkstra.AddNodes(graph, cells...)

	grid.Iterate(func(x, y int, r rune) iterate.Step {
		c := &Point{x, y}
		height := GetHeight(grid, c)

		// Add edges
		for _, p := range GetNeighbors(grid, c) {
			// in bounds
			if p.X < 0 || p.Y < 0 || p.X >= len(grid[0]) || p.Y >= len(grid) {
				continue
			}

			// the elevation of the destination square can be at most one higher than the elevation of your current square
			if GetHeight(grid, p) > height+1 {
				continue
			}

			graph.AddEdge(nodes[c.String()], nodes[p.String()], 1)
		}

		return iterate.Continue
	})

	dijkstra.Dijkstra(graph, nodes["(0, 0)"])

	return
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
