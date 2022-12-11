package main

import (
	"fmt"
	"strings"

	"github.com/dwethmar/adventofcode2022/pkg/number"
)

type Monkey struct {
	// Name of the monkey
	Name string
	// StartingItems lists your worry level for each item the monkey is
	// currently holding in the order they will be inspected.
	StartingItems []int

	// Operation
	// example: new = old * 7
	Operation string

	// test
	Test      string
	TestTrue  string
	TestFalse string

	// Inspected total number of times each monkey inspects items
	inspected int
}

// hows how your worry level changes as that monkey inspects an item.
// (An operation like new = old * 5 means that your worry level after
// the monkey inspected the item is five times whatever your worry level was before inspection.)
func (m *Monkey) DoOperation(index int) (string, int) {
	item := m.StartingItems[index]

	p := strings.Split(m.Operation, " ")

	a := p[2]
	op := p[3]
	b := p[4]

	var valueA, valueB int
	if a == "old" {
		valueA = item
	} else {
		valueA = number.GetAllIntsFromString(a)[0]
	}

	if b == "old" {
		valueB = item
	} else {
		valueB = number.GetAllIntsFromString(b)[0]
	}

	result := 0
	operatorStr := ""
	valueBStr := fmt.Sprint(valueB)
	if valueBStr == fmt.Sprint(item) {
		valueBStr = "itself"
	}

	switch op {
	case "+":
		result = valueA + valueB
		operatorStr = "increases"
	case "*":
		result = valueA * valueB
		operatorStr = "is multiplied"
	}

	return fmt.Sprintf("Worry level %s by %s to %d.", operatorStr, valueBStr, result), result
}

// Test shows how the monkey uses your worry level to decide where
// to throw an item next.
// If true shows what happens with an item if the Test was true.
// If false shows what happens with an item if the Test was false.
func (m *Monkey) DoTest(worryLevel int) (string, int) {
	v := float64(number.GetAllIntsFromString(m.Test)[0])
	p := float64(worryLevel)

	// check if p is divisible by v
	isDivisible := p/v == float64(int(p/v))

	test := ""
	if isDivisible {
		test = m.TestTrue
	} else {
		test = m.TestFalse
	}

	successStr := ""
	if isDivisible {
		successStr = "is"
	} else {
		successStr = "is not"
	}

	return fmt.Sprintf("Current worry level %s divisible by %d.", successStr, int(v)), number.GetAllIntsFromString(test)[0]
}

func (m *Monkey) Inspect() {
	m.inspected++
}

func (m *Monkey) Inspected() int {
	return m.inspected
}
