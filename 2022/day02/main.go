package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	games := strings.Split(string(input), "\n")
	part1 := part1(games)
	fmt.Println(part1)
	part2 := part2(games)
	fmt.Println(part2)
}

func part2(games []string) int {
	rulesMap := map[string]int{
		"AX": 3,
		"BX": 1,
		"CX": 2,
		"AY": 4,
		"BY": 5,
		"CY": 6,
		"AZ": 8,
		"BZ": 9,
		"CZ": 7,
	}
	sum := 0
	for _, game := range games {
		if game == "" {
			continue
		}
		sum += rulesMap[game]
	}
	return sum
}

func part1(games []string) int {
	rulesMap := map[string]int{
		"AX": 4,
		"BX": 1,
		"CX": 7,
		"AY": 8,
		"BY": 5,
		"CY": 2,
		"AZ": 3,
		"BZ": 9,
		"CZ": 6,
	}
	sum := 0
	for _, game := range games {
		if game == "" {
			continue
		}
		sum += rulesMap[game]
	}
	return sum
}
