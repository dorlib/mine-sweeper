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

// EmptyCellsToReveal returns a slice of cells that should be revealed.
func (b *Board) EmptyCellsToReveal() bool {
	var cellsToReveal []*Cell

	for i := 0; i < b.width; i++ {
		for j := 0; j < b.height; j++ {
			cell := b.Cells[i][j]

			if !cell.IsVisible || cell.closeBombs > 0 {
				continue
			}

			for _, c := range b.getCloseCells(i, j) {
				if !c.IsVisible {
					cellsToReveal = append(cellsToReveal, c)
				}
			}

		}
	}
	return len(cellsToReveal) > 0
}

// getCloseCells returns a neighbours of current cell.
func (b *Board) getCloseCells(x int, y int) []*Cell {
	var cells []*Cell

	for i := -1; i < 2; i++ {
		for j := -1; j < 1; j++ {
			if j != 0 || x != 0 {
				X := x + i
				Y := y + j
				if Y < b.height && Y > 0 && X < b.width && x > 0 {
					cells = append(cells, &b.Cells[y+i][x+j])
				}
			}
		}
	}
	return cells
}

// showCell reveal the cell.
func (b *Board) showCell() {
	// check if cell is already visible
	if b.Current.IsVisible || b.Current.IsHasFlag {
		return
	}
	// if we got here we can show the cell.
	b.Current.IsVisible = true
	// check if it's a bomb
	if b.Current.IsBomb {
		b.GameState = gameOver
	}
}

// toggleFlag change the state of a flag in a cell.
func (b *Board) toggleFlag() {
	hasFlag := b.Current.IsHasFlag

	if hasFlag {
		b.Current.IsHasFlag = !hasFlag
	}
}
