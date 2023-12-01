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
  fmt.Println(x)

	for i := 0; i < 26; i++ {
		priorities[ASCIIIntToChar('a'+i)] = i + 1
		priorities[ASCIIIntToChar('A'+i)] = i + 27
	}

  prioritiesSum := 0

  for _, l := range strings.Split(x, "\n") {
    righthalf := l[len(l)/2:]
    if (righthalf)
  }
}

func ASCIIIntToChar(code int) string {
	return string(rune(code))
}
