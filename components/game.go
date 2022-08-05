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
			b.Pointer.moveUp(&b)
		case "down", "s":
			b.Pointer.moveDown(&b)
		case "left", "a":
			b.Pointer.moveLeft(&b)
		case "right", "d":
			b.Pointer.moveRight(&b)
		case "enter", " ":
			b.showCell()
		case "f":
			b.toggleFlag()
		}
	}
	if b.EmptyCellsToReveal() {
		return b, tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
			return msg
		})
	}
	return b, nil
}
