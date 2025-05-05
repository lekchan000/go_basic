package main

import "fmt"

func main() {
	hello("Calculator")

	fmt.Println("Sum:", add(2, 3))
	fmt.Println("Difference:", subtract(10, 4))
	fmt.Println("Product:", multiply(2, 6))
	fmt.Println("Quotient:", divide(8, 2))
}

func hello(name string) {
	fmt.Println("Hello", name)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) float64 {
	if b == 0 {
		fmt.Println("Cannot divide by zero")
		return 0
	}
	return float64(a) / float64(b)
}
