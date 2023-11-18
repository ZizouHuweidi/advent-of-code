package main

import (
	"os"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	// x := strings.Split(string(input), "\n")
}
