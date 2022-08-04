package main

import tea "github.com/charmbracelet/bubbletea"

type GameState int

const (
	Noraml GameState = iota
	gameOver
	gameWon
)

type Board struct {
	height    int
	width     int
	Cells     [][]Cell
	Pointer   Pointer
	Current   *Cell
	GameState GameState
}

func (m Board) Init() tea.Cmd {
	return nil
}
