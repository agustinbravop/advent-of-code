package main

import (
	"fmt"
	"os"
	"strings"
)

func Day7Part1() string {
	text, err := os.ReadFile("07-input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading input file: %v\n", err)
	}

	lines := strings.Split(string(text), "\n")
	beamStartPosition := strings.Index(lines[0], "S")
	// This set stores active beams. This set is updated as we process every line.
	beams := map[int]struct{}{beamStartPosition: {}}
	splits := 0

	for _, line := range lines {
		for i, char := range line {
			_, hasBeam := beams[i]
			if char == '^' && hasBeam {
				// A beam has collided with a splitter, so we split the beam.
				delete(beams, i)
				beams[i-1] = struct{}{}
				beams[i+1] = struct{}{}
				splits += 1
			}
		}
	}

	return fmt.Sprintf("%d", splits)
}
