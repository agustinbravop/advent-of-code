package main

import (
	"fmt"
	"maps"
	"os"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	cursor  int
	choices []string
	output  string
}

var (
	headerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true).
			Padding(0, 1)

	cursorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("212"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			Faint(true)

	dayMap = map[string]map[string]func() string{
		"1":  {"1": Day1Part1, "2": Day1Part2},
		"2":  {"1": Day2Part1, "2": Day2Part2},
		"3":  {"1": Day3Part1, "2": Day3Part2},
		"4":  {"1": Day4Part1, "2": Day4Part2},
		"5":  {"1": Day5Part1, "2": Day5Part2},
		"6":  {"1": Day6Part1, "2": Day6Part2},
		"7":  {"1": Day7Part1, "2": Day7Part2},
		"8":  {"1": Day8Part1, "2": Day8Part2},
		"9":  {"1": Day9Part1, "2": Day9Part2},
		"10": {"1": Day10Part1},
		"11": {"1": Day11Part1, "2": Day11Part2},
	}
)

func initialModel() model {
	choices := make([]string, 0)
	for _, day := range slices.Sorted(maps.Keys(dayMap)) {
		for part := range dayMap[day] {
			choices = append(choices, fmt.Sprintf("Day %s - Part %s", day, part))
		}
	}
	slices.Sort(choices)
	return model{choices: choices, output: ""}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q":
			if m.output != "" {
				fmt.Println()
			}
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			output := runFunction(m.choices[m.cursor])
			fmt.Println("\n\nOutput: " + output)
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := headerStyle.Render("Advent of Code 2025") + "\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		line := fmt.Sprintf("%s %s", cursor, choice)
		if m.cursor == i {
			line = cursorStyle.Render(line)
		}
		s += line + "\n"
	}

	s += "\n" + helpStyle.Render("↑/↓: Navigate  Enter: Run  q: Quit")

	if m.output != "" {
		s += "\n\nOutput: " + m.output
	}

	return s
}

func runFunction(choice string) string {
	var day, part int
	fmt.Sscanf(choice, "Day %d - Part %d", &day, &part)

	dayStr := fmt.Sprintf("%d", day)
	partStr := fmt.Sprintf("%d", part)

	if function, ok := dayMap[dayStr][partStr]; ok {
		return function()
	}
	return "Error: Function not found"
}

func main() {
	if len(os.Args) > 1 {
		if len(os.Args) < 3 {
			fmt.Printf("Missing arguments. Use: go run . <day> <part>\n")
			os.Exit(1)
		}

		day := os.Args[1]
		part := os.Args[2]

		if function, ok := dayMap[day][part]; ok {
			result := function()
			fmt.Println("Output:", result)
		} else {
			fmt.Printf("Invalid day or part. Use: go run . <day> <part>\n")
			os.Exit(1)
		}
	} else {
		p := tea.NewProgram(initialModel())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}
}
