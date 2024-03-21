package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Println(Part1(string(input)))
	fmt.Println(Part2(string(input)))
}

func Part1(input string) int {
	floor := 0

	for _, x := range input {
		switch x {
		case '(':
			floor++
		case ')':
			floor--
		default:
			continue
		}
	}
	return floor
}

func Part2(input string) int {
	floor := 0

	for i, x := range input {

		switch x {
		case '(':
			floor++
		case ')':
			floor--
		default:
			continue
		}

		if floor == -1 {
			return i + 1
		}
	}

	return -1
}
