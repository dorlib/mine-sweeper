package main

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

}
