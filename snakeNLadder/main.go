package main

func main() {
	p1 := NewPlayer(1, 0)
	p2 := NewPlayer(2, 0)

	dice := NewDice(1, 6)

	board := NewBoard(10, 10, 5)

	game := NewGame(board, dice, []*Player{p1, p2})

	game.StartGame()
}
