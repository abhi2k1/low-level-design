package main

type peice int

const (
	Empty peice = iota
	Cross
	Circle
)

type PlayerPiece interface{
	GetPiece() string
}

type cross struct{}

func NewCross() PlayerPiece{
	return &cross{}
}

func (c *cross) GetPiece() string{
	return "X"
}

type circle struct{}

func NewCircle() PlayerPiece{
	return &circle{}
}

func (c *circle) GetPiece() string{
	return "0"
}

