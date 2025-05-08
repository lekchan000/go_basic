package main

type (
	Aircraft interface {
		Fly()
	}

	FighterAircraft interface {
		Aircraft
		Fire()
	}

	CommercialAircraft interface {
		Aircraft
		AirCondition()
	}

	Boeing747 struct{}
	Spitfire  struct{}
)

func (b Boeing747) Fly() {
	println("Boeing747 flying")
}

func (b Boeing747) AirCondition() {
	println("Boeing747 air conditioning")
}

func (s Spitfire) Fly() {
	println("Spitfire flying")
}

func (s Spitfire) Fire() {
	println("Spitfire firing")
}

func main() {

	spitfire := Spitfire{}
	spitfire.Fly()
	spitfire.Fire()

	boeing747 := Boeing747{}
	boeing747.Fly()
	boeing747.AirCondition()

}
