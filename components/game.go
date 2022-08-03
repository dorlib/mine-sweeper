package main

import (
	"github.com/nsf/termbox-go"
)

var (
	ScoreChan         = make(chan int)
	keyboardEventChan = make(chan keyboardEvent)
)

type Game struct {
	Board  *Board
	Score  int
	IsOver bool
}

func InitPointer() *Pointer {
	return newPointer(RIGHT, Cordiante{X: 0, Y: 0})
}

func InitScore() int {
	return 0
}

func InitBoard() *Board {
	return newBoard(InitPointer(), 10, 10, ScoreChan)
}

func (g *Game) Retry() {
	g.Board = InitBoard()
	g.Score = InitScore()
	g.IsOver = false
}

func (g *Game) AddScore(p int) {
	g.Score += p
}

func NewGame() *Game {
	board := InitBoard()
	pointer := InitPointer()
	printBoard(board, 10, 10, true, 10, 10, pointer)
	return &Game{Board: board, Score: InitScore()}

}

func (g *Game) End() {
	g.IsOver = true
}

// Start starts the game
func (g *Game) Start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	go listenToKeyBoard(keyboardEventChan)

	//if err := g.Render(); err != nil {
	//	panic(err)
	//}

mainloop:
	for {
		select {
		case p := <-ScoreChan:
			g.AddScore(p)
		case e := <-keyboardEventChan:
			switch e.eventType {
			case MOVE:
				d := keyToDirection(e.key)
				g.Board.pointer.changeDirection(d)
			case RETRY:
				g.Retry()
			case END:
				break mainloop
			}
		default:
			if !g.IsOver {
				if err := g.Board.MovePointer; err != nil {
					g.End()
				}
			}

			//if err := g.Render(); err != nil {
			//	panic(err)
			//}
		}
	}
}
