package main

import (
	"fmt"
)

func mayPanic() {
	fmt.Println("About to panic...")
	panic("a problem happened") // Causes a panic with the message "a problem happened"
}

func main() {
	// Defer an anonymous function to recover from any panics in this goroutine
	defer func() {
		if r := recover(); r != nil {
			// 'r' is the value passed to panic()
			fmt.Println("Recovered from panic:", r)
		}
	}()

	mayPanic()
	fmt.Println("This line is reached because we recovered from the panic.")
}
