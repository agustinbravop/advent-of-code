package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var INPUT_FILE = "01-input.txt"

func main() {
	text, err := os.ReadFile(INPUT_FILE)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	counter := 0
	dial := 50
	lines := strings.Split(string(text), "\n")
	for _, line := range lines {
		direction := line[:1]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Printf("Error parsing distance: %v\n", err)
			continue
		}

		if direction == "L" {
			dial = ((dial-distance)%100 + 100) % 100
		} else if direction == "R" {
			dial = ((dial+distance)%100 + 100) % 100
		}

		if dial == 0 {
			counter += 1
		}
	}

	fmt.Println(counter)
}
