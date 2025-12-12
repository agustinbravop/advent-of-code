package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day9Part1() string {
	text, err := os.ReadFile("09-input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading input file: %v\n", err)
	}

	// Process input file to get the position of all red tiles.
	redTiles := [][]int{}
	for line := range strings.SplitSeq(string(text), "\n") {
		numbers := strings.Split(line, ",")
		x, _ := strconv.Atoi(numbers[0])
		y, _ := strconv.Atoi(numbers[1])
		redTiles = append(redTiles, []int{x, y})
	}

	// The largest rectangle will either be formed
	// by the furthest bottom-left and top-right red tiles
	// or by the furthest bottom-right and top-left red tiles.
	topLeftMost := redTiles[0]
	topRightMost := redTiles[0]
	bottomLeftMost := redTiles[0]
	bottomRightMost := redTiles[0]

	// Find those four red tiles closest to each corner.
	for _, tile := range redTiles {
		x, y := tile[0], tile[1]

		if x+y <= topLeftMost[0]+topLeftMost[1] {
			topLeftMost = tile
		}
		if -x+y <= -topRightMost[0]+topRightMost[1] {
			topRightMost = tile
		}
		if x-y <= bottomLeftMost[0]-bottomLeftMost[1] {
			bottomLeftMost = tile
		}
		if -x-y <= -bottomRightMost[0]-bottomRightMost[1] {
			bottomRightMost = tile
		}
	}

	// Calculate the area of the largest rectangle.
	firstRectangleArea := (bottomRightMost[0] - topLeftMost[0] + 1) * (bottomRightMost[1] - topLeftMost[1] + 1)
	secondRectangleArea := (topRightMost[0] - bottomLeftMost[0] + 1) * (bottomLeftMost[1] - topRightMost[1] + 1)

	if firstRectangleArea > secondRectangleArea {
		return fmt.Sprintf("%d", firstRectangleArea)
	} else {
		return fmt.Sprintf("%d", secondRectangleArea)
	}
}
