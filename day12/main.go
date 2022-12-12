package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
	"sync"
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

	p := FindPoints(grid, 'S', 'E')
	if len(p) != 2 {
		log.Fatal("Could not find start and end point")
	}

	return Walk(grid, p[0], p[1], []string{p[0].String()})
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

func Walk(grid Grid, start *Point, end *Point, walked []string) []*Point {
	path := []*Point{
		start,
	}

	currentHeight := GetHeight(grid, start)

	neighbors := []*Point{
		// up
		{start.X, start.Y - 1},
		// down
		{start.X, start.Y + 1},
		// left
		{start.X - 1, start.Y},
		// right
		{start.X + 1, start.Y},
	}

	// sort neighbors by distance to end
	sort.Slice(neighbors, func(i, j int) bool {
		return neighbors[i].Distance(end) < neighbors[j].Distance(end)
	})

	directions := map[string][]*Point{}

	wg := sync.WaitGroup{}

	for _, n := range neighbors {
		// Check if we are out of bounds
		if n.X < 0 || n.Y < 0 || n.X >= len(grid[0]) || n.Y >= len(grid) {
			continue
		}

		// Check if we have already walked this path
		if contains(walked, n.String()) {
			continue
		}

		height := GetHeight(grid, n)

		// Check if we can walk to this point
		if math.Abs(float64(currentHeight-height)) > 1 {
			continue
		}

		walked = append(walked, n.String())

		newWalked := make([]string, len(walked))
		copy(newWalked, walked)

		wg.Add(1)

		go func(p Point) {
			// Walk to the next point
			directions[p.String()] = Walk(grid, &p, end, newWalked)

			defer wg.Done()
		}(*n)
	}

	wg.Wait()

	shortest := 0
	shortestPath := ""

	// Find the shortest path
	for n, p := range directions {
		// Check if the path is empty
		if len(p) == 0 {
			continue
		}

		// has reached end?
		if p[len(p)-1].String() != end.String() {
			continue
		}

		// Check if this is the shortest path
		if shortest == 0 || len(p) < shortest {
			shortest = len(p)
			shortestPath = n
		}
	}

	// Add the shortest path to the path
	if shortestPath != "" {
		path = append(path, directions[shortestPath]...)
	}

	return path
}
