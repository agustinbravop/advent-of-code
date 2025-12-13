package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Toggles the state of the lights the button controls only.
func pressButton(lights []bool, button []int) []bool {
	newLights := make([]bool, len(lights))
	copy(newLights, lights)

	for _, pos := range button {
		newLights[pos] = !newLights[pos]
	}
	return newLights
}

// Returns the fewest button presses required to activate the target lights.
func fewestButtonPressesToSolve(targetLights []bool, buttons [][]int) int {
	// Initially, all lights are off.
	lights := make([]bool, len(targetLights))

	// Do a breadth-first search of solutions.
	queue := [][]bool{lights}
	visited := map[string]bool{}
	visited[fmt.Sprintf("%v", lights)] = true

	for depth := 0; len(queue) > 0; depth++ {
		nextQueue := [][]bool{}
		for _, state := range queue {
			if reflect.DeepEqual(state, targetLights) {
				return depth
			}

			// Consider each button as a step towards a solution.
			for _, button := range buttons {
				nextState := pressButton(state, button)
				if !visited[fmt.Sprintf("%v", nextState)] {
					nextQueue = append(nextQueue, nextState)
					visited[fmt.Sprintf("%v", nextState)] = true
				}
			}
		}
		queue = nextQueue
	}

	return -1
}

func Day10Part1() string {
	text, err := os.ReadFile("10-input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading input file: %v\n", err)
	}

	fewestButtonPresses := 0
	for line := range strings.SplitSeq(string(text), "\n") {
		// Parse indicator lights diagram.
		lights := []bool{}
		endOfLights := strings.Index(line, "]")
		for _, char := range line[:endOfLights] {
			switch char {
			case '.':
				lights = append(lights, false)
			case '#':
				lights = append(lights, true)
			}
		}

		// Parse button wiring schematics.
		buttons := [][]int{}
		startOfJoltages := strings.Index(line, "{")
		schematics := strings.SplitSeq(line[endOfLights+3:startOfJoltages-2], ") (")
		for schematic := range schematics {
			button := []int{}
			parts := strings.SplitSeq(schematic, ",")
			for part := range parts {
				num, _ := strconv.Atoi(part)
				button = append(button, num)
			}
			buttons = append(buttons, button)
		}

		fewestButtonPresses += fewestButtonPressesToSolve(lights, buttons)
	}

	return fmt.Sprintf("%d", fewestButtonPresses)
}
