package main

import "fmt"

type tokenType int

type Token struct {
	Content    rune
	Type       tokenType
	IsSelected bool
}

const (
	TableSpace tokenType = iota
	TableComponent
	Neighbours
	Flag
	Unknown
	Empty
	Bomb
)

const (
	flag    = '⚑'
	bomb    = '💣'
	unknown = '❔'
)

const (
	backgroundColor          = "\033[1;40m"
	foregroundColor          = "\033[1;37m"
	tableForegroundColor     = "\033[1;30m"
	selectedBackgroundColor  = "\033[1;44m"
	neighbourForegroundColor = "\033[1;%sm"
	flagForegroundColor      = "\033[1;35m"
	bombForegroundColor      = "\033[1;31m"
)

// CreateBoardCell create a token with emoji to any type to cell
func CreateBoardCell(c Cell) Token {
	if c.IsVisible {
		item := rune('0' + c.closeBombs)
		itemType := Neighbours
		if c.closeBombs == 0 {
			item = ' '
			itemType = Empty
		}
		if c.IsBomb {
			item = bomb
			itemType = Bomb
		}
		return Token{Content: item, Type: itemType}
	}
	if c.IsHasFlag {
		return Token{Content: flag, Type: Flag}
	}
	return Token{Content: unknown, Type: Unknown}
}

func (t Token) print() string {
	c := t.Content
	backgroundStyle := backgroundColor
	foregroundStyle := foregroundColor

	if t.Type == Neighbours {
		foregroundStyle = neighbourForegroundColor
	}
	if t.Type == Flag {
		foregroundStyle = flagForegroundColor
	}
	if t.Type == Bomb {
		foregroundStyle = bombForegroundColor
	}
	if t.IsSelected {
		backgroundStyle = selectedBackgroundColor
	}
	if t.Type == TableComponent {
		foregroundStyle = tableForegroundColor
	}

	return fmt.Sprintf("%s%s%c", backgroundStyle, foregroundStyle, c)
}

func colorNeighbour(r rune) string {
	switch r {
	case '1':
		return "32"
	case '2':
		return "33"
	case '3':
		return "31"
	case '4':
		return "35"
	case '5':
		return "36"
	default:
		return "36"
	}
}

func addStructRow(row []Token, numOfElements int, start rune, separator rune, end rune) {
	row[0] = Token{Content: start, Type: TableComponent}
	for i := 0; i < numOfElements; i++ {
		base := 1 + 4*i
		row[base] = Token{Content: '─', Type: TableComponent}
		row[base+1] = Token{Content: '─', Type: TableComponent}
		row[base+2] = Token{Content: '─', Type: TableComponent}
		row[base+3] = Token{Content: separator, Type: TableComponent}
	}
	row[len(row)-1] = Token{Content: end, Type: TableComponent}
}

func boardPositionToViewModelPosition(y int, x int) (int, int) {
	return 1 + y*2, 2 + x*4
}
