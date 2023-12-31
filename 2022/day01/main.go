package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	calories := strings.Split(string(input), "\n")

	var maxCalories []int

	currentCalories := 0

	for _, calorie := range calories {
		if calorie == "" {
			maxCalories = append(maxCalories, currentCalories)
			currentCalories = 0
			continue
		}
		calorieInt, _ := strconv.Atoi(calorie)
		currentCalories = currentCalories + calorieInt
	}
	maxCalories = append(maxCalories, currentCalories)

	sort.Ints(maxCalories)

	part1 := maxCalories[len(maxCalories)-1]
	part2 := maxCalories[len(maxCalories)-1] + maxCalories[len(maxCalories)-2] + maxCalories[len(maxCalories)-3]

	fmt.Println(part1)
	fmt.Println(part2)

}
