package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

type Grid [][]int

func (g Grid) Get(x, y int) int { return g[y][x] }

func main() {
	g, err := CreateGrid(os.Stdin)
	if err != nil {
		panic(err)
	}

	count, err := CountVisibleTrees(g)
	if err != nil {
		panic(err)
	}

	fmt.Printf("part 1: %d\n", count)

	score, err := HighestScenicScore(g)
	if err != nil {
		panic(err)
	}

	fmt.Printf("part 2: %d\n", score)
}

func CreateGrid(input io.Reader) (grid Grid, err error) {
	var y int
	err = IterateLines(input, func(s string) Step {
		if grid == nil {
			grid = Grid{}

			for i := 0; i < len(s); i++ {
				grid = append(grid, make([]int, len(s)))
			}
		}

		for x, c := range s {
			if v, err := strconv.Atoi(string(c)); err == nil {
				grid[y][x] = v
			} else {
				panic(err)
			}
		}

		y++

		return Continue
	})

	return
}

// CountVisibleTrees counts the number of visible trees
// in the grid.
func CountVisibleTrees(grid Grid) (count int, err error) {
	grideSize := len(grid)

	count = (grideSize * 4) - 4

	// iterate trees (don't do the edges)
	for y := 1; y < grideSize-1; y++ {
		for x := 1; x < grideSize-1; x++ {
			if IsVisibleTree(grid, x, y) {
				count++
			}
		}
	}

	return
}

func IsVisibleTree(grid Grid, x, y int) (visible bool) {
	treeHeight := grid.Get(x, y)
	visibleSides := 0

	callBack := func(xb, yb int) Step {
		if t := grid.Get(xb, yb); treeHeight > t {
			// check if edge
			if IsEdge(grid, xb, yb) {
				// if edge then it's visible
				visibleSides++
				return Break
			}

			return Continue
		}

		return Break
	}

	// check up
	WalkGrid(grid, x, y-1, 0, -1, callBack)
	// check down
	WalkGrid(grid, x, y+1, 0, 1, callBack)
	// check left
	WalkGrid(grid, x-1, y, -1, 0, callBack)
	// check right
	WalkGrid(grid, x+1, y, 1, 0, callBack)

	return visibleSides >= 1
}

func HighestScenicScore(grid Grid) (score int, err error) {
	grideSize := len(grid)

	for y := 1; y < grideSize; y++ {
		for x := 1; x < grideSize; x++ {
			if s := ScenicScore(grid, x, y); s > score {
				score = s
			}
		}
	}

	return
}

func IsEdge(grid Grid, x, y int) bool {
	return x == 0 || y == 0 || x == len(grid)-1 || y == len(grid)-1
}

func ScenicScore(grid Grid, x, y int) (score int) {
	treeHeight := grid.Get(x, y)

	lineOfSights := []int{}
	lineOfSight := 0

	// stop if you reach an edge or at the first tree that is the same height or taller than the tree under consideration.
	callBack := func(xb, yb int) Step {
		lineOfSight++

		if t := grid.Get(xb, yb); treeHeight > t {
			return Continue
		}

		return Break
	}

	// check up
	WalkGrid(grid, x, y-1, 0, -1, callBack)
	lineOfSights = append(lineOfSights, lineOfSight)
	lineOfSight = 0

	// check left
	WalkGrid(grid, x-1, y, -1, 0, callBack)
	lineOfSights = append(lineOfSights, lineOfSight)
	lineOfSight = 0

	// check right
	WalkGrid(grid, x+1, y, 1, 0, callBack)
	lineOfSights = append(lineOfSights, lineOfSight)
	lineOfSight = 0

	// check down
	WalkGrid(grid, x, y+1, 0, 1, callBack)
	lineOfSights = append(lineOfSights, lineOfSight)
	lineOfSight = 0

	// fmt.Printf("x: %d, y: %d, lineOfSights: %v \n", x, y, lineOfSights)

	score = lineOfSights[0]
	for _, l := range lineOfSights[1:] {
		score *= l
	}

	return
}

func WalkGrid(grid Grid, x, y int, dX, dY int, f func(x, y int) Step) (err error) {
	for {
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid) {
			break
		}

		if c := f(x, y); c == Break {
			break
		}

		x += dX
		y += dY
	}

	return
}
