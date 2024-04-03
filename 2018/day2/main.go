package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	countTwo := 0
	countThree := 0

	var boxIDs []string

	for scanner.Scan() {
		line := scanner.Text()

		hasTwo, hasThree := countOccurrences(line)
		if hasTwo {
			countTwo++
		}
		if hasThree {
			countThree++
		}

		boxIDs = append(boxIDs, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	checksum := countTwo * countThree
	commonLetters := findCommonLetters(boxIDs)
	fmt.Printf("Checksum: %d\n", checksum)
	fmt.Println("Common letters between the correct box IDs:", commonLetters)
}

func countOccurrences(input string) (bool, bool) {
	counts := make(map[rune]int)

	for _, letter := range input {
		counts[letter]++
	}

	hasTwo := false
	hasThree := false

	for _, count := range counts {
		if count == 2 {
			hasTwo = true
		}
		if count == 3 {
			hasThree = true
		}
	}

	return hasTwo, hasThree
}

func findCommonLetters(boxIDs []string) string {
	for i := 0; i < len(boxIDs); i++ {
		for j := i + 1; j < len(boxIDs); j++ {
			commonLetters := compareBoxIDs(boxIDs[i], boxIDs[j])
			if commonLetters != "" {
				return commonLetters
			}
		}
	}
	return ""
}

func compareBoxIDs(id1, id2 string) string {
	differenceCount := 0
	var commonLetters []rune

	for i := 0; i < len(id1); i++ {
		if id1[i] != id2[i] {
			differenceCount++
		} else {
			commonLetters = append(commonLetters, rune(id1[i]))
		}
	}

	if differenceCount == 1 {
		return string(commonLetters)
	}
	return ""
}
