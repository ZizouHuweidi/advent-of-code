package main

import (
	"os"
	"strings"

	"github.com/zizouhuweidi/advent-of-code/pkg/cast"
	"github.com/zizouhuweidi/advent-of-code/pkg/util"
)

func main() {
	input := util.ReadFile(os.Args[1])
	arr := strings.Split(input, ",")

	var numArr []int
	for _, v := range arr {
		numArr = append(numArr, cast.ToInt(v))
	}

	println(Part1(numArr))
}

func Part1(input []int) int {
	for i := 0; i < len(input); i += 4 {
		opcode := input[i]
		if opcode == 99 {
			break
		}
		switch opcode {
		case 1:
			input[input[i+3]] = input[input[i+1]] + input[input[i+2]]
		case 2:
			input[input[i+3]] = input[input[i+1]] * input[input[i+2]]
		default:
			continue
		}
	}

	return input[0]
}
