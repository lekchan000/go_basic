package main

import "fmt"

type (
	Keyboard interface {
		Input()
	}

	MechanicalKeyboard struct {
		SwitchType string
		Size       string
		OS         string
	}

	NomalKeyboard struct {
		IsCanPress bool
	}
)

func (m MechanicalKeyboard) Input() {
	fmt.Println("Press the key on the Mechanical Keyboard :", m.SwitchType, m.Size, m.OS)
}

func (n NomalKeyboard) Input() {
	fmt.Println("Press the key on the Nomal Keyboard")
}

func main() {
	mechanicalKeyboard := MechanicalKeyboard{
		SwitchType: "Blue",
		Size:       "Full",
		OS:         "Windows",
	}
	nomalKeyboard := NomalKeyboard{
		IsCanPress: true,
	}

	mechanicalKeyboard.Input()
	nomalKeyboard.Input()
}
