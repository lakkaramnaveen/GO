package main

import (
	"log"
	"os"
)

func main() {
	data := []byte("hello\ngo\n")
	// Write data to a file with 0644 permissions
	err := os.WriteFile("dat1", data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// For more granular writes
	f, err := os.Create("dat2")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.WriteString("writes\n"); err != nil {
		log.Fatal(err)
	}

	f.Sync() // Flush writes to stable storage
}
