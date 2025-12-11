package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// isIngredientFresh returns true if the ingredient with the given ID is included
// in at least one of the provided ranges of fresh ingredients.
func isIngredientFresh(id int, ranges [][]int) bool {
	for _, r := range ranges {
		if id >= r[0] && id <= r[1] {
			return true
		}
	}
	return false
}

func Day5Part1() {
	text, err := os.ReadFile("05-input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	// We split the file first and then process ingredients line by line.
	// Split the file by the empty line to get both halves.
	halves := strings.Split(string(text), "\n\n")
	freshIngredientRangesText := halves[0]
	ingredientIdsText := halves[1]

	var freshIngredientRanges [][]int
	for line := range strings.SplitSeq(freshIngredientRangesText, "\n") {
		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		freshIngredientRanges = append(freshIngredientRanges, []int{start, end})
	}

	freshIngredientCount := 0
	for line := range strings.SplitSeq(ingredientIdsText, "\n") {
		id, _ := strconv.Atoi(line)
		if isIngredientFresh(id, freshIngredientRanges) {
			freshIngredientCount += 1
		}
	}

	fmt.Println(freshIngredientCount)
}
