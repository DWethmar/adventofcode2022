package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	Alphabet = "abcdefghijklmnopqrstuvwxyz"
)

var (
	priority = Alphabet + strings.ToUpper(Alphabet)
)

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

	p1, err := PartOne(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", p1)

	// Part 2
	file.Seek(0, 0)

	p2, err := PartTwo(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 2: %d\n", p2)
}

func IterateLines(r io.Reader, f func(s string)) (err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	return scanner.Err()
}

func PartOne(input *os.File) (prioritySum int, err error) {
	err = IterateLines(input, func(s string) {
		c2 := s[len(s)/2:]

		for _, c := range s[0 : len(s)/2] {
			if strings.ContainsRune(c2, c) {
				prioritySum += strings.IndexRune(priority, c) + 1
				break
			}
		}
	})
	return
}

func PartTwo(input *os.File) (prioritySum int, err error) {
	rucksacks := []string{}
	countedBadges := ""

	err = IterateLines(input, func(s string) {
		rucksacks = append(rucksacks, s)

		if len(rucksacks) == 3 {
			for _, c := range strings.Join(rucksacks, "") {
				if strings.ContainsRune(countedBadges, c) {
					continue
				}

				o := 0

				for i, r := range rucksacks {
					if strings.ContainsRune(r, c) {
						o++
					} else if i > 0 {
						continue
					}
				}

				if o == 3 {
					prioritySum += strings.IndexRune(priority, c) + 1
				}

				countedBadges += string(c)
			}

			rucksacks = []string{}
			countedBadges = ""
		}
	})

	return
}
