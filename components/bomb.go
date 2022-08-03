package main

import (
	"os"
	"strings"
)

type Bomb struct {
	Emoji string
	X     int
	Y     int
}

func NewBomb(x int, y int) *Bomb {
	return &Bomb{
		Emoji: bombSymbol,
		X:     x,
		Y:     y,
	}
}

func hasUnicodeSupport() bool {
	return strings.Contains(os.Getenv("LANG"), "UTF-8")
}
