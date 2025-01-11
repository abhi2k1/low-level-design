package main

type Player struct {
	id         int
	currentPos int
}

func NewPlayer(id, currentPos int) *Player {
	return &Player{
		id:         id,
		currentPos: currentPos,
	}
}

func (p *Player) GetCurrentPos() int {
	return p.currentPos
}

func (p *Player) SetCurrentPos(pos int) {
	p.currentPos = pos
}
