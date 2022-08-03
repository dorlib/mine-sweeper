package main

import (
	"os"
	"strings"
)

type Bomb struct {
	Emoji rune
	X     int
	Y     int
}

func NewBomb(x int, y int) *Food {
	return &Food{
		Emoji: bombSymbol,
		X:     x,
		Y:     y,
	}
}

func hasUnicodeSupport() bool {
	return strings.Contains(os.Getenv("LANG"), "UTF-8")
}
