package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile("[0-9]+")

func main() {
	var filePath string
	if len(os.Args) < 1 {
		fmt.Println("input : " + os.Args[0] + " not found")
		os.Exit(1)
	} else {
		filePath = os.Args[1]
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	p1, err := CountOverlappingSectionAssignment(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", p1)
}

func MapSections(i []string) (ax, ay, bx, by int) {
	ax, _ = strconv.Atoi(i[0])
	ay, _ = strconv.Atoi(i[1])
	bx, _ = strconv.Atoi(i[2])
	by, _ = strconv.Atoi(i[3])
	return
}

func CountOverlappingSectionAssignment(input *os.File) (overlap int, err error) {
	err = IterateLines(input, func(s string) {
		var ax, ay, bx, by int = MapSections(re.FindAllString(s, -1))
		// check if A contains B
		if ax <= bx && ay >= by || bx <= ax && by >= ay {
			overlap++
		}
	})
	return
}
