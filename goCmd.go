package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// Command and arguments are specified separately
	cmd := exec.Command("ls", "-ltr") 

	// Run the command and get the output
	out, err := cmd.Output()
	if err != nil {
		// Log error if command failed
		log.Fatal(err)
	}

	// Print the output
	fmt.Println("Command Successfully Executed")
	fmt.Println(string(out))
}
