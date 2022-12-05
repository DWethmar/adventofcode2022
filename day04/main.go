package main

import (
	"fmt"
	"io"
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

	file.Seek(0, 0)
	p2, err := CountPartialOverlappingSectionAssignment(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part 2: %d\n", p2)

}

func MapSections(i []string) (x1, x2, y1, y2 int) {
	x1, _ = strconv.Atoi(i[0])
	x2, _ = strconv.Atoi(i[1])
	y1, _ = strconv.Atoi(i[2])
	y2, _ = strconv.Atoi(i[3])
	return
}

func CountOverlappingSectionAssignment(input io.Reader) (overlap int, err error) {
	err = IterateLines(input, func(s string) {
		var x1, x2, y1, y2 int = MapSections(re.FindAllString(s, -1))
		// check if A contains B
		if x1 <= y1 && x2 >= y2 || y1 <= x1 && y2 >= x2 {
			overlap++
		}
	})
	return
}

func CountPartialOverlappingSectionAssignment(input io.Reader) (overlap int, err error) {
	err = IterateLines(input, func(s string) {
		var x1, x2, y1, y2 int = MapSections(re.FindAllString(s, -1))
		// for any overlap
		if x1 <= y2 && y1 <= x2 {
			overlap++
		}
	})
	return
}
