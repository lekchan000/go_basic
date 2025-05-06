package main

import (
	"fmt"
	"sync"
)

func main() {

	/* 	fmt.Println("Hello, Mainthread 1")

	   	go hello()

	   	time.Sleep(1 * time.Microsecond)
	   	fmt.Println("Hello, Mainthread 2")
	*/

	var wg sync.WaitGroup
	wg.Add(3)

	for i := range 3 {
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Printf("Hello, Goroutine %d\n", i)
		}(i, &wg)
	}
	wg.Wait()
}

/* func hello() {
	for {
		fmt.Println("Hello, Goroutine")
	}
} */
