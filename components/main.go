package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"math/rand"
	"os"
)

type configuration struct {
	Bold int
}

var Config = configuration{
	Bold: 1,
}

// check returns 1 if cell is bomb
func check(field [][]Cell, x int, y int) int {
	if y < 0 || x < 0 || y >= len(field) || x >= len(field[0]) {
		return 0
	}
	if field[x][y].IsBomb {
		return 1
	}
	return 0
}

func generateField(height int, width int) [][]Cell {
	chanceToBomb := 10
	var field = make([][]Cell, height)
	for h := range field {
		field[h] = make([]Cell, width)
		for w := range field[h] {
			c := Cell{IsBomb: rand.Intn(100) < chanceToBomb}
			field[h][w] = c
		}
	}
	for h := range field {
		for w := range field[h] {
			for i := -1; i < 1; i++ {
				for j := -1; j < 1; j++ {
					if j != 0 || i != 0 {
						field[h][w].closeBombs += check(field, h+i, w+j)
					}
				}
			}
		}
	}
	return field
}

func initBoard() Board {
	field := generateField(20, 40)
	return Board{
		height:    10,
		width:     30,
		Cells:     field,
		Current:   &field[0][0],
		Pointer:   Pointer{X: 0, Y: 0},
		GameState: Normal,
	}
}
func main() {
	p := tea.NewProgram(initBoard())
	err := p.Start()
	if err != nil {
		fmt.Printf("error with %v", err)
		os.Exit(1)
	}
}
