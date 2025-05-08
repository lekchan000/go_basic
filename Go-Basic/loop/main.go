package main

import "fmt"

func main() {
	/* 	for i := 1; i <= 3; i++ {
		println(i)
	} */

	var sum int
	for i := 1; i <= 10; i++ {
		sum += i
		fmt.Println(sum)
	}

	for i := range 3 {
		println(i)
	}
}
