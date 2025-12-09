package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var INPUT_FILE = "01-input.txt"

func part_01() {
	text, err := os.ReadFile(INPUT_FILE)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	counter := 0
	dial := 50
	lines := strings.SplitSeq(string(text), "\n")
	for line := range lines {
		direction := line[:1]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Printf("Error parsing distance: %v\n", err)
			continue
		}

		switch direction {
		case "L":
			dial = ((dial-distance)%100 + 100) % 100
		case "R":
			dial = ((dial+distance)%100 + 100) % 100
		}

		if dial == 0 {
			counter += 1
		}
	}

	fmt.Println(counter)
}

func part_02() {
	text, err := os.ReadFile(INPUT_FILE)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	counter := 0
	dial := 50
	lines := strings.SplitSeq(string(text), "\n")
	for line := range lines {
		direction := line[:1]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Printf("Error parsing distance: %v\n", err)
			continue
		}

		switch direction {
		case "L":
			if dial-distance <= 0 {
				extra_click := 1
				if dial == 0 {
					extra_click = 0
				}
				counter += extra_click - (dial-distance)/100
			}
			dial = ((dial-distance)%100 + 100) % 100
		case "R":
			counter += (dial + distance) / 100
			dial = ((dial+distance)%100 + 100) % 100
		}
	}

	fmt.Println(counter)
}

func main() {
	part_02()
}
