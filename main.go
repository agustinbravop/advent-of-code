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

	dayMap = map[string]map[string]string{
		"1": {"1": "Day1Part1", "2": "Day1Part2"},
		"2": {"1": "Day2Part1", "2": "Day2Part2"},
		"3": {"1": "Day3Part1", "2": "Day3Part2"},
		"4": {"1": "Day4Part1", "2": "Day4Part2"},
		"5": {"1": "Day5Part1", "2": "Day5Part2"},
	}

	fileMap = map[string]string{
		"Day1Part1": "01-secret-entrance.go",
		"Day1Part2": "01-secret-entrance.go",
		"Day2Part1": "02-gift-shop.go",
		"Day2Part2": "02-gift-shop.go",
		"Day3Part1": "03-lobby.go",
		"Day3Part2": "03-lobby.go",
		"Day4Part1": "04-printing-department.go",
		"Day4Part2": "04-printing-department.go",
		"Day5Part1": "05-cafeteria.go",
		"Day5Part2": "05-cafeteria.go",
	}
)

func initialModel() model {
	choices := make([]string, 0)
	for day := range dayMap {
		for part := range dayMap[day] {
			choices = append(choices, fmt.Sprintf("Day %s - Part %s", day, part))
		}
	}
	return model{choices: choices}
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
	var day, part int
	fmt.Sscanf(choice, "Day %d - Part %d", &day, &part)

	dayStr := fmt.Sprintf("%d", day)
	partStr := fmt.Sprintf("%d", part)

	if funcName, ok := dayMap[dayStr][partStr]; ok {
		runFile(fileMap[funcName], funcName)
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

	fmt.Printf("\nOutput: %s\n", strings.TrimSpace(string(output)))
}

func main() {
	if len(os.Args) > 1 {
		if len(os.Args) < 3 {
			fmt.Printf("Missing arguments. Use: go run main.go <day> <part>\n")
			os.Exit(1)
		}

		day := os.Args[1]
		part := os.Args[2]

		if funcName, ok := dayMap[day][part]; ok {
			runFile(fileMap[funcName], funcName)
		} else {
			fmt.Printf("Invalid day or part. Use: go run main.go <day> <part>\n")
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
