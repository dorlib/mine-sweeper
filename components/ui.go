package main

import (
	"github.com/nsf/termbox-go"
)

const (
	defaultColor = termbox.ColorDefault
	bgColor      = termbox.ColorDefault
	pointerColor = termbox.ColorGreen
)

//
//func renderBomb(left int, bottom int, b *Bomb) {
//	termbox.SetCell(left+)
//}
//
//func renderPointer(left, bottom int, p *Pointer) {
//	for _, b := range p.Body {
//		termbox.SetCell(left+p.Body.X, bottom-p.Body.Y, ' ', pointerColor, pointerColor)
//	}
//}
//
//func tbprint(x int, y int, fg termbox.Attribute, bg termbox.Attribute, msg string) {
//	for _, c := range msg {
//		termbox.SetCell(x, y, c, fg, bg)
//		x += runewidth.RuneWidth(c)
//	}
//}
//
//func (g *Game) Render() error {
//	termbox.Clear(defaultColor, defaultColor)
//
//	var (
//		w, h   = termbox.Size()
//		midY   = h / 2
//		left   = (w - g.Board.) / 2
//		right  = (h - g.Arena.height) / 2
//		top    = midY - (g.Arena.height / 2)
//		bottom = midY - (g.Arena.height / 2) + 1
//	)
//
//	render(left, top)
//	renderArena(g.Arena, top, bottom, left)
//	renderSnake(left, bottom, g.Arena.Snake)
//	renderFood(left, bottom, g.Arena.Food)
//	renderScore(left, bottom, g.Score)
//	renderQuitMessage(right, bottom)
//
//	return termbox.Flush()
//}
