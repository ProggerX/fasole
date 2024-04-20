package main

import tea "github.com/charmbracelet/bubbletea"
import "os"
import "fmt"

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			saveFile(m.tasks, os.Args[1])
			return m, tea.Quit
		} else if m.creationd.isShown == false {
			switch msg.String() {
			case "j":
			if m.cursor < len(m.tasks) - 1 {
				m.cursor++
			}
			case "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "a":
				m.creationd.isShown = true
			case "o":
				m.creationd.isShown = true
			case " ":
				if len(m.tasks) > 0 {
					m.tasks[m.cursor].is_complete = !m.tasks[m.cursor].is_complete
				}
			case "d":
				if len(m.tasks) > 0 {
					m.tasks = append(m.tasks[:m.cursor], m.tasks[m.cursor + 1:]...)
					if m.cursor > 0 {
						m.cursor--
					}
				}
			case "h":
				m.showingHelp = !m.showingHelp
			}
		} else {
			if msg.String() == "enter" {
				m.tasks = append(m.tasks, task{m.creationd.name, false})
				m.creationd.name = ""
				m.creationd.isShown = false
			} else if msg.String() == "backspace" {
				if len(m.creationd.name) > 0 {
					m.creationd.name = m.creationd.name[:len(m.creationd.name) - 1]
				}
			} else if msg.String() == "esc" {
				m.creationd.name = ""
				m.creationd.isShown = false
			} else {
				m.creationd.name += msg.String()
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	var s string
	if m.creationd.isShown == false {
		for i, itm := range m.tasks {
			if m.cursor == i {
				s += fmt.Sprintf("> [%s] %s\n", func() string { if itm.is_complete { return "X" } else { return " " }}(), itm.name)
			} else {
				s += fmt.Sprintf("  [%s] %s\n", func() string { if itm.is_complete { return "X" } else { return " " }}(), itm.name)
			}
		}
		s += "\n"
		if !m.showingHelp {
			s += "h - toggle help"
		} else {
			s += "ctrl+c or q - save and quit | j, k - select tasks | a or o - add task | d - delete task | space - mark as done"
		}
	} else {
		s += fmt.Sprintf("Enter name of the task, press enter when you finish, or esc to cancel:\n%s\n", m.creationd.name)
	}
	return s
}

func (m model) Init() tea.Cmd {
	return nil
}
