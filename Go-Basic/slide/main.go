package main

import "fmt"

func main() {
	b := []int{1, 2, 3}
	fmt.Println(b)

	c := make([]int, 0)

	c = append(c, 4, 5, 6)

	fmt.Println(c)

	c = append(c, 7, 8, 9)

	fmt.Println(c)

	d := []int{1, 2, 3}

	double(d)

	fmt.Println(d)
}

func double(nums []int) {
	for i := range nums {
		nums[i] *= 2
	}
}
