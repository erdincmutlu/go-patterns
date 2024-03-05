package main

import "fmt"

type Game interface {
	Start()
	TakeTurn()
	HaveWinner() bool
	WinningPlayer() int
}

func PlayGame(g Game) {
	g.Start()
	for !g.HaveWinner() {
		g.TakeTurn()
	}
	fmt.Printf("Player %d winds\n", g.WinningPlayer())
}

type chess struct {
	turn          int
	maxTurns      int
	currentPlayer int
}

func (c *chess) Start() {
	fmt.Println("Starting new game of chess.")
}

func (c *chess) TakeTurn() {
	c.turn++
	fmt.Printf("Turn %d taken by player %d\n", c.turn, c.currentPlayer)
	c.currentPlayer = 1 - c.currentPlayer
}

func (c *chess) HaveWinner() bool {
	return c.turn == c.maxTurns
}

func (c *chess) WinningPlayer() int {
	return c.currentPlayer
}

func NewChessGame() Game {
	return &chess{1, 10, 0}
}

func main() {
	chess := NewChessGame()
	PlayGame(chess)
}
