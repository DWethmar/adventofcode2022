package main

import (
	"strings"

	"github.com/dwethmar/adventofcode2022/pkg/number"
)

type Monkey struct {
	// Name of the monkey
	Name string
	// Items lists your worry level for each item the monkey is
	// currently holding in the order they will be inspected.
	Items []int

	// Operation
	Operation string

	// test
	Test      int
	TestTrue  int
	TestFalse int

	// Inspected total number of times each monkey inspects items
	inspected int
}

// hows how your worry level changes as that monkey inspects an item.
// (An operation like new = old * 5 means that your worry level after
// the monkey inspected the item is five times whatever your worry level was before inspection.)
func (m *Monkey) DoOperation(index int) int {
	item := m.Items[index]

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

	switch op {
	case "+":
		result = valueA + valueB
	case "*":
		result = valueA * valueB
	}

	return result
}

// Test shows how the monkey uses your worry level to decide where
// to throw an item next.
// If true shows what happens with an item if the Test was true.
// If false shows what happens with an item if the Test was false.
func (m *Monkey) DoTest(w int) int {
	if w%m.Test == 0 {
		return m.TestTrue
	}
	return m.TestFalse
}

func (m *Monkey) Inspect() {
	m.inspected++
}

func (m *Monkey) Inspected() int {
	return m.inspected
}
