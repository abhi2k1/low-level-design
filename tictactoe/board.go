package main

import (
	"errors"
	"fmt"
)

type Board interface {
	PlacePiece(x, y int, piece PlayerPiece) error
	CheckforWinner() (PlayerPiece, bool)
	PrintBoard()
}

type board struct {
	size int
	grid [][]PlayerPiece
}

func NewBoard(size int) *board {
	grid := make([][]PlayerPiece, size)
	for i := range grid {
		grid[i] = make([]PlayerPiece, size)
	}
	return &board{
		size: size,
		grid: grid,
	}
}

func (b *board) PlacePiece(x, y int, piece PlayerPiece) (err error) {
	if x < 0 || x >= b.size || y < 0 || y >= b.size {
		return errors.New("Invalid move")
	}

	if b.grid[x][y] != nil {
		return errors.New("Invalid move")
	}

	b.grid[x][y] = piece

	return
}

func (b *board) CheckforWinner() (winnerPiece PlayerPiece, isboardFull bool) {

	// check for rows
	for i := 0; i < b.size; i++ {
		for j := 1; j < b.size; j++ {
			if b.grid[i][j] == nil || b.grid[i][j] != b.grid[i][j-1] {
				break
			}

			if j == b.size-1 {
				return b.grid[i][j], false
			}
		}
	}

	// check for columns
	for i := 0; i < b.size; i++ {
		for j := 1; j < b.size; j++ {
			if b.grid[j][i] == nil || b.grid[j][i] != b.grid[j-1][i] {
				break
			}

			if j == b.size-1 {
				return b.grid[j][i], false
			}
		}
	}

	// check for diagonals
	for i := 1; i < b.size; i++ {
		if b.grid[i][i] == nil || b.grid[i][i] != b.grid[i-1][i-1] {
			break
		}

		if i == b.size-1 {
			return b.grid[i][i], false
		}
	}

	for i := 1; i < b.size; i++ {
		if b.grid[i][b.size-i-1] == nil || b.grid[i][b.size-i-1] != b.grid[i-1][b.size-i] {
			break
		}

		if i == b.size-1 {
			return b.grid[i][b.size-i-1], false
		}
	}

	// check if board is full
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if b.grid[i][j] == nil {
				return nil, false
			}
		}
	}

	return nil, true
}

func (b *board) PrintBoard() {
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if b.grid[i][j] == nil {
				fmt.Print("_")
			} else {
				fmt.Print(b.grid[i][j].GetPiece())
			}

			fmt.Print(" | ")
		}
		fmt.Println()
	}
	fmt.Println()
}
