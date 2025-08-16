package main

import (
	"fmt"
	"sync"
)

// write a go code to print "ping" "pong" one at a time N times

func ping(pingCh, pongCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= N; i++ {
		<-pingCh
		fmt.Println("PING")
		pongCh <- struct{}{}
	}

}

func pong(pingCh, pongCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= N; i++ {
		<-pongCh
		fmt.Println("PONG")
		if i < N { // avoid extra send at the end
			pingCh <- struct{}{}
		}
	}
}

const N int = 5

func main() {

	pingCh := make(chan struct{})
	pongCh := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(2)

	go ping(pingCh, pongCh, &wg)
	go pong(pingCh, pongCh, &wg)

	pingCh <- struct{}{}

	wg.Wait()
}
