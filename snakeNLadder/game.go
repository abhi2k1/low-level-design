package main

import "fmt"

type Game struct {
	board   *Board
	dice    *Dice
	players []*Player
}

func NewGame(board *Board, dice *Dice, players []*Player) *Game {
	return &Game{
		board:   board,
		dice:    dice,
		players: players,
	}
}

func (g *Game) StartGame() {
	winnerPlayerId := -1

	for winnerPlayerId == -1 {
		playingPlayer := g.players[0]
		fmt.Println("Player Turn: ", playingPlayer.id)
		g.players = g.players[1:]

		// roll the dice
		nextMove := g.dice.RollDice()
		nextPos := playingPlayer.GetCurrentPos() + nextMove

		cellObj := g.board.GetCell(nextMove)

		if cellObj.HasSnakeNLadder() {
			_, nextPos = cellObj.GetJumpPos()
			fmt.Println(fmt.Sprintf("Player got a jump, start: %v ,end: %v", playingPlayer.currentPos, nextPos))
		}

		if g.board.IsPlayerWinner(nextPos) {
			winnerPlayerId = playingPlayer.id
			break
		}

		playingPlayer.currentPos = nextPos

		g.players = append(g.players, playingPlayer)
		fmt.Println("Player Next move: ", nextPos)
	}

	fmt.Println("Winner Player Id: ", winnerPlayerId)
}
