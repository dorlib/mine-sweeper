package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

var (
	ScoreChan         = make(chan int)
	keyboardEventChan = make(chan keyboardEvent)
)

type Game struct {
	Board  [][]string
	Score  int
	IsOver bool
}

func InitPointer() *Pointer {
	return newPointer(RIGHT, Cordiante{X: 0, Y: 0})
}

func InitScore() int {
	return 0
}

func (g *Game) Retry() {
	g.Board = initBoard(10, 10)
	g.Score = InitScore()
	g.IsOver = false
}

func (g *Game) AddScore(p int) {
	g.Score += p
}

func NewGame() *Game {
	return &Game{Board: initBoard(10, 10), Score: InitScore()}
}

// Start starts the game
func (g *Game) Start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	go listenToKeyBoard(keyboardEventChan)

	if err := g.render(); err != nil {
		panic(err)
	}

mainLoop:
	for {
		select {
		case p := <-ScoreChan:
			g.AddScore(p)
		case e := <-keyboardEventChan:
			switch e.eventType {
			case MOVE:
				d := keyToDirection(e.key)
				g.Board.Pointer.changeDirection(d)
			case RETRY:
				g.retry()
			case END:
				break mainLoop
			}
		default:
			if !g.isOver {
				if err := g.arena.moveSnake(); err != nil {
					g.end()
				}
			}

			if err := g.render(); err != nil {
				panic(err)
			}

			time.Sleep(g.moveInterval())
		}
	}
}
