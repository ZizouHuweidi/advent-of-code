package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	output := Part1(string(input))
	fmt.Println(output)

	output = Part2(string(input))
	fmt.Println(output)
}

func getCalibration(str string) int64 {
	var res string

	var first string
	for _, s := range str {
		if unicode.IsDigit(s) {
			first = string(s)
			break
		}
	}

	var last string
	for i := len(str) - 1; i > 0; i-- {
		if unicode.IsDigit(rune(str[i])) {
			last = string(str[i])
			break
		}
	}

	res = first + last
	ans, _ := strconv.ParseInt(res, 10, 64)

	return ans

}

func convertInput(input string) string {
	for _, line := range strings.Split(input, "\n") {
		for key := range valid_num_strings {

		}
	}
	return input
}

func Part1(input string) int64 {
	var total int64
	for _, line := range strings.Split(input, "\n") {
		total += getCalibration(line)
	}
	return total
}

var valid_num_strings = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func Part2(input string) int64 {
	var res int64
	inp := convertInput(input)
	println(inp)
	res = Part1(inp)

	return res
}
