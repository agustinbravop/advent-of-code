package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
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

func Day5Part1() string {
	text, err := os.ReadFile("05-input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading input file: %v\n", err)
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

	return fmt.Sprintf("%d", freshIngredientCount)
}

func Day5Part2() string {
	text, err := os.ReadFile("05-input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading input file: %v\n", err)
	}

	// We split the file and only keep the first half with the ID ranges.
	halves := strings.Split(string(text), "\n\n")
	freshIngredientRangesText := halves[0]

	var freshIngredientRanges [][]int
	for line := range strings.SplitSeq(freshIngredientRangesText, "\n") {
		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		// Iterate over existing ingredient ranges to avoid range overlap.
		for i := 0; i < len(freshIngredientRanges); i++ {
			r := freshIngredientRanges[i]
			if r[0] <= start && start <= r[1] {
				// Existing range overlaps with new range, so we clip the new range.
				start = r[1] + 1
			}
			if r[0] <= end && end <= r[1] {
				// Existing range overlaps with new range, so we clip the new range.
				end = r[0] - 1
			}
			if start <= r[0] && r[1] <= end {
				// New range covers completely the existing range, so we remove the existing range.
				freshIngredientRanges = slices.Delete(freshIngredientRanges, i, i+1)
				// Reduce the index by one to account for the removed element.
				i -= 1
			}
		}

		// Insert new range only if its not completely covered by an existing range.
		if start <= end {
			freshIngredientRanges = append(freshIngredientRanges, []int{start, end})
		}
		// The list is sorted to ensure that the ranges are iterated in ascending order.
		sort.Slice(freshIngredientRanges, func(i, j int) bool {
			return freshIngredientRanges[i][0] < freshIngredientRanges[j][0]
		})
	}

	// Sum the lengths of the ranges.
	sum := 0
	for _, r := range freshIngredientRanges {
		sum += r[1] - r[0] + 1
	}

	return fmt.Sprintf("%d", sum)
}
