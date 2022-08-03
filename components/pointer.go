package main

const (
	RIGHT Direction = 1 + iota
	LEFT
	UP
	DOWN
)

type Direction int

type Pointer struct {
	Direction Direction
	X int
	Y int
}

func newPointer(d Direction, b []Cordiante) *Pointer {
	return &Pointer{
		Direction: d,
		Cordinate: b,
	}
}

func (p *Pointer) Move() error {
	cord := Cordiante{X: p.X, Y: p.Y}

	switch p.Direction {
	case RIGHT:
		cord.X++
	case LEFT:
		if cord.X > 0 {
			cord.X--
		}
	case UP:
		cord.Y++
	case DOWN:
		if cord.Y > 0 {
			cord.Y--
		}
	}
}

func (p *Pointer) changeDirection(x, y int) {
	moves := map[Direction][]
}
