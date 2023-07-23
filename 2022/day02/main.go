package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	games := strings.Split(string(content), "\n")

}
