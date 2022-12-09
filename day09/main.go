package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func (c *Point) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
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

	positionsVisited := []*Point{}

	pointsVisited := map[string]struct{}{}

	err = IterateLines(input, func(s string) Step {
		direction := ""
		distance := 0

		direction = s[0:1]
		if distance, err = strconv.Atoi(s[2:]); err != nil {
			return Break
		}

		for i := 0; i < distance; i++ {

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

			move := DrawGrid(head, tail, []*Point{})

			Follow(head, tail)

			if _, ok := pointsVisited[tail.String()]; !ok {
				pointsVisited[tail.String()] = struct{}{}
				positionsVisited = append(positionsVisited, &Point{X: tail.X, Y: tail.Y})
				positionVisitedByTail = len(positionsVisited)
			}

			if debug {
				fmt.Printf("Head: %v, Tail: %v Direction: %s, Distance: %d/%d Positions: %d \n", head, tail, direction, i+1, distance, positionVisitedByTail)

				fmt.Printf("%s", DrawLinesNextToEachOther(
					"Head:\n"+move.String(),
					"Tail:\n"+DrawGrid(head, tail, []*Point{}).String(),
					"visited:\n"+DrawGrid(nil, nil, positionsVisited).String(),
				))
			}
		}

		return Continue
	})

	return
}

func Follow(head *Point, tail *Point) {
	// If the head is right next to or directly on top of the tail then we do nothing
	diffInX := math.Abs(float64(tail.X - head.X))
	diffInY := math.Abs(float64(tail.Y - head.Y))

	if diffInY <= 1 && diffInX <= 1 {
		return
	}

	//fmt.Println(fmt.Sprintf("Tail moving (%s) to meet head (%s)", tail.String(), head.String()))
	xDirection := 1
	yDirection := 1

	if head.Y < tail.Y {
		yDirection = -1
	}

	if head.X < tail.X {
		xDirection = -1
	}

	if head.X == tail.X {
		tail.Y += 1 * yDirection
		return
	}

	if head.Y == tail.Y {
		tail.X += 1 * xDirection
		return
	}

	// Diagonal
	tail.X += 1 * xDirection
	tail.Y += 1 * yDirection
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
