package main

import "fmt"

type (
	Player struct{}
	Item   struct {
		Name    string
		Ability []Ability
	}
	Heal       struct{}
	DamageBuff struct{}

	Ability interface {
		Execute()
	}
)

func (p *Player) Use(item Item) {
	fmt.Println("Using item:", item.Name)

	for _, item := range item.Ability {
		item.Execute()
	}
}

func (h Heal) Execute() {
	fmt.Println("Healing")
}

func (d DamageBuff) Execute() {
	fmt.Println("Buffing damage 100%")
}

func main() {
	p := &Player{}
	steak := Item{
		Name:    "Steak",
		Ability: []Ability{Heal{}, DamageBuff{}},
	}
	p.Use(steak)
}
