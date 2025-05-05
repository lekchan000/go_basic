package main

import "fmt"

func main() {
	a := 10
	b := 20

	fmt.Println("a + b =", a+b) // Addition
	fmt.Println("a - b =", a-b) // Subtraction
	fmt.Println("a * b =", a*b) // Multiplication
	fmt.Println("a / b =", a/b) // Division
	fmt.Println("a % b =", a%b) // Modulus

	fmt.Println()

	fmt.Println("a == b =", a == b) // Equal to
	fmt.Println("a != b =", a != b) // Not equal to
	fmt.Println("a > b =", a > b)   // Greater than
	fmt.Println("a < b =", a < b)   // Less than
	fmt.Println("a >= b =", a >= b) // Greater than or equal to
	fmt.Println("a <= b =", a <= b) // Less than or equal to

	fmt.Println()

	fmt.Println("And", true && true) // Logical AND
	fmt.Println("Or", true || false) // Logical OR
	fmt.Println("Not", !true)        // Logical NOT

}
