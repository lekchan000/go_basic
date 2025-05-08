package main

import "fmt"

func main() {
	jobch := make(chan int, 10)
	resultch := make(chan int, 10)

	for i := range 10 {
		jobch <- i + 1
	}
	close(jobch)

	go double(jobch, resultch)

	for range 10 {
		result := <-resultch
		fmt.Println(result)
	}
}

func double(jobch <-chan int, resultch chan<- int) {
	for j := range jobch {
		resultch <- j * 2
	}
}
