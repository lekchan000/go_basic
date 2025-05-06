package main

import "fmt"

type (
	Number interface {
		int | float64
	}

	Player[T Number] struct {
		Name   string
		Age    int
		Damage T
	}

	Database[T Number] interface {
	}
)

func Sum[T Number](nums []T) T {
	var sum T
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	numsInt := []int{1, 2, 3}
	numsFloat := []float64{1.1, 2.2, 3.3}

	intSum := Sum(numsInt)
	floatSum := Sum(numsFloat)

	fmt.Printf("Sum of integers: %d\n", intSum)
	fmt.Printf("Sum of floats: %.1f\n", floatSum)
}
