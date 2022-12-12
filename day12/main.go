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

func main() {
	grid, err := CreateGrid(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	graph := MakeGraph(grid)

	// part1(grid, graph)
	part2(grid, graph)
}

func part1(grid Grid, graph *dijkstra.Graph) {
	points := FindPoints(grid, 'S', 'E')
	if len(points) != 2 {
		log.Fatal("Could not find start points")
	}

	length, _ := GetPathToSignal(graph, points[0], points[1])

	fmt.Printf("Part 1: %d\n", length)
}

func part2(grid Grid, graph *dijkstra.Graph) {
	startPoints := FindPoints(grid, 'S', 'a')
	if len(startPoints) < 1 {
		log.Fatal("Could not find start points")
	}

	fmt.Printf("Start points: %v \n", len(startPoints))

	endPoints := FindPoints(grid, 'E')
	if len(endPoints) != 1 {
		log.Fatal("Could not find end points")
	}

	PrintGrid(grid, append(startPoints, endPoints[0])...)

	var length = math.MaxInt32

	for _, start := range startPoints {
		l, path := GetPathToSignal(graph, start, endPoints[0])

		if len(path) > 0 && l < length {
			length = l
		}
	}

	fmt.Printf("Part 2: %d\n", length)
}

func MakeGraph(grid Grid) *dijkstra.Graph {
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

	return graph
}

func GetPathToSignal(graph *dijkstra.Graph, start *Point, end *Point) (int, []string) {
	return graph.Path(start.String(), end.String())
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
