package main

import "fmt"

func someRandomFunction() {
	defer func(msg string) {
		if r := recover(); r != nil {
			fmt.Println()
			fmt.Println("Successfully Recovered from Panic")
		}
	}("Panic Recovery Mechanism")

	fmt.Println("PANIC OCCURRED !!")
	panic("Something went wrong")
}

func main() {
	fmt.Println()
	fmt.Println("Program Starts...")
	fmt.Println()
	someRandomFunction()
	fmt.Println()
	fmt.Println("Program Continues even after panic...")
}
