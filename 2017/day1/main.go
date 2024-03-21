package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zizouhuweidi/advent-of-code/pkg/cast"
	"github.com/zizouhuweidi/advent-of-code/pkg/util"
)

func main() {
	input := util.ReadFile(os.Args[1])

	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}

func Part1(input string) int {
	sum := 0
	digits := parseInput(input)

	for i := 0; i < len(digits); i++ {
		if digits[i] == digits[(i+1)%len(digits)] {
			sum += digits[i]
		}
	}

	return sum
}

func Part2(input string) int {
	digits := parseInput(input)
	var sum int
	offset := len(digits) / 2

	for i := 0; i < len(digits); i++ {
		if digits[i] == digits[(i+offset)%len(digits)] {
			sum += digits[i]
		}
	}

	return sum
}

func parseInput(input string) (res []int) {
	for _, num := range strings.Split(input, "") {
		res = append(res, cast.ToInt(num))
	}
	return res
}
