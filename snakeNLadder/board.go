package main

import "math/rand"

type Board struct {
	cellObj    [][]*cell
	dimensions int // it can be ignored, as because we can find the dimension via using length of cellObj
}

func NewBoard(dimension int, snakeCnt, ladderCnt int) *Board {
	cellObj := make([][]*cell, 0)

	for i := 0; i < dimension; i++ {
		temp := make([]*cell, 0)

		for j := 0; j < dimension; j++ {
			newCell := NewCell(nil)

			temp = append(temp, newCell)
		}

		cellObj = append(cellObj, temp)
	}

	for snakeCnt > 0 {
		startPos := rand.Intn(dimension * dimension)
		endPos := rand.Intn(dimension * dimension)

		if startPos <= endPos {
			continue
		}

		startCellObj := cellObj[(startPos)/dimension][(startPos)%dimension]
		if startCellObj.HasSnakeNLadder() {
			continue
		}

		jumpObj := NewJump(startPos, endPos)

		startCellObj.jump = jumpObj
		snakeCnt--
	}

	for ladderCnt > 0 {
		startPos := rand.Intn(dimension * dimension)
		endPos := rand.Intn(dimension * dimension)

		if startPos >= endPos {
			continue
		}

		startCellObj := cellObj[(startPos)/dimension][(startPos)%dimension]
		if startCellObj.HasSnakeNLadder() {
			continue
		}

		jumpObj := NewJump(startPos, endPos)

		startCellObj.jump = jumpObj
		ladderCnt--
	}

	return &Board{
		cellObj:    cellObj,
		dimensions: dimension,
	}
}

func (b *Board) IsPlayerWinner(nextMove int) bool {
	if nextMove > (b.dimensions*b.dimensions)-1 {
		return true
	}

	return false
}

func (b *Board) GetCell(pos int) *cell {
	row := pos / b.dimensions

	col := pos % b.dimensions

	return b.cellObj[row][col]
}
