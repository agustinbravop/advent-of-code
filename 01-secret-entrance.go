package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day_1_part_1() {
	text, err := os.ReadFile("01-input.txt")
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

func day_1_part_2() {
	text, err := os.ReadFile("01-input.txt")
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
