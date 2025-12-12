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

// Connects the two closest junction boxes.
// Both their circuits become the same circuit, which also requires updating
// the circuit field of each junction box belonging to the second circuit.
func connectOnePair(junctionBoxes []JunctionBox, directlyConnected map[string]struct{}) ([]JunctionBox, map[string]struct{}) {
	// Find the closest pair of junction boxes.
	minDistance := squaredDistance(junctionBoxes[0], junctionBoxes[1])
	j1, j2 := 0, 1
	for i := 1; i < len(junctionBoxes); i++ {
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

	// Connect the two closest junction boxes.
	key := fmt.Sprintf("%d-%d", j1, j2)
	directlyConnected[key] = struct{}{}
	deletedCircuit := junctionBoxes[j2].circuit
	for i := range junctionBoxes {
		if junctionBoxes[i].circuit == deletedCircuit {
			junctionBoxes[i].circuit = junctionBoxes[j1].circuit
		}
	}

	return junctionBoxes, directlyConnected
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
	for range 1000 {
		junctionBoxes, directlyConnected = connectOnePair(junctionBoxes, directlyConnected)
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
