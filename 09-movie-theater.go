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

// Check if rectangle is valid: no green lines intersect the interior of the rectangle.
func hasOnlyRedAndGreenTiles(redTiles [][]int, corner1, corner2 []int) bool {
	minX := min(corner1[0], corner2[0])
	maxX := max(corner1[0], corner2[0])
	minY := min(corner1[1], corner2[1])
	maxY := max(corner1[1], corner2[1])

	// Check if any red tiles are strictly inside rectangle (not allowed).
	for _, tile := range redTiles {
		x, y := tile[0], tile[1]
		if minX < x && x < maxX && minY < y && y < maxY {
			if (x == corner1[0] && y == corner1[1]) || (x == corner2[0] && y == corner2[1]) {
				// Corner tiles are allowed.
				continue
			}
			// Other red tiles inside the rectangle are not allowed.
			return false
		}
	}

	// Check if any green line segment intersects the interior of the rectangle.
	n := len(redTiles)
	for i := range n {
		curr := redTiles[i]
		next := redTiles[(i+1)%n] // Handle wrapping.

		if lineSegmentIntersectsRectangle(curr, next, minX, maxX, minY, maxY) {
			return false
		}
	}

	return true
}

// Check if line segment intersects the interior of the rectangle.
func lineSegmentIntersectsRectangle(p1, p2 []int, minX, maxX, minY, maxY int) bool {
	x1, y1 := p1[0], p1[1]
	x2, y2 := p2[0], p2[1]

	// Check if line segment is completely outside rectangle bounds.
	if (x1 < minX && x2 < minX) || (x1 > maxX && x2 > maxX) ||
		(y1 < minY && y2 < minY) || (y1 > maxY && y2 > maxY) {
		return false
	}

	// Check if line segment crosses the interior of the rectangle.
	if x1 == x2 {
		if minX < x1 && x1 < maxX {
			if (y1 <= minY && y2 >= maxY) || (y1 >= maxY && y2 <= minY) {
				// Crosses rectangle horizontally.
				return true
			}
		}
	} else {
		if minY < y1 && y1 < maxY {
			if (x1 <= minX && x2 >= maxX) || (x1 >= maxX && x2 <= minX) {
				// Crosses rectangle vertically.
				return true
			}
		}
	}

	return false
}

func Day9Part2() string {
	text, err := os.ReadFile("09-input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading input file: %v\n", err)
	}

	// Process input file to get the position of all red tiles and green lines between them.
	redTiles := [][]int{}
	for line := range strings.SplitSeq(string(text), "\n") {
		numbers := strings.Split(line, ",")
		x, _ := strconv.Atoi(numbers[0])
		y, _ := strconv.Atoi(numbers[1])
		redTiles = append(redTiles, []int{x, y})
	}

	// Search all combinations of red tiles to find the largest valid rectangle.
	bestArea := 0
	for i := 0; i < len(redTiles)-1; i++ {
		for j := i + 1; j < len(redTiles); j++ {
			corner1 := redTiles[i]
			corner2 := redTiles[j]

			// Only update the area if the rectangle is valid and the area is better.
			if hasOnlyRedAndGreenTiles(redTiles, corner1, corner2) {
				width := max(corner1[0], corner2[0]) - min(corner1[0], corner2[0]) + 1
				height := max(corner1[1], corner2[1]) - min(corner1[1], corner2[1]) + 1
				area := width * height
				if area > bestArea {
					bestArea = area
				}
			}
		}
	}

	return fmt.Sprintf("%d", bestArea)
}
