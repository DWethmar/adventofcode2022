package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func runningtime(s string) (string, time.Time) {
	log.Println("Start: ", s)
	return s, time.Now()
}

func track(s string, startTime time.Time) {
	endTime := time.Now()
	log.Printf("End: %s	%s", s, endTime.Sub(startTime))
}

func main() {
	defer track(runningtime("ultra top secret strategy guide"))

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

	fmt.Printf("Score: %d\n", calculateScore(file))
}

func calculateScore(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)

	var score int

	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		lineParts := strings.Split(line, " ")

		opponentMove := MapOpponentMove(lineParts[0])

		r := Round{
			OpponentMove: opponentMove,
			PlayerMove:   SelectPlayerMove(opponentMove, MapPlayerStrategy(lineParts[1])),
		}

		score += r.Score()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return score
}
