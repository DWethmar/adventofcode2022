package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var filePath string
	if len(os.Args) < 1 {
		fmt.Println("input : " + os.Args[0] + " file name")
		os.Exit(1)
	} else {
		fmt.Println("input file name : " + os.Args[1])
		filePath = os.Args[1]
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var largestSum int
	var sum int
	counter := NewSummarizer(0)

	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")

		if line == "" {
			counter = NewSummarizer(0)
			continue
		}

		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		if i == 0 {
			continue
		}

		if sum = counter(i); sum > largestSum {
			largestSum = sum
		}
	}

	if sum > largestSum {
		largestSum = sum
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("largest sum : ", largestSum)
	}
}

func NewSummarizer(sum int) func(i int) int {
	return func(i int) int {
		sum += i
		return sum
	}
}
