package main

import (
	"fmt"
)

type Game interface{
	InitializeGame(boardSize int, players []Player)
	StartGame()
}

type game struct{
	board Board
	players []Player
	currentPlayerIndex int
}


func NewGame() Game{
	return &game{}
}

func (g *game) InitializeGame(boardSize int, players []Player){
	// Initialize the board
	g.board = NewBoard(boardSize)
	g.players = players
	g.currentPlayerIndex = 0
}

func (g *game) StartGame(){
	for {
		g.board.PrintBoard()
		player := g.players[g.currentPlayerIndex]

		// take  a input from the player
		var x, y int
		fmt.Printf("Player %s, enter the coordinates x and y to place your piece: ", player.Name())
		fmt.Scanf("%d %d", &x, &y)
		
		err := g.board.PlacePiece(x, y, player.Piece())
		if err != nil {
			fmt.Println(err)
			continue
		}

		if winner, isBoardFull := g.board.CheckforWinner(); winner != nil || isBoardFull {
			if winner != nil {
				g.board.PrintBoard()

				winningPlayer := findWinningPlayer(winner, g.players)

				fmt.Println("Player", winningPlayer.Name(), "has won the game with piece", winningPlayer.Piece().GetPiece())
			} else {
				fmt.Println("The game is a draw")
			}
			break
		}
		g.currentPlayerIndex = (g.currentPlayerIndex + 1) % len(g.players)
	}

	return
}

func findWinningPlayer(piece PlayerPiece, players []Player) Player{
	for _, player := range players{
		if player.Piece() == piece{
			return player
		}
	}
	return nil
}

