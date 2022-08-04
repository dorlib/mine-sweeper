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

func (b Board) Init() tea.Cmd {
	return nil
}

// getCloseCellsToReveal returns a slice of close cells to the current cell which we need to reveal.
func (b *Board) getCloseCellsToReveal(x int, y int) []*Cell {
	var cellsToReveal []*Cell

	for i := -1; i < 2; i++ {
		for j := -1; j < 1; j++ {
			if j != 0 || x != 0 {
				X := x + i
				Y := y + j
				if Y < b.height && Y > 0 && X < b.width && x > 0 {
					cellsToReveal = append(cellsToReveal, &b.Cells[y+i][x+j])
				}
			}
		}
	}
	return cellsToReveal
}

// showCell reveal the cell
func (b *Board) showCell() {
	// check if cell is already visible
	if b.Current.IsVisible || b.Current.IsHasFlag {
		return
	}
	// if we got here we can show the cell
	b.Current.IsVisible = true
	// check if it's a bomb
	if b.Current.IsBomb {
		b.GameState = gameOver
	}
}

// toggleFlag change the state of a flag in a cell
func (b *Board) toggleFlag() {
	hasFlag := b.Current.IsHasFlag

	if hasFlag {
		b.Current.IsHasFlag = !hasFlag
	}
}
