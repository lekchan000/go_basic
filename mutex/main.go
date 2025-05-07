package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu      sync.Mutex
	balance int = 1000
)

func main() {
	doneCh := make(chan bool, 3)

	go updateBalance(doneCh, 100)
	go updateBalance(doneCh, 150)
	go updateBalance(doneCh, 200)

	<-doneCh
	<-doneCh
	<-doneCh

	fmt.Println("Final balance:", balance)
}

func updateBalance(doneCh chan<- bool, amount int) {

	mu.Lock()

	fmt.Println("Updating balance...")
	time.Sleep(1 * time.Second)

	balance += amount
	doneCh <- true

	mu.Unlock()
	fmt.Println("Balance updated successfully.")
}
