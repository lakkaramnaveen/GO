package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Alice", "Bob", "Vera"}
	n, found := slices.BinarySearch(names, "Alice")
	fmt.Println("Alice:", n, found)

	n, found = slices.BinarySearch(names, "Bill")
	fmt.Println("Bill:", n, found)
}