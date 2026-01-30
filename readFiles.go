package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Read the entire file content
	content, err := os.ReadFile("book_data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	// For more control, open the file
	f, err := os.Open("book_data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // Ensure the file is closed

	// Read some bytes
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
}
