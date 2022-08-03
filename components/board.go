package main

import (
	"fmt"
)

var hideSymbol string = "\U0001F972"
var pointerSymbol string = "➞"
var bombSymbol string = "💣"
var flagSymbol string = "🏳️"
var ok string = "🙂"
var lost string = "☹"

type Board struct {
	pointer   *Pointer
	board     [][]string
	scoreChan chan int
}

func initBoard(x, y int) *Board {
	board := make([][]string, x)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			board[i] = append(board[i], hideSymbol)
		}
	}
	return newBoard(board)
}

func printBoard(board *Board, detected int, left int, status bool, x int, y int, p *Pointer) {

	var s string
	if status {
		s = ok
	} else {
		s = lost
	}

	fmt.Printf("####################################\n")
	fmt.Printf("boms detected: %v     status: %v      bombs left: %v\n", detected, s, left)
	for i := 0; i < y; i++ {
		fmt.Printf("|          ")
		for j := 0; j < x; j++ {
			if j == p.Body.X && i == p.Body.Y {
				fmt.Printf(pointerSymbol)
			} else {
				fmt.Printf(" ")
			}
			fmt.Printf(board.board[i][j])
			fmt.Printf(" ")
		}
		fmt.Printf("          |\n")
	}
	fmt.Printf("####################################\n")
	fmt.Printf("Controls :\n")
	fmt.Printf("Use keybora arrows to navigate\n")
	fmt.Printf("Press space to place a flag on a spot\n")
	fmt.Printf("Press enter to reveal a spot\n")
	fmt.Printf("Press ESC to quit\n")
}

//func checkIfWin(board [][]string) bool {
//	switch {
//	case board[0][0] == board[0][1] && board[0][1] == board[0][2] && board[0][2] != " ":
//		return true
//	case board[1][0] == board[1][1] && board[1][1] == board[1][2] && board[1][2] != " ":
//		return true
//	case board[2][0] == board[2][1] && board[2][1] == board[2][2] && board[2][2] != " ":
//		return true
//	case board[0][0] == board[1][0] && board[1][0] == board[2][0] && board[2][0] != " ":
//		return true
//	case board[0][1] == board[1][1] && board[1][1] == board[2][1] && board[2][1] != " ":
//		return true
//	case board[0][2] == board[1][2] && board[1][2] == board[2][2] && board[2][2] != " ":
//		return true
//	case board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[2][2] != " ":
//		return true
//	case board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[2][0] != " ":
//		return true
//	default:
//		return false
//	}
//}

func (b *Board) MovePointer() error {
	if err := b.pointer.Move(); err != nil {
		return err
	}
	return nil
}
