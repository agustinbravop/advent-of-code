package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day6Part1() string {
	text, err := os.ReadFile("06-input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading input file: %v\n", err)
	}

	lines := strings.Split(string(text), "\n")

	// For each problem we store the numbers as a slice of integers.
	problems := make([][]int, len(strings.Fields(lines[0])))

	// Read the problems (all numbers but not the operators at the end).
	for i := 0; i < len(lines)-1; i++ {
		// Split each line by whitespace to get one number for each problem.
		nums := strings.Fields(lines[i])
		for j, num := range nums {
			numInt, err := strconv.Atoi(num)
			if err != nil {
				return fmt.Sprintf("Error converting string to int: %v\n", err)
			}
			problems[j] = append(problems[j], numInt)
		}
	}

	// Read each operator and calculate the answer to each problem.
	total := 0
	ops := strings.Fields(lines[len(lines)-1])
	for j, op := range ops {
		answer := 0
		switch op {
		case "+":
			for i := 0; i < len(problems[j]); i++ {
				answer += problems[j][i]
			}
		case "*":
			answer = 1
			for i := 0; i < len(problems[j]); i++ {
				answer *= problems[j][i]
			}
		}
		total += answer
	}

	return fmt.Sprintf("%d", total)
}

// getLargestLength returns the length of the longest string in the slice.
func getLargestLength(slice []string) int {
	if len(slice) == 0 {
		return 0
	}

	max := len(slice[0])
	for _, str := range slice {
		if len(str) > max {
			max = len(str)
		}
	}
	return max
}

func Day6Part2() string {
	text, err := os.ReadFile("06-input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading input file: %v\n", err)
	}

	lines := strings.Split(string(text), "\n")

	// For each problem we store the numbers as a slice of strings.
	problems := make([][]string, len(strings.Fields(lines[0])))
	// The positions of the operators of the last line help separate numbers into columns.
	separators := []int{}
	for i, op := range lines[len(lines)-1] {
		if op != ' ' {
			separators = append(separators, i)
		}
	}

	// Read the problems (all numbers but not the operators at the end).
	for i := 0; i < len(lines)-1; i++ {
		// Split each line by whitespace to get one number for each problem.
		nums := []string{}
		num := ""
		for k := 0; k < len(lines[i]); k++ {
			if slices.Contains(separators, k+1) {
				// The current number ended one character ago, and a new number starts in the next character.
				nums = append(nums, num)
				num = ""
			} else {
				num += string(lines[i][k])
			}
		}
		// Add the last number in the line.
		nums = append(nums, num)

		for j, num := range nums {
			problems[j] = append(problems[j], num)
		}
	}

	// Read each operator and calculate the answer to each problem.
	total := 0
	ops := strings.Fields(lines[len(lines)-1])
	for j, op := range ops {
		answer := 0
		largestLength := getLargestLength(problems[j])

		switch op {
		case "+":
			// Iterate over each digit position.
			for d := range largestLength {
				// Get the number by concatenating digits from top to bottom.
				number := ""
				for i := 0; i < len(problems[j]); i++ {
					if problems[j][i][d] != ' ' {
						number += string(problems[j][i][d])
					}
				}
				integer, _ := strconv.Atoi(number)
				answer += integer
			}
		case "*":
			answer = 1
			// Iterate over each digit position.
			for d := range largestLength {
				// Get the number by concatenating digits from top to bottom.
				number := ""
				for i := 0; i < len(problems[j]); i++ {
					if problems[j][i][d] != ' ' {
						number += string(problems[j][i][d])
					}
				}
				integer, _ := strconv.Atoi(number)
				answer *= integer
			}
		}

		total += answer
	}

	return fmt.Sprintf("%d", total)
}
