package main

import (
	"fmt"
	"os"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func main() {
	// os.Args[1] captures the "nani" argument you typed in the terminal
	if len(os.Args) > 1 {
		name := os.Args[1]
		fmt.Println(Hello(name))
	} else {
		fmt.Println("Please provide a name!")
	}
}