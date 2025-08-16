package main

import (
	"fmt"
	"time"
)

// write a code to implement deadlock in go

func main() {
	// A deadlock occurs when two or more goroutines are waiting for each other to release resources, causing them to be stuck indefinitely.
	// Here is a simple example of a deadlock in Go:

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	// Goroutine 1
	go func() {
		fmt.Println("G1: trying to send to ch1")
		ch1 <- struct{}{} // waits until someone receives
		fmt.Println("G1: sent to ch1")

		fmt.Println("G1: trying to receive from ch2")
		<-ch2 // waits until someone sends
		fmt.Println("G1: received from ch2")
	}()

	// Goroutine 2
	go func() {
		fmt.Println("G2: trying to send to ch2")
		ch2 <- struct{}{} // waits until someone receives
		fmt.Println("G2: sent to ch2")

		fmt.Println("G2: trying to receive from ch1")
		<-ch1 // waits until someone sends
		fmt.Println("G2: received from ch1")
	}()

	// Give some time before runtime detects deadlock
	time.Sleep(2 * time.Second)

	fmt.Println("Main: program stuck → deadlock!")
	// This will cause a deadlock because both goroutines are waiting for each other.
	select {}
}
