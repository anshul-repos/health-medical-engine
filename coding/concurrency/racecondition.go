package main

import (
	"fmt"
	"sync"
)

// Withouth sync.mutex, there will be race condition, if you run code multiple times you'll get different output.
// fix with sync.mutex, always get output balance : 1000

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
