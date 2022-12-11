package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
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

	monkeys, _, err := CreateMonkeys(io.NopCloser(bytes.NewBuffer(rawBody)))
	if err != nil {
		log.Fatal(err)
	}

	part1(monkeys)

	monkeys, lcm, err := CreateMonkeys(io.NopCloser(bytes.NewBuffer(rawBody)))
	if err != nil {
		log.Fatal(err)
	}

	part2(monkeys, lcm)
}

func part1(monkeys []*Monkey) {
	relieve := func(w int) int {
		return w / 3
	}

	for i := 0; i < 20; i++ {
		PlayKeepAway(monkeys, relieve)
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

func part2(monkeys []*Monkey, lcm int) {
	relieve := func(worryLvl int) int {
		return worryLvl % lcm
	}

	for i := 0; i < 10000; i++ {
		PlayKeepAway(monkeys, relieve)

		if i == 0 {
			fmt.Printf("After %d rounds:\n", i)
			for _, monkey := range monkeys {
				fmt.Printf("%s inspected items %d times.\n", monkey.Name, monkey.Inspected())
			}
			fmt.Println()
		}
	}

	for _, monkey := range monkeys {
		fmt.Printf("%s inspected items %d times.\n", monkey.Name, monkey.Inspected())
	}

	// sort monkeys based on Inspected()
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Inspected() > monkeys[j].Inspected()
	})

	answer := monkeys[0].Inspected() * monkeys[1].Inspected()

	fmt.Printf("Part 2: %s(%d) * %s(%d) = %d\n", monkeys[0].Name, monkeys[0].Inspected(), monkeys[1].Name, monkeys[1].Inspected(), answer)
}

func CreateMonkeys(input io.Reader) ([]*Monkey, int, error) {
	var monkeys []*Monkey

	currentMonkey := &Monkey{}
	i := 0
	lcm := 1

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
			currentMonkey.Items = number.GetAllIntsFromString(v[1])
		case 2:
			currentMonkey.Operation = strings.Trim(v[1], " ")
		case 3:
			currentMonkey.Test = number.GetAllIntsFromString(instruction)[0]
			lcm *= currentMonkey.Test
		case 4:
			currentMonkey.TestTrue = number.GetAllIntsFromString(instruction)[0]
		case 5:
			currentMonkey.TestFalse = number.GetAllIntsFromString(instruction)[0]
		}

		if i == 5 {
			i = 0
		} else {
			i++
		}

		return iterate.Continue
	})

	return monkeys, lcm, err
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
func PlayKeepAway(monkeys []*Monkey, relieve func(w int) int) {
	for _, monkey := range monkeys {
		for i, _ := range monkey.Items {
			w := relieve(monkey.DoOperation(i))
			t := monkey.DoTest(w)

			receivingMonkey := monkeys[t]
			receivingMonkey.Items = append(receivingMonkey.Items, w)

			monkey.Inspect()
		}

		monkey.Items = nil
	}
}
