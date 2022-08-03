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

func NewGame() *Game {
	return &Game{Board: initBoard(10, 10), Score: initialScore()}
}

// Start starts the game
func (g *Game) Start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	go listenToKeyBoard(keyboardEventsChan)

	if err := g.render(); err != nil {
		panic(err)
	}

mainLoop:
	for {
		select {
		case p := <-pointsChan:
			g.addPoints(p)
		case e := <-keyboardEventsChan:
			switch e.eventType {
			case MOVE:
				d := keyToDirection(e.key)
				g.arena.snake.changeDirection(d)
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
