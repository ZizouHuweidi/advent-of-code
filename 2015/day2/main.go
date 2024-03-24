package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zizouhuweidi/advent-of-code/pkg/cast"
	"github.com/zizouhuweidi/advent-of-code/pkg/util"
)

func main() {
	s := util.ReadFile(os.Args[1])

	input := parseInput(s)

	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}

func Part1(input [][]int) int {
	sum := 0

	for _, v := range input {
		l, w, h := v[0], v[1], v[2]

		area := 2*l*w + 2*w*h + 2*h*l

		sum += area + min(l*w, w*h, h*l)
	}

	return sum
}

func Part2(input [][]int) int {
	sum := 0

	for _, v := range input {
		l, w, h := v[0], v[1], v[2]

		perimeter := min(2*(l+w), 2*(w+h), 2*(h+l))

		bow := l * w * h

		sum += perimeter + bow
	}

	return sum
}

func parseInput(input string) [][]int {
	res := make([][]int, 0)

	tmp := strings.Split(input, "\n")

	for _, v := range tmp {
		dimensions := strings.Split(v, "x")
		d := make([]int, len(dimensions))
		for i, str := range dimensions {
			d[i] = cast.ToInt(str)
		}
		res = append(res, d)
	}

	return res
}
