package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/dwethmar/adventofcode2022/pkg/iterate"
)

const (
	Lit  string = "##"
	Dark string = ".."
)

type Command struct {
	Value  int
	Cycles int
}

func main() {
	rawBody, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	x, err := ReadAndParse(io.NopCloser(bytes.NewBuffer(rawBody)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", x)
}

func ReadAndParse(in io.Reader) (signalStrengths int, err error) {
	x := 1
	output := &strings.Builder{}

	noop := func() Command {
		return Command{
			Value:  0,
			Cycles: 1,
		}
	}

	addX := func(x int) Command {
		return Command{
			Value:  x,
			Cycles: 2,
		}
	}

	crtRow := []string{}
	cycle := 0

	err = iterate.Lines(in, func(instruction string) iterate.Step {
		var cmd Command

		switch instruction[:4] {
		case "noop":
			cmd = noop()
		case "addx":
			x, err := strconv.Atoi(instruction[5:])
			if err != nil {
				panic(err)
			}
			cmd = addX(x)
		}

		for i := 0; i < cmd.Cycles; i++ {
			cycle++
			fmt.Printf("cycle: %d X: %d \t cmd: %s \t %d/%d", cycle, x, instruction, i+1, cmd.Cycles)

			// every 20th, 60th, 100th, 140th, 180th, and 220th
			if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
				signalStrengths += cycle * x
				fmt.Printf(" [signalStrengths: %d * %d = %d sum: %d]", cycle, x, cycle*x, signalStrengths)
			}

			/// Sprite position is x

			if i == cmd.Cycles-1 && cmd.Value != 0 {
				x += cmd.Value
				fmt.Printf(" [ended x = %d]", x)
			}

			// sprite
			// if sprite has overlap in
			crtPos := len(crtRow) + 1
			if crtPos == x-1 || crtPos == x || crtPos == x+1 {
				crtRow = append(crtRow, Lit)
			} else {
				crtRow = append(crtRow, Dark)
			}

			// output
			if cycle%40 == 0 {
				output.WriteString(strings.Join(crtRow, "") + "\n")
				crtRow = []string{}
			}
			// output]

			fmt.Println()
		}

		return iterate.Continue
	})

	fmt.Println(output.String())

	return
}
