package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/dwethmar/adventofcode2022/pkg/iterate"
	"github.com/dwethmar/adventofcode2022/pkg/number"
)

func main() {
	rawBody, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	monkeys, err := CreateMonkeys(io.NopCloser(bytes.NewBuffer(rawBody)))
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 20; i++ {
		PlayKeepAway(i+1, monkeys)
	}

	fmt.Println()

	for _, monkey := range monkeys {
		fmt.Printf("%s inspected items %d times.\n", monkey.Name, monkey.Inspected())
	}

	// sort monkeys based on Inspected()
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Inspected() > monkeys[j].Inspected()
	})

	answer := monkeys[0].Inspected() * monkeys[1].Inspected()

	fmt.Printf("Part 1: %s(%d) * %s(%d) = %d\n", monkeys[0].Name, monkeys[0].Inspected(), monkeys[1].Name, monkeys[1].Inspected(), answer)
}

func CreateMonkeys(input io.Reader) ([]*Monkey, error) {
	var monkeys []*Monkey

	currentMonkey := &Monkey{}
	i := 0

	err := iterate.Lines(input, func(instruction string) iterate.Step {
		if instruction == "" {
			return iterate.Continue
		}

		v := strings.Split(instruction, ":")

		switch i {
		case 0:
			currentMonkey = &Monkey{}
			monkeys = append(monkeys, currentMonkey)

			currentMonkey.Name = strings.Trim(v[0], " ")
		case 1:
			currentMonkey.StartingItems = number.GetAllIntsFromString(v[1])
		case 2:
			currentMonkey.Operation = strings.Trim(v[1], " ")
		case 3:
			currentMonkey.Test = strings.Trim(v[1], " ")
		case 4:
			currentMonkey.TestTrue = strings.Trim(v[1], " ")
		case 5:
			currentMonkey.TestFalse = strings.Trim(v[1], " ")
		}

		if i == 5 {
			i = 0
		} else {
			i++
		}

		return iterate.Continue
	})

	return monkeys, err
}

/**
Monkey 0:
  Monkey inspects an item with a worry level of 79.
    Worry level is multiplied by 19 to 1501.
    Monkey gets bored with item. Worry level is divided by 3 to 500.
    Current worry level is not divisible by 23.
    Item with worry level 500 is thrown to monkey 3.
  Monkey inspects an item with a worry level of 98.
    Worry level is multiplied by 19 to 1862.
    Monkey gets bored with item. Worry level is divided by 3 to 620.
    Current worry level is not divisible by 23.
    Item with worry level 620 is thrown to monkey 3.
**/
func PlayKeepAway(round int, monkeys []*Monkey) {

	for _, monkey := range monkeys {
		fmt.Print(monkey.Name, ":\n")

		for i, item := range monkey.StartingItems {
			fmt.Printf("\tMonkey inspects an item with a worry level of %d.\n", item)

			d, newWorryLevel := monkey.DoOperation(i)
			fmt.Printf("\t\t%s\n", d)

			newWorryLevel = int(math.Floor(float64(newWorryLevel) / 3))
			fmt.Printf("\t\tMonkey gets bored with item. Worry level is divided by 3 to %d.\n", newWorryLevel)

			t, throwToMonkeyIndex := monkey.DoTest(newWorryLevel)
			fmt.Printf("\t\t%s\n", t)

			fmt.Printf("\t\tItem with worry level %d is thrown to monkey %d.\n", newWorryLevel, throwToMonkeyIndex)
			receivingMonkey := monkeys[throwToMonkeyIndex]
			receivingMonkey.StartingItems = append(receivingMonkey.StartingItems, newWorryLevel)

			monkey.Inspect()
		}

		monkey.StartingItems = []int{}
	}

	fmt.Println("")

	fmt.Printf("After round %d, the monkeys are holding items with these worry levels:\n", round)
	for _, monkey := range monkeys {
		fmt.Printf("%s: %d\n", monkey.Name, monkey.StartingItems)
	}

	return
}
