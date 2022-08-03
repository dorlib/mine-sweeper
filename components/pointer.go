package main

import "errors"

const (
	RIGHT Direction = 1 + iota
	LEFT
	UP
	DOWN
)

type Direction int

type Pointer struct {
	Direction Direction
	Body      Cordiante
}

func newPointer(d Direction, b Cordiante) *Pointer {
	return &Pointer{
		Direction: d,
		Body:      b,
	}
}

func (p *Pointer) Move() error {
	cord := Cordiante{X: p.Body.X, Y: p.Body.Y}

	switch p.Direction {
	case RIGHT:
		if cord.X < 10 {
			cord.X++
		}
	case LEFT:
		if cord.X > 0 {
			cord.X--
		}
	case UP:
		if cord.Y < 10 {
			cord.Y++
		}
	case DOWN:
		if cord.Y > 0 {
			cord.Y--
		}
	}
	return nil
}

func (p *Pointer) changeDirection(d Direction) {
	opposites := map[Direction]Direction{
		RIGHT: LEFT,
		LEFT:  RIGHT,
		UP:    DOWN,
		DOWN:  UP,
	}

	opp := opposites[d]

	if opp != 0 && opp != p.Direction {
		p.Direction = d
	}
}

func (p *Pointer) Die() error {
	return errors.New("game over")

}
