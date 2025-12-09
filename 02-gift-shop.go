package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// isInvalidPart1 returns true if the ID is made only of some sequence of digits repeated twice.
// Examples of invalid IDs: 55 (5 twice), 6464 (64 twice), and 123123 (123 twice).
func isInvalidPart1(id string) bool {
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
			if isInvalidPart1(strconv.Itoa(i)) {
				invalid_id_sum += i
			}
		}
	}

	fmt.Println(invalid_id_sum)
}

// factors returns the list of factors of a given number.
// Ignore the full length factor.
func factors(n int) []int {
	var factors []int
	for i := 1; i < n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

// isInvalidPart2 returns true if the ID is made only of some sequence of digits repeated at least twice.
// Examples of invalid IDs: 123123123 (123 three times) and 1111111 (1 seven times).
func isInvalidPart2(id string) bool {
	// factorial of the length of the string
	factors := factors(len(id))
	for _, factor := range factors {
		// Extract the potential repeating pattern
		pattern := id[:factor]

		// Build the repeated string and compare
		repeated := strings.Repeat(pattern, len(id)/factor)
		if repeated == id {
			return true
		}
	}
	return false
}

func Day2Part2() {
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
			if isInvalidPart2(strconv.Itoa(i)) {
				invalid_id_sum += i
			}
		}
	}

	fmt.Println(invalid_id_sum)
}
