package main

import "fmt"

func main() {
	n := 3

	if n%2 == 0 {
		fmt.Println("Is Even")
	} else {
		fmt.Println("Is Odd")

	}

	score := 65

	if score >= 80 {
		fmt.Println("Grade A")
	} else if score >= 70 {
		fmt.Println("Grade B")
	} else if score >= 60 {
		fmt.Println("Grade C")
	} else if score >= 50 {
		fmt.Println("Grade D")
	} else {
		fmt.Println("Grade F")
	}

	if a := 10 > 20; a {
		fmt.Println("10 is greater than 20")
	} else {
		fmt.Println("10 is less than 20")

		fmt.Println(a)
	}
}
