package main

import "fmt"

type (
	Player  struct{}
	Display struct{}
)

func (p *Player) Move() {
	fmt.Println("Player is moving")
}

func (p *Player) Attack() {
	fmt.Println("Player is attacking")
}

func (d *Display) DisplayScoreBoard() {
	fmt.Println("Displaying scoreboard")
}

func main() {
	player := &Player{}
	display := &Display{}

	player.Move()
	player.Attack()

	display.DisplayScoreBoard()
}
