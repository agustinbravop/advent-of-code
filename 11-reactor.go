package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func Day11Part1() string {
	text, err := os.ReadFile("11-input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading input file: %v\n", err)
	}

	devices := make(map[string][]string)
	for line := range strings.SplitSeq(string(text), "\n") {
		// Parse devices: add each device and its outputs to the hashmap.
		device, outputs, _ := strings.Cut(line, ": ")
		for output := range strings.SplitSeq(outputs, " ") {
			devices[device] = append(devices[device], output)
		}
	}

	successfulPaths := make([][]string, 0)
	// Breadth-first search to find all paths from device 'you' to device 'out'.
	// Each possible path is a slice of device names.
	queue := [][]string{{"you"}}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		currentDevice := path[len(path)-1]

		if currentDevice == "out" {
			// Found a successful path.
			successfulPaths = append(successfulPaths, path)
		}

		for _, output := range devices[currentDevice] {
			if !slices.Contains(path, output) {
				// This path has not passed through this device before, so visit it.
				// Do an immutable append to avoid affecting the Contains() check.
				newPath := make([]string, len(path)+1)
				copy(newPath, path)
				newPath[len(newPath)-1] = output
				queue = append(queue, newPath)
			}
		}
	}

	return fmt.Sprintf("%v", len(successfulPaths))
}
