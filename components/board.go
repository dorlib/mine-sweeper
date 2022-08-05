package main

import tea "github.com/charmbracelet/bubbletea"

type GameState int

const (
	Normal GameState = iota
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

func (b Board) View() string {
	height := len(b.Cells)*2 + 1
	width := len(b.Cells[0])*4 + 1
	viewModel := make([][]Token, height)

	for i := 0; i < len(viewModel); i++ {
		viewModel[i] = make([]Token, width)
	}

	numOfElements := len(b.Cells[0])
	addStructRow(viewModel[0], numOfElements, '┌', '┬', '┐')
	for h, boardRow := range b.Cells {
		if h != 0 {
			addStructRow(viewModel[h*2], numOfElements, '├', '┼', '┤')
		}
		viewRow := viewModel[h*2+1]
		for w, cell := range boardRow {
			base := w * 4
			item := CreateBoardCell(cell)
			viewRow[base] = Token{Content: '│', Type: TableComponent}
			viewRow[base+1] = Token{Content: ' ', Type: TableSpace}
			viewRow[base+2] = item
			viewRow[base+3] = Token{Content: ' ', Type: TableSpace}
		}
		viewRow[len(viewRow)-1] = Token{Content: '│', Type: TableComponent}
	}
	addStructRow(viewModel[height-1], numOfElements, '└', '┴', '┘')

	// Select the 'selected' cell
	vmY, vmX := boardPositionToViewModelPosition(b.Pointer.Y, b.Pointer.X)
	for offsetY := -1; offsetY <= 1; offsetY++ {
		for offsetX := -2; offsetX <= 2; offsetX++ {
			viewModel[vmY+offsetY][vmX+offsetX].IsSelected = true
		}
	}
	s := ""
	for _, row := range viewModel {
		for _, cell := range row {
			s += cell.print()
		}
		s += "\n"
	}
	return s
}

type Cell struct {
	IsBomb     bool
	IsVisible  bool
	IsHasFlag  bool
	closeBombs int
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
