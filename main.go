package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	cursor  int
	choices []string
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
)

func initialModel() model {
	return model{
		choices: []string{
			"Day 1 - Part 1",
			"Day 1 - Part 2",
			"Day 2 - Part 1",
			"Day 2 - Part 2",
			"Day 3 - Part 1",
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
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
			runFunction(m.choices[m.cursor])
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
	return s
}

func runFunction(choice string) {
	switch choice {
	case "Day 1 - Part 1":
		runFile("01-secret-entrance.go", "Day1Part1")
	case "Day 1 - Part 2":
		runFile("01-secret-entrance.go", "Day1Part2")
	case "Day 2 - Part 1":
		runFile("02-gift-shop.go", "Day2Part1")
	case "Day 2 - Part 2":
		runFile("02-gift-shop.go", "Day2Part2")
	case "Day 3 - Part 1":
		runFile("03-lobby.go", "Day3Part1")
	}
}

func runFile(filename, funcName string) {
	content, _ := os.ReadFile(filename)
	code := string(content) + fmt.Sprintf("\n\nfunc main() { %s() }", funcName)

	tempFile := "temp.go"
	os.WriteFile(tempFile, []byte(code), 0644)
	defer os.Remove(tempFile)

	cmd := exec.Command("go", "run", tempFile)
	output, _ := cmd.CombinedOutput()

	// Print output with spacing to avoid collision
	fmt.Printf("\nOutput: %s\n", strings.TrimSpace(string(output)))
}

func main() {
	if len(os.Args) > 1 {
		// Direct function execution mode
		funcName := os.Args[1]
		switch funcName {
		case "Day1Part1":
			runFile("01-secret-entrance.go", "Day1Part1")
		case "Day1Part2":
			runFile("01-secret-entrance.go", "Day1Part2")
		case "Day2Part1":
			runFile("02-gift-shop.go", "Day2Part1")
		case "Day2Part2":
			runFile("02-gift-shop.go", "Day2Part2")
		case "Day3Part1":
			runFile("03-lobby.go", "Day3Part1")
		default:
			fmt.Printf("Unknown function: %s\n", funcName)
		}
		return
	}

	// Interactive menu mode if no argument was provided
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
