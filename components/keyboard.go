package main

import (
	"github.com/nsf/termbox-go"
	_ "github.com/nsf/termbox-go"
)

type keyboardEventType int

//keyboard events
const (
	MOVE keyboardEventType = 1 + iota
	PLACE
	RETRY
	END
)

type keyboardEvent struct {
	eventType keyboardEventType
	key       termbox.Key
}

func keyToDirection(k termbox.Key) direction {
	switch k {
	case termbox.KeyArrowLeft:
		return LEFT
	case termbox.KeyArrowRight:
		return RIGHT
	case termbox.KeyArrowUp:
		return UP
	case termbox.KeyArrowDown:
		return DOWN
	default:
		return 0
	}
}

func listenToKeyBoard(evChan chan keyboardEvent) {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				evChan <- keyboardEvent{eventType: MOVE, key: ev.Key}
			case termbox.KeyArrowRight:
				evChan <- keyboardEvent{eventType: MOVE, key: ev.Key}
			case termbox.KeyArrowUp:
				evChan <- keyboardEvent{eventType: MOVE, key: ev.Key}
			case termbox.KeyArrowDown:
				evChan <- keyboardEvent{eventType: MOVE, key: ev.Key}
			case termbox.KeyEsc:
				evChan <- keyboardEvent{eventType: END, key: ev.Key}
			case termbox.KeySpace:
				evChan <- keyboardEvent{eventType: PLACE, key: ev.Key}
			case termbox.KeyEnter:
				evChan <- keyboardEvent{eventType: PLACE, key: ev.Key}
			default:
				if string(ev.Ch) == "r" {
					evChan <- keyboardEvent{eventType: RETRY, key: ev.Key}
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
