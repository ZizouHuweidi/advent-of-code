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
	arr := strings.Split(input, "\n")

	var numArr []int
	for _, v := range arr {
		numArr = append(numArr, cast.ToInt(v))
	}

	fmt.Println(Part1(numArr))
	fmt.Println(Part2(numArr))
}

func Part1(numArr []int) int {
	sum := 0

	for _, v := range numArr {
		sum += int(v/3) - 2
	}

	return sum
}

func Part2(numArr []int) int {
	sum := 0

	for _, v := range numArr {
		sum += calculateFuel(v)
	}

	return sum
}

func calculateFuel(mass int) int {
	fuel := (mass / 3) - 2
	if fuel <= 0 {
		return 0
	}
	return fuel + calculateFuel(fuel)
}
