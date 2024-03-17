package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	resPart1 := 0
	resPart2 := 0

	for lineNumber := 1; scanner.Scan(); lineNumber++ {
		line := scanner.Text()

		if Part1(line) {
			resPart1 += lineNumber
		}

		powerNumber := Part2(line)
		resPart2 += powerNumber
	}

	fmt.Printf("Part_1: %d\n", resPart1)
	fmt.Printf("Part_2: %d\n", resPart2)
}

func Part1(input string) bool {
	game_subset := strings.Split((strings.Split(input, ":")[1]), ";")

	for _, val := range game_subset {

		pattern := regexp.MustCompile(`(\d+)\s+([a-zA-Z]+)`)
		matches := pattern.FindAllStringSubmatch(val, -1)

		for _, match := range matches {
			number, _ := strconv.Atoi(match[1])
			color := match[2]

			if (color == "red" && number > 12) ||
				(color == "green" && number > 13) ||
				(color == "blue" && number > 14) {
				return false
			}

		}
	}

	return true
}

func Part2(input string) int {
	game_subset := strings.Split((strings.Split(input, ":")[1]), ";")
	red := 0
	blue := 0
	green := 0
	for _, val := range game_subset {

		pattern := regexp.MustCompile(`(\d+)\s+([a-zA-Z]+)`)
		matches := pattern.FindAllStringSubmatch(val, -1)

		for _, match := range matches {
			number, _ := strconv.Atoi(match[1])
			color := match[2]
			if color == "red" && red < number {
				red = number
			}
			if color == "blue" && blue < number {
				blue = number
			}
			if color == "green" && green < number {
				green = number
			}

		}
	}

	return red * blue * green
}
