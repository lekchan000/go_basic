package main

import "fmt"

func main() {
	var a [3]int = [3]int{1, 2, 3}

	fmt.Println(a)

	a[0] = 4

	fmt.Println(a)

	for _, v := range a {
		fmt.Println(v)
	}
}
