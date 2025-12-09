package main

import (
	"fmt"
	"os"
)

func day_2_part_1() {
	text, err := os.ReadFile("02-input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	fmt.Println(string(text))
}
