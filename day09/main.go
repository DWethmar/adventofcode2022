package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func main() {
	positionVisitedByTail, err := MoveRope(os.Stdin, (len(os.Args) >= 2 && os.Args[1] == "debug"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", positionVisitedByTail)
}

func MoveRope(input io.Reader, debug bool) (positionVisitedByTail int, err error) {
	head := &Point{X: 0, Y: 0}
	tail := &Point{X: 0, Y: 0}

	// positionVisitedByTail++
	positionsVisited := []*Point{}

	err = IterateLines(input, func(s string) Step {
		direction := ""
		distance := 0

		direction = s[0:1]
		if distance, err = strconv.Atoi(s[2:]); err != nil {
			return Break
		}

		for i := 0; i < distance; i++ {
			move := DrawGrid(head, tail, []*Point{})

			switch direction {
			case "U":
				head.Y++
			case "D":
				head.Y--
			case "L":
				head.X--
			case "R":
				head.X++
			}

			x, y := Follow(tail, head)
			tail.X += x
			tail.Y += y

			if x != 0 || y != 0 { // tail moved
				positionsVisited = append(positionsVisited, &Point{X: tail.X, Y: tail.Y})
				positionVisitedByTail++
			}

			if debug {
				fmt.Printf("Head: %v, Tail: %v Direction: %s, Distance: %d/%d Positions: %d \n", head, tail, direction, i+1, distance, positionVisitedByTail)

				fmt.Printf("%s", DrawLinesNextToEachOther(
					"move:\n"+move.String(),
					"follow:\n"+DrawGrid(head, tail, []*Point{}).String(),
					"visited:\n"+DrawGrid(nil, nil, positionsVisited).String(),
				))
			}
		}

		return Continue
	})

	return
}

func DrawGrid(head *Point, tail *Point, otherPoints []*Point) *strings.Builder {
	var strBuilder = &strings.Builder{}

	const width = 6
	const height = 5

	// draw grid with flipped Y axis
	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {
			if head != nil && x == head.X && y == head.Y {
				strBuilder.WriteString("H")
			} else if tail != nil && x == tail.X && y == tail.Y {
				strBuilder.WriteString("T")
			} else if x == 0 && y == 0 {
				strBuilder.WriteString("s")
			} else {
				match := false
				for _, p := range otherPoints {
					if p.X == x && p.Y == y {
						strBuilder.WriteString("#")
						match = true
						break
					}
				}

				if !match {
					strBuilder.WriteString(".")
				}
			}
		}

		strBuilder.WriteString("\n")
	}

	return strBuilder
}

func Follow(point *Point, target *Point) (x, y int) {
	// only follow if distance between two point is greater then 1
	if point.X == target.X && point.Y == target.Y {
		return 0, 0
	}

	// check if adjacent
	for tx := -1; tx <= 1; tx++ {
		for ty := -1; ty <= 1; ty++ {
			if point.X+tx == target.X && point.Y+ty == target.Y {
				return 0, 0
			}
		}
	}

	if point.X < target.X {
		x = 1
	} else if point.X > target.X {
		x = -1
	}

	if point.Y < target.Y {
		y = 1
	} else if point.Y > target.Y {
		y = -1
	}

	return
}
