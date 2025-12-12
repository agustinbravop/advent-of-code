package main

import (
	"fmt"
	"os"
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
