package main

import (
    tea "github.com/charmbracelet/bubbletea"
	"os"
	"strings"
	"fmt"
	"github.com/charmbracelet/glamour"
	"net/http"
)

func parseFile(path string) ([]task, error) {
	var result []task

	bts, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var str string = string(bts)
	for _, t := range strings.Split(str, "\n")[:len(strings.Split(str, "\n")) - 1] {
		result = append(result, task{
			name: t[1:],
			is_complete: func() bool { if t[0] == 'X' {return true} else {return false}}(),
		})
	}

	return result, nil
}

func saveFile(tasks []task, path string) (error) {
	var str string = ""
	for _, t := range tasks {
		str += fmt.Sprintf("%s%s\n", func() string { if t.is_complete { return "X" } else { return " "} }(), t.name)
	}
	bts := []byte(str)
	err := os.WriteFile(path, bts, 0666)
	return err
}

func main() {
	var m model

	if len(os.Args[1:]) != 1 {
		fmt.Println("Expected exactly one argument - db filename. For more info use \"fasole help\"")
		os.Exit(1)
	} else if os.Args[1:][0] == "help" {
		 
	} else {
		tasks, err := parseFile(os.Args[1:][0])
		if err != nil {
			saveFile([]task{}, os.Args[1:][0])
			tasks, _ = parseFile(os.Args[1:][0])
		}
		m = initialModel(tasks)
		p := tea.NewProgram(m)
		p.Run()
	}
}
