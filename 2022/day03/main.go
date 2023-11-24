package main

import (
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	x := strings.TrimRight(string(input), "\n")

	priorities := make(map[string]int)

	for i := 0; i < 26; i++ {
		priorities[ASCIIIntToChar('a'+i)] = i + 1
		priorities[ASCIIIntToChar('A'+i)] = i + 27
	}
}

func ASCIIIntToChar(code int) string {
	return string(rune(code))
}
