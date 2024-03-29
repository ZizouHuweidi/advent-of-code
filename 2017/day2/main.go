package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	checksumPart1 := 0
	checksumPart2 := 0

	for scanner.Scan() {
		line := scanner.Text()

		checkSum := Part1(line)
		checksumPart1 += checkSum

		checkSum = Part2(line)
		checksumPart2 += checkSum
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("checksum of part 1: %d\n", checksumPart1)
	fmt.Printf("checksum of part 2: %d\n", checksumPart2)
}

func Part1(line string) int {
	intArr := make([]int, len(strings.Fields(line)))

	for i, v := range strings.Fields(line) {
		num, _ := strconv.Atoi(v)
		intArr[i] = num
	}

	mn, mx := slices.Min(intArr), slices.Max(intArr)

	res := mx - mn

	return res
}

func Part2(line string) int {
	intArr := make([]int, len(strings.Fields(line)))

	for i, v := range strings.Fields(line) {
		num, _ := strconv.Atoi(v)
		intArr[i] = num
	}

	for i := 0; i < len(intArr); i++ {
		for j := i + 1; j < len(intArr); j++ {
			if intArr[i]%intArr[j] == 0 {
				return intArr[i] / intArr[j]
			} else if intArr[j]%intArr[i] == 0 {
				return intArr[j] / intArr[i]
			}
		}
	}

	return 0
}
