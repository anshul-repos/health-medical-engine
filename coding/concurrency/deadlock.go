package main

// write a code to implement deadlock in go

func main() {
	// A deadlock occurs when two or more goroutines are waiting for each other to release resources, causing them to be stuck indefinitely.
	// Here is a simple example of a deadlock in Go:

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	go func() {
		ch1 <- struct{}{} // Send signal to ch1
		ch2 <- struct{}{} // Wait for signal from ch2 (this will block)
	}()

	go func() {
		ch2 <- struct{}{} // Send signal to ch2
		ch1 <- struct{}{} // Wait for signal from ch1 (this will block)
	}()

	// This will cause a deadlock because both goroutines are waiting for each other.
	select {}
}
