package main

type Player interface{
	Name() string
	Piece() PlayerPiece
}

type player struct{
	name string
	piece PlayerPiece
}

func NewPlayer(name string, piece PlayerPiece) Player{
	return &player{
		name: name,
		piece: piece,
	}
}

func (p *player) Name() string{
	return p.name
}

func (p *player) Piece() PlayerPiece{
	return p.piece
}