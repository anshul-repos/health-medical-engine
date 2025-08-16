package main

import (
	"fmt"
	"sync"
)

var balance int = 1000 // shared resource
var mu sync.Mutex

func deposit(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i < 1000; i++ {
		mu.Lock()
		balance = balance + i
		mu.Unlock()
	}

}

func withdraw(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i < 1000; i++ {
		mu.Lock()
		balance = balance - i
		mu.Unlock()
	}

}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go deposit(&wg)
	go withdraw(&wg)

	wg.Wait()

	fmt.Printf("Balance: %+v", balance)
}
