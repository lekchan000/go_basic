package main

import "fmt"

type Airplane struct {
	Name string
}

func (a Airplane) Fly() {
	fmt.Printf("%s is flying\n", a.Name)
}

func main() {
	spitfire := Airplane{Name: "spitfire"}
	boeing := Airplane{Name: "boeing"}

	spitfire.Fly()
	boeing.Fly()
}
