package main

import (
	"testing"
)

func Test(t *testing.T) {
	total := Part1(testInput)

	if total != 142 {
		t.Errorf("wrong total: %d", total)
	}
}

const testInput = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
