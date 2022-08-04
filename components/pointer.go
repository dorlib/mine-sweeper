package main

type Pointer struct {
	X int
	Y int
}

func (p *Pointer) moveUp(b *Board) {
	if p.Y > 0 {
		p.Y--
		b.Current = &b.Cells[b.Pointer.Y][b.Pointer.X]
	}
}

func (p *Pointer) moveDown(b *Board) {
	if p.Y < len(b.Cells) {
		p.Y++
		b.Current = &b.Cells[b.Pointer.Y][b.Pointer.X]
	}
}

func (p *Pointer) moveLeft(b *Board) {
	if p.X > 0 {
		p.X--
		b.Current = &b.Cells[b.Pointer.Y][b.Pointer.X]
	}
}

func (p *Pointer) moveRight(b *Board) {
	if p.X < len(b.Cells[0]) {
		p.X++
		b.Current = &b.Cells[b.Pointer.Y][b.Pointer.X]
	}
}
