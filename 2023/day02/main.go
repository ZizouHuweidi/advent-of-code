package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/zizouhuweidi/advent-of-code/utils"
)

func main() {
	data := strings.Split(strings.Trim(utils.GetInput("input.txt"), "\n"), "\n")

	validGames := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	maxCubes := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	var sum1 int
	var sum2 int

	for _, game := range data {
		re := regexp.MustCompile(`Game (\d+): (.+)`)
		match := re.FindStringSubmatch(game)
		if match == nil {
			continue
		}

		gameID, err := strconv.Atoi(match[1])
		if err != nil {
			continue
		}

		cubesSequence := match[2]

		if isValidGame(cubesSequence, validGames) {
			sum1 += gameID
		}
	}

	fmt.Println(sum1)
}

func isValidGame(cubesSequence string, validConfig map[string]int) bool {
	re := regexp.MustCompile(`(\d+) (\w+)`)
	matches := re.FindAllStringSubmatch(cubesSequence, -1)

	for _, match := range matches {
		count, err := strconv.Atoi(match[1])
		if err != nil {
			return false
		}

		color := match[2]

		if validCount, ok := validConfig[color]; ok {
			if count > validCount {
				return false
			}
		}
	}
	return true
}

func gamePower(cubesSequence string, maxCubes map[string]int) int {
	re := regexp.MustCompile(`(\d+) (\w+)`)
	matches := re.FindAllStringSubmatch(cubesSequence, -1)

	for _, match := range matches {
		count, _ := strconv.Atoi(match[1])

		color := match[2]

		currentCount := maxCubes[color]

	}
}
