package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func runningtime(s string) (string, time.Time) {
	log.Println("Start:	", s)
	return s, time.Now()
}

func track(s string, startTime time.Time) {
	endTime := time.Now()
	log.Printf("End: %s	%s", s, endTime.Sub(startTime))
}

func main() {
	defer track(runningtime("get largest sum"))

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

	// fmt.Printf("Largest sums: %d\n", getLargestSums(file, 3))

	sum := 0
	for i, s := range getLargestSums(file, 3) {
		fmt.Printf("Sum: %d %d\n", i, s)
		sum += s
	}

	fmt.Printf("Sum of largest sums: %d\n", sum)
}

func getLargestSums(reader io.Reader, topN int) []int {
	scanner := bufio.NewScanner(reader)

	var sum int
	largestSums := make([]int, topN)
	counter := NewSummarizer(0)

	for scan := scanner.Scan(); scan; {
		line := strings.TrimSuffix(scanner.Text(), "\n")

		if line == "" {
			counter = NewSummarizer(0)
			// check if sum is larger than any of the largest sums
			for i, largestSum := range largestSums {
				if sum > largestSum {
					// shift all smaller sums down
					for j := topN - 1; j > i; j-- {
						largestSums[j] = largestSums[j-1]
					}

					for j := len(largestSums) - 1; j > i; j-- {
						largestSums[j] = largestSums[j-1]
					}

					largestSums[i] = sum
					break
				}
			}

			sum = 0
			scan = scanner.Scan()
			continue
		}

		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		sum = counter(i)
		scan = scanner.Scan()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return largestSums
}

func NewSummarizer(sum int) func(i int) int {
	return func(i int) int {
		sum += i
		return sum
	}
}
