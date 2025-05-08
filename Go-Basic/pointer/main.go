package main

import (
	"fmt"
)

func main() {
	b := new(int)
	c := new(*int)

	a := 1
	b = &a
	c = &b

	fmt.Printf("address a %v\n", &a)
	fmt.Printf("address b %v\n", &b)
	fmt.Printf("address c %v\n", &c)

	fmt.Println()

	fmt.Printf("value a %v\n", a)
	fmt.Printf("value b %v\n", b)
	fmt.Printf("value c %v\n", c)

	fmt.Println()

	fmt.Printf("*b %v\n", *b)
	fmt.Printf("*c %v\n", *c)
	fmt.Printf("**c %v\n", **c)

	fmt.Println()

	w := 2
	fmt.Println("w = ", w)

	Add(&w)

	fmt.Println("w = ", w)
}

func Add(w *int) {
	*w += 2
}
