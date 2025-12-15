package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	indicators []bool
	buttons    [][]int
	joltages   []int
}

func parseLine(line string) Machine {
	endOfLights := strings.Index(line, "]")
	startOfJoltages := strings.Index(line, "{")

	// Parse indicator lights.
	indicators := make([]bool, endOfLights-1)
	for i, char := range line[1:endOfLights] {
		indicators[i] = char == '#'
	}

	// Parse buttons.
	var buttons [][]int
	schematics := strings.SplitSeq(line[endOfLights+3:startOfJoltages-2], ") (")
	for schematic := range schematics {
		var button []int
		for part := range strings.SplitSeq(schematic, ",") {
			if num, err := strconv.Atoi(part); err == nil {
				button = append(button, num)
			}
		}
		buttons = append(buttons, button)
	}

	// Parse joltages.
	var joltages []int
	requirements := strings.SplitSeq(line[startOfJoltages+1:len(line)-1], ",")
	for req := range requirements {
		if num, err := strconv.Atoi(req); err == nil {
			joltages = append(joltages, num)
		}
	}

	return Machine{indicators, buttons, joltages}
}

func boolsToString(lights []bool) string {
	result := make([]byte, len(lights))
	for i, light := range lights {
		if light {
			result[i] = '1'
		} else {
			result[i] = '0'
		}
	}
	return string(result)
}

func intsToString(nums []int) string {
	return fmt.Sprintf("%v", nums)
}

func generatePatterns(buttons [][]int, numLights int) map[string][][]int {
	patterns := make(map[string][][]int)

	for mask := 0; mask < (1 << len(buttons)); mask++ {
		lights := make([]bool, numLights)
		var pressed []int

		for i, button := range buttons {
			if mask&(1<<i) != 0 {
				pressed = append(pressed, i)
				for _, pos := range button {
					lights[pos] = !lights[pos]
				}
			}
		}

		key := boolsToString(lights)
		patterns[key] = append(patterns[key], pressed)
	}

	return patterns
}

func minButtonPressesForJoltages(joltages []int, patterns map[string][][]int, buttons [][]int, cache map[string]int) int {
	// Base case: all zeros.
	if allZeros(joltages) {
		return 0
	}

	// Check cache.
	key := intsToString(joltages)
	if result, exists := cache[key]; exists {
		return result
	}

	// Find indicator pattern for odd joltages.
	indicators := make([]bool, len(joltages))
	for i, j := range joltages {
		indicators[i] = j%2 == 1
	}

	indicatorsKey := boolsToString(indicators)
	buttonCombos, exists := patterns[indicatorsKey]
	if !exists {
		cache[key] = -1
		return -1
	}

	minPresses := -1

	// Try each button combination.
	for _, pressed := range buttonCombos {
		targetAfter := make([]int, len(joltages))
		copy(targetAfter, joltages)

		// Apply button presses.
		for _, btnIdx := range pressed {
			for _, joltIdx := range buttons[btnIdx] {
				targetAfter[joltIdx]--
			}
		}

		// Skip negative values.
		if anyNegative(targetAfter) {
			continue
		}

		// Calculate half target (all values are even).
		halfTarget := make([]int, len(targetAfter))
		for i, val := range targetAfter {
			halfTarget[i] = val / 2
		}

		halfPresses := minButtonPressesForJoltages(halfTarget, patterns, buttons, cache)
		if halfPresses == -1 {
			continue
		}

		totalPresses := len(pressed) + 2*halfPresses
		if minPresses == -1 || totalPresses < minPresses {
			minPresses = totalPresses
		}
	}

	cache[key] = minPresses
	return minPresses
}

func allZeros(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

func anyNegative(nums []int) bool {
	for _, num := range nums {
		if num < 0 {
			return true
		}
	}
	return false
}

func Day10Part1() string {
	content, err := os.ReadFile("10-input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading file: %v", err)
	}

	total := 0
	for line := range strings.SplitSeq(string(content), "\n") {
		if line == "" {
			continue
		}
		machine := parseLine(line)
		patterns := generatePatterns(machine.buttons, len(machine.indicators))

		indicatorsKey := boolsToString(machine.indicators)
		if combos, exists := patterns[indicatorsKey]; exists {
			minPresses := len(combos[0])
			for _, combo := range combos[1:] {
				if len(combo) < minPresses {
					minPresses = len(combo)
				}
			}
			total += minPresses
		}
	}

	return fmt.Sprintf("%d", total)
}

func Day10Part2() string {
	content, err := os.ReadFile("10-input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading file: %v", err)
	}

	total := 0
	for line := range strings.SplitSeq(string(content), "\n") {
		if line == "" {
			continue
		}
		machine := parseLine(line)
		patterns := generatePatterns(machine.buttons, len(machine.joltages))
		cache := make(map[string]int)

		result := minButtonPressesForJoltages(machine.joltages, patterns, machine.buttons, cache)
		if result != -1 {
			total += result
		}
	}

	return fmt.Sprintf("%d", total)
}
