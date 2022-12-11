package number

import (
	"regexp"
	"strconv"
)

var (
	numbers = regexp.MustCompile(`-?\d+`)
)

func GetAllIntsFromString(s string) []int {
	var ints []int
	for _, match := range numbers.FindAllString(s, -1) {
		ints = append(ints, MustAtoi(match))
	}
	return ints
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
