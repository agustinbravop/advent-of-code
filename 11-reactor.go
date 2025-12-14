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
		devices[device] = strings.Split(outputs, " ")
	}

	successfulPaths := 0
	// Breadth-first search to find all paths from device 'you' to device 'out'.
	// Each possible path is a slice of device names.
	queue := [][]string{{"you"}}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		currentDevice := path[len(path)-1]

		if currentDevice == "out" {
			// Found a successful path.
			successfulPaths += 1
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

	return fmt.Sprintf("%d", successfulPaths)
}

// Recursive depth-first search to find all paths from current device to 'out'.
// Memoization: `cache` stores the number of paths that can be reached from each device.
func findAllPaths(devices map[string][]string, current string, seenDac, seenFft bool, cache map[cacheKey]int) int {
	if current == "out" && seenDac && seenFft {
		// Found a successful path.
		return 1
	}

	if current == "dac" {
		seenDac = true
	} else if current == "fft" {
		seenFft = true
	}

	// Check if the result is already cached.
	cacheKey := cacheKey{device: current, seenDac: seenDac, seenFft: seenFft}
	if val, ok := cache[cacheKey]; ok {
		return val
	}

	// Visit children.
	total := 0
	for _, output := range devices[current] {
		total += findAllPaths(devices, output, seenDac, seenFft, cache)
	}

	cache[cacheKey] = total
	return total
}

type cacheKey struct {
	device           string
	seenDac, seenFft bool
}

func Day11Part2() string {
	text, err := os.ReadFile("11-input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading input file: %v\n", err)
	}

	devices := make(map[string][]string)
	for line := range strings.SplitSeq(string(text), "\n") {
		// Parse devices: add each device with its outputs to the hashmap.
		device, outputs, _ := strings.Cut(line, ": ")
		devices[device] = strings.Split(outputs, " ")
	}

	cache := make(map[cacheKey]int)
	paths := findAllPaths(devices, "svr", false, false, cache)
	return fmt.Sprintf("%d", paths)
}
