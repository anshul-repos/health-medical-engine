package main

import (
	"fmt"
	"sync"
)

func printEven(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		if num%2 == 0 {
			fmt.Printf("Even:%d\n", num)
		}
	}
}

func printOdd(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		if num%2 != 0 {
			fmt.Printf("Odd:%d\n", num)
		}
	}
}

func main() {
	fmt.Println()
	chEv := make(chan int)
	chOd := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go printEven(chEv, &wg)
	go printOdd(chOd, &wg)

	for i := 1; i < 10; i++ {
		chOd <- i
		chEv <- i
	}

	close(chEv)
	close(chOd)
	wg.Wait()

	fmt.Println()
}
