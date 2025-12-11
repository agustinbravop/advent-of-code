package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func Day3Part1() string {
	text, err := os.ReadFile("03-input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading input file: %v\n", err)
	}

	joltage := 0
	for bank := range strings.SplitSeq(string(text), "\n") {
		// read bank as an array of single digit integers.
		numbers := make([]int, len(bank))
		for i, digit := range bank {
			numbers[i] = int(digit - '0')
		}

		firstDigitIndex := 0
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[firstDigitIndex] {
				firstDigitIndex = i
			}
		}

		secondDigit := 0
		for i := firstDigitIndex + 1; i < len(numbers); i++ {
			if numbers[i] > secondDigit {
				secondDigit = numbers[i]
			}
		}

		joltage += numbers[firstDigitIndex]*10 + secondDigit
	}

	return fmt.Sprintf("%d", joltage)
}

func Day3Part2() string {
	text, err := os.ReadFile("03-input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading input file: %v\n", err)
	}
	joltage := 0
	for bank := range strings.SplitSeq(string(text), "\n") {
		// read bank as an array of single digit integers.
		numbers := make([]int, len(bank))
		for i, digit := range bank {
			numbers[i] = int(digit - '0')
		}

		// Store the position of each selected battery (index of the digit in the numbers array).
		positions := make([]int, 12)
		prevPos := -1
		// For every position, find the index of the largest digit valid for the battery at that position.
		for i := 0; i < len(positions); i++ {
			positions[i] = prevPos + 1
			// Find the largest digit among valid digits for the current position.
			for j := prevPos + 1; j < len(numbers)-len(positions)+i+1; j++ {
				if numbers[j] > numbers[positions[i]] {
					positions[i] = j
				}
			}
			prevPos = positions[i]
		}

		for i, pos := range positions {
			joltage += numbers[pos] * (int(math.Pow10(len(positions) - i - 1)))
		}
	}

	return fmt.Sprintf("%d", joltage)
}
