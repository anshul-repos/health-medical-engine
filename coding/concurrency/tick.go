package main

import (
	"fmt"
	"sync"
	"time"
)

// Timer-based Worker
// Write a goroutine that prints "tick" every second until stopped.

func ticker(done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(time.Second * 1)

	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		case <-done:
			ticker.Stop()
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	wg.Add(1)

	go ticker(done, &wg)

	time.Sleep(time.Second * 10)

	done <- struct{}{}
	wg.Wait()
}
