package main

func main() {

	player1 := NewPlayer("Player 1", NewCross())
	player2 := NewPlayer("Player 2", NewCircle())

	game := NewGame()

	game.InitializeGame(3, []Player{player1, player2})

	game.StartGame()
}
