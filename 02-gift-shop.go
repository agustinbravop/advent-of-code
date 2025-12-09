package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// isInvalid returns true if the ID is made only of some sequence of digits repeated twice.
// Examples of invalid IDs: 55 (5 twice), 6464 (64 twice), and 123123 (123 twice).
func isInvalid(id string) bool {
	// split id by the middle
	middle := len(id) / 2
	firstHalf := id[:middle]
	secondHalf := id[middle:]

	return firstHalf == secondHalf
}

func Day2Part1() {
	text, err := os.ReadFile("02-input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	invalid_id_sum := 0
	lines := strings.SplitSeq(string(text), ",")
	for line := range lines {
		// assume each line is two numbers divided by a '-'
		parts := strings.Split(line, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("Error parsing start number: %v\n", err)
			continue
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("Error parsing end number: %v\n", err)
			continue
		}

		for i := start; i <= end; i++ {
			if isInvalid(strconv.Itoa(i)) {
				invalid_id_sum += i
				fmt.Println(i)
			}
		}
		fmt.Println(start, end, invalid_id_sum)
	}

	fmt.Println(invalid_id_sum)
}
