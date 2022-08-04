package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

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
		case "enter", " ":
			return b.showCell()
		case "f":
			return b.toggleFlag()
		}
	}
	if b.revealEmptyNeighbours() {
		return b, tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
			return msg
		})
	}
	return b, nil
}
