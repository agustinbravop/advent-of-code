package main

import (
	"fmt"
	"os"
	"strings"
)

func Day4Part1() {
	text, err := os.ReadFile("04-input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	accessiblePaperRolls := 0
	// Append a fake line filled with dots at the start and end of the diagram.
	// This pads the diagram, allowing us to use a sliding window of 3 lines.
	lines := strings.Split(string(text), "\n")
	dotLine := strings.Repeat(".", len(lines[0]))
	lines = append([]string{dotLine}, lines...)
	lines = append(lines, dotLine)

	// Use a sliding window of three lines to check all neighbours of every cell.
	for i := 1; i < len(lines)-1; i++ {
		// For every cell in the line, count it's neighbours.
		for j := 0; j < len(lines[i]); j++ {
			// Ignore cells that do not contain a paper roll.
			if lines[i][j] != '@' {
				continue
			}

			count := 0
			if j > 0 {
				// Count left-side neighbours.
				count += strings.Count(lines[i-1][j-1:j], "@")
				count += strings.Count(lines[i][j-1:j], "@")
				count += strings.Count(lines[i+1][j-1:j], "@")
			}
			if j < len(lines[i])-1 {
				// Count right-side neighbours.
				count += strings.Count(lines[i-1][j+1:j+2], "@")
				count += strings.Count(lines[i][j+1:j+2], "@")
				count += strings.Count(lines[i+1][j+1:j+2], "@")
			}
			count += strings.Count(lines[i-1][j:j+1], "@")
			count += strings.Count(lines[i+1][j:j+1], "@")

			if count < 4 {
				accessiblePaperRolls += 1
			}
		}
	}

	fmt.Println(accessiblePaperRolls)
}

// findAccessiblePaperRolls finds all accessible paper rolls in a given grid.
func findAccessiblePaperRolls(lines []string) [][]int {
	var accessiblePaperRolls [][]int
	// Use a sliding window of three lines to check all neighbours of every cell.
	for i := 1; i < len(lines)-1; i++ {
		// For every cell in the line, count it's neighbours.
		for j := 0; j < len(lines[i]); j++ {
			// Ignore cells that do not contain a paper roll.
			if lines[i][j] != '@' {
				continue
			}

			count := 0
			if j > 0 {
				// Count left-side neighbours.
				count += strings.Count(lines[i-1][j-1:j], "@")
				count += strings.Count(lines[i][j-1:j], "@")
				count += strings.Count(lines[i+1][j-1:j], "@")
			}
			if j < len(lines[i])-1 {
				// Count right-side neighbours.
				count += strings.Count(lines[i-1][j+1:j+2], "@")
				count += strings.Count(lines[i][j+1:j+2], "@")
				count += strings.Count(lines[i+1][j+1:j+2], "@")
			}
			count += strings.Count(lines[i-1][j:j+1], "@")
			count += strings.Count(lines[i+1][j:j+1], "@")

			if count < 4 {
				accessiblePaperRolls = append(accessiblePaperRolls, []int{i, j})
			}
		}
	}

	return accessiblePaperRolls
}

func Day4Part2() {
	text, err := os.ReadFile("04-input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	// Append a fake line filled with dots at the start and end of the diagram.
	// This pads the diagram, allowing us to use a sliding window of 3 lines.
	lines := strings.Split(string(text), "\n")
	dotLine := strings.Repeat(".", len(lines[0]))
	lines = append([]string{dotLine}, lines...)
	lines = append(lines, dotLine)

	removedPaperRolls := 0
	accessiblePaperRollsLocations := findAccessiblePaperRolls(lines)
	for len(accessiblePaperRollsLocations) > 0 {
		// Replace the locations with a dot to remove the accesible paper rolls.
		for _, location := range accessiblePaperRollsLocations {
			lines[location[0]] = lines[location[0]][:location[1]] + "." + lines[location[0]][location[1]+1:]
		}
		removedPaperRolls += len(accessiblePaperRollsLocations)
		// Find the new accessible paper rolls.
		accessiblePaperRollsLocations = findAccessiblePaperRolls(lines)
	}

	fmt.Println(removedPaperRolls)
}
