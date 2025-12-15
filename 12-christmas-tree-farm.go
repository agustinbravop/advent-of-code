package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 3x3 shapes.
type shape [3][3]bool

type region struct {
	width, length int
	// Map of shape index to quantity.
	presents map[int]int
}

// Rotates a 3x3 shape 90 degrees clockwise.
func rotateShape(s shape) shape {
	rotated := shape{}
	for i := range 3 {
		for j := range 3 {
			rotated[i][j] = s[2-j][i]
		}
	}
	return rotated
}

// Flips a 3x3 shape horizontally.
func flipShape(s shape) shape {
	flipped := shape{}
	flipped[0] = s[2]
	flipped[1] = s[1]
	flipped[2] = s[0]
	return flipped
}

// Returns true if the given region can fit all its presents inside.
func canRegionFitPresents(region region, shapes map[int]shape) bool {
	area := region.width * region.length
	requiredArea := 0
	totalPresents := 0
	for i, quantity := range region.presents {
		minimumSize := 0
		for _, row := range shapes[i] {
			for _, cell := range row {
				if cell {
					minimumSize += 1
				}
			}
		}
		requiredArea += minimumSize * quantity
		totalPresents += quantity
	}

	// Check if the region is too small to fit all presents.
	if area <= requiredArea {
		return false
	}

	maxPresentsInWidth := region.width / 3
	maxPresentsInLength := region.length / 3

	// Check if the region is so large that packing is not an issue.
	if maxPresentsInWidth*maxPresentsInLength >= totalPresents {
		return true
	}

	// The two previous checks seem to be good enough heuristics for the given input.
	// TODO: implement a general case solution.
	panic("general case shape packing solution not implemented")
}

func Day12Part1() string {
	text, err := os.ReadFile("12-input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading input file: %v\n", err)
	}

	sections := strings.Split(string(text), "\n\n")

	// Parse all shapes.
	shapes := make(map[int]shape)
	for i := 0; i < len(sections)-1; i++ {
		lines := strings.Split(sections[i], "\n")
		shapeIndexStr, _ := strings.CutSuffix(lines[0], ":")
		shapeIndex, _ := strconv.Atoi(shapeIndexStr)

		// Parse one shape.
		shape := shape{}
		for j := 1; j < len(lines); j++ {
			for k := 0; k < len(lines[j]); k++ {
				shape[j-1][k] = lines[j][k] == '#'
			}
		}
		shapes[shapeIndex] = shape
	}

	validRegionsCount := 0

	// Parse and process regions.
	lastSection := sections[len(sections)-1]
	for line := range strings.SplitSeq(lastSection, "\n") {
		size, presents, _ := strings.Cut(line, ": ")
		width, length, _ := strings.Cut(size, "x")
		w, _ := strconv.Atoi(width)
		l, _ := strconv.Atoi(length)

		// Each region has a map of shape index to quantity.
		// This saves memory as all shapes are only stored once.
		p := make(map[int]int)
		for i, quantity := range strings.Split(presents, " ") {
			q, _ := strconv.Atoi(quantity)
			p[i] = q
		}

		region := region{width: w, length: l, presents: p}
		if canRegionFitPresents(region, shapes) {
			validRegionsCount += 1
		}
	}

	return fmt.Sprintf("%d", validRegionsCount)
}
