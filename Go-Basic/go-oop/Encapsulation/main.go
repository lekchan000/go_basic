package main

import "fmt"

type (
	PlayerClass interface {
		Attack()
	}

	Warrior struct {
		Weapon string
	}

	Mage struct {
		Spell string
	}
)

func PlayerAttack(playerClass PlayerClass) {
	playerClass.Attack()
}

func (w Warrior) Attack() {
	fmt.Println("Warrior attacks with", w.Weapon)
}

func (m Mage) Attack() {
	fmt.Println("Mage casts", m.Spell)
}

func main() {
	warrior := Warrior{
		Weapon: "Sword",
	}
	mage := Mage{
		Spell: "Fireball",
	}

	warrior.Attack()
	mage.Attack()

	PlayerAttack(warrior)
	PlayerAttack(mage)

}
