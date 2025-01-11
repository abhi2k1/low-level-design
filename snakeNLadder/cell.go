package main

type jump struct {
	startPos int
	endPos   int
}

func NewJump(startPos, endPos int) *jump {
	return &jump{
		startPos: startPos,
		endPos:   endPos,
	}
}

func (j *jump) GetPos() (int, int) {
	return j.startPos, j.endPos
}

type cell struct {
	jump *jump
}

func NewCell(jump *jump) *cell {
	return &cell{
		jump: jump,
	}
}

func (c *cell) SetJump(jum *jump) {
	c.jump = jum
}

func (c *cell) HasSnakeNLadder() bool {
	return c.jump != nil
}

func (c *cell) GetJumpPos() (int, int) {
	return c.jump.GetPos()
}
