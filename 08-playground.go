package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type JunctionBox struct {
	x, y, z, circuit int
}

// Returns the squared Euclidean distance between two junction boxes.
// The square root is omitted given that it doesn't change the order.
func squaredDistance(j1, j2 JunctionBox) int {
	return (j1.x-j2.x)*(j1.x-j2.x) + (j1.y-j2.y)*(j1.y-j2.y) + (j1.z-j2.z)*(j1.z-j2.z)
}

// Finds the indexes of the two closest junction boxes that are not directly connected.
func findClosestUnconnectedPair(junctionBoxes []JunctionBox, directlyConnected map[string]struct{}) (int, int) {
	j1, j2 := 0, 1
	minDistance := squaredDistance(junctionBoxes[0], junctionBoxes[1])
	for i := 0; i < len(junctionBoxes); i++ {
		for j := i + 1; j < len(junctionBoxes); j++ {
			distance := squaredDistance(junctionBoxes[i], junctionBoxes[j])
			if distance < minDistance {
				// Check that the junction boxes aren't already directly connected.
				key := fmt.Sprintf("%d-%d", i, j)
				_, ok := directlyConnected[key]
				if !ok {
					minDistance = distance
					j1, j2 = i, j
				}
			}
		}
	}

	return j1, j2
}

func Day8Part1() string {
	text, err := os.ReadFile("08-input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading input file: %v\n", err)
	}

	// Parse entire input file for easy processing in a later stage.
	junctionBoxes := []JunctionBox{}
	for i, line := range strings.Split(string(text), "\n") {
		positions := strings.Split(line, ",")
		x, _ := strconv.Atoi(positions[0])
		y, _ := strconv.Atoi(positions[1])
		z, _ := strconv.Atoi(positions[2])
		// At the start, each junction box belongs to its own isolated circuit.
		junctionBoxes = append(junctionBoxes, JunctionBox{x, y, z, i})
	}

	// Keep track of directly connected junction boxes so they're not connected again.
	directlyConnected := make(map[string]struct{})
	for range 10 {
		j1, j2 := findClosestUnconnectedPair(junctionBoxes, directlyConnected)

		// Connect the two closest junction boxes.
		key := fmt.Sprintf("%d-%d", j1, j2)
		directlyConnected[key] = struct{}{}
		deletedCircuit := junctionBoxes[j2].circuit
		for i := range junctionBoxes {
			if junctionBoxes[i].circuit == deletedCircuit {
				junctionBoxes[i].circuit = junctionBoxes[j1].circuit
			}
		}
	}

	// Count the amout of junction boxes in each circuit.
	circuitSizes := make(map[int]int)
	for _, junctionBox := range junctionBoxes {
		circuitSizes[junctionBox.circuit]++
	}

	// Find the three largest circuits.
	var largestSizes []int
	for _, size := range circuitSizes {
		if len(largestSizes) < 3 {
			largestSizes = append(largestSizes, size)
			sort.Ints(largestSizes)
		} else if size > largestSizes[0] {
			largestSizes[0] = size
			sort.Ints(largestSizes)
		}
	}

	answer := largestSizes[0] * largestSizes[1] * largestSizes[2]
	return fmt.Sprintf("%d", answer)
}

func Day8Part2() string {
	text, err := os.ReadFile("08-input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading input file: %v\n", err)
	}

	// Parse entire input file for easy processing in a later stage.
	junctionBoxes := []JunctionBox{}
	for i, line := range strings.Split(string(text), "\n") {
		positions := strings.Split(line, ",")
		x, _ := strconv.Atoi(positions[0])
		y, _ := strconv.Atoi(positions[1])
		z, _ := strconv.Atoi(positions[2])
		// At the start, each junction box belongs to its own isolated circuit.
		junctionBoxes = append(junctionBoxes, JunctionBox{x, y, z, i})
	}

	// Keep track of directly connected junction boxes so they're not connected again.
	directlyConnected := make(map[string]struct{})
	for true {
		j1, j2 := findClosestUnconnectedPair(junctionBoxes, directlyConnected)

		// Connect the two closest junction boxes.
		key := fmt.Sprintf("%d-%d", j1, j2)
		directlyConnected[key] = struct{}{}
		deletedCircuit := junctionBoxes[j2].circuit
		onlyOneCircuitRemains := true
		for i := range junctionBoxes {
			if junctionBoxes[i].circuit == deletedCircuit {
				junctionBoxes[i].circuit = junctionBoxes[j1].circuit
			} else if junctionBoxes[i].circuit != junctionBoxes[j1].circuit {
				// Check if there's only one circuit left.
				onlyOneCircuitRemains = false
			}
		}

		if onlyOneCircuitRemains {
			answer := junctionBoxes[j1].x * junctionBoxes[j2].x
			return fmt.Sprintf("%d", answer)
		}
	}
	return "No solution found"
}
