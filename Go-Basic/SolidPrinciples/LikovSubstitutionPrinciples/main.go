package main

import "fmt"

type (
	Aircraft interface {
		Fly()
	}

	Boeing747 struct{ Aircraft }

	Spitfire struct{ Aircraft }
)

func (b *Boeing747) Fly() {
	fmt.Println("Boeing747 is flying")
}

func (s *Spitfire) Fly() {
	fmt.Println("Spitfire is flying")
}

func (s *Spitfire) Fire() {
	fmt.Println("Spitfire is firing")
}

func main() {
	spitfire := &Spitfire{}
	spitfire.Fly()
	spitfire.Fire()

	boeing747 := &Boeing747{}
	boeing747.Fly()
}
