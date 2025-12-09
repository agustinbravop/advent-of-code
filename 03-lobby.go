package main

import (
	"fmt"
	"os"
	"strings"
)

func Day3Part1() {
	text, err := os.ReadFile("03-input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
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

	fmt.Println(joltage)
}
