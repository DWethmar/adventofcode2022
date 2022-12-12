package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dwethmar/adventofcode2022/day12/dijkstra"
	"github.com/dwethmar/adventofcode2022/pkg/iterate"
	"github.com/dwethmar/adventofcode2022/pkg/number"
)

const Alphabet = "abcdefghijklmnopqrstuvwxyz"

type Point struct {
	X, Y int
}

func (p Point) String() string { return fmt.Sprintf("(%d, %d)", p.X, p.Y) }

func main() {
	grid, err := CreateGrid(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	Part1(grid, CreateGraph(grid, func(a, b int) bool {
		// the elevation of the destination square can be at most one higher than the elevation of your current square
		return a >= b || a+1 == b
	}))

	Part2(grid, CreateGraph(grid, func(a, b int) bool {
		return a <= b || a == b+1
	}))
}

func Part1(grid Grid, graph *dijkstra.Graph) {
	start := FindPoints(grid, 'S')[0]

	length, path := graph.Path(start.String(), func(v interface{}) bool {
		return v == 'E'
	})

	points := make([]*Point, len(path))
	for i, p := range path {
		s := number.GetAllIntsFromString(p)
		points[i] = &Point{X: s[0], Y: s[1]}
	}

	PrintGrid(grid, points...)

	fmt.Printf("Part 1: %d\n", length)
}

func Part2(grid Grid, graph *dijkstra.Graph) {
	start := FindPoints(grid, 'E')[0]

	length, path := graph.Path(start.String(), func(v interface{}) bool {
		return v == 'a' || v == 'S'
	})

	points := make([]*Point, len(path))
	for i, p := range path {
		s := number.GetAllIntsFromString(p)
		points[i] = &Point{X: s[0], Y: s[1]}
	}

	PrintGrid(grid, points...)

	fmt.Printf("Part 2: %d\n", length)
}

func CreateGraph(grid Grid, isEdge func(a, b int) bool) *dijkstra.Graph {
	graph := dijkstra.NewGraph()

	grid.Iterate(func(x, y int, v rune) iterate.Step {
		c := &Point{X: x, Y: y}

		graph.AddNode(c.String(), v)

		return iterate.Continue
	})

	grid.Iterate(func(x, y int, v rune) iterate.Step {
		c := &Point{x, y}

		heightA := GetHeight(v)

		// Add edges
		for _, p := range GetNeighbors(grid, c) {
			if p.X < 0 || p.Y < 0 || p.X >= len(grid[0]) || p.Y >= len(grid) { // out of bounds
				continue
			}

			heightB := GetHeight(grid.Get(p.X, p.Y))

			if isEdge(heightA, heightB) {
				graph.AddEdge(c.String(), p.String(), 1)
			}
		}

		return iterate.Continue
	})

	return graph
}

func GetHeight(r rune) int {
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
