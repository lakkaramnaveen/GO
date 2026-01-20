package main

import "fmt"

func add (x int, y int) int {
	return x + y
}

func multiply (x int, y int) int {
	return x * y
}

func divide (x int, y int) int {
	return x / y
}

func main() {
	fmt.Println(add(2, 2))
	fmt.Println(multiply(2, 2))
	fmt.Println(divide(2, 2))
}