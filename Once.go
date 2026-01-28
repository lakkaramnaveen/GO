package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	once sync.Once
	wg   sync.WaitGroup
	config string
)

func initializeConfig() {
	// This function will be called exactly once
	fmt.Println("Initializing configuration...")
	time.Sleep(time.Second) // Simulate a heavy initialization task
	config = "Loaded Config"
	fmt.Println("Configuration initialized.")
}

func worker(id int) {
	defer wg.Done()
	
	// All workers call this, but initializeConfig() runs only once
	once.Do(initializeConfig) 

	fmt.Printf("Worker %d is using config: %s\n", id, config)
}

func main() {
	numWorkers := 5
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker(i)
	}

	// The main goroutine waits for all workers to finish
	wg.Wait()
	fmt.Println("All workers finished. Main goroutine exiting.")
}
