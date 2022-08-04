package main

import tea "github.com/charmbracelet/bubbletea"

func (b Board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return b, tea.Quit
		case "up", "w":
			return b.Pointer.moveUp(&b)
		case "down", "s":
			return b.Pointer.moveDown(&b)
		case "left", "a":
			return b.Pointer.moveRight(&b)
		case "right", "d":
			return b.Pointer.moveLeft(&b)
		}
	}
	return
}
