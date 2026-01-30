package main

import "fmt"

func calculate(a, b int) (int, int) {
    sum := a + b
    product := a * b
    return sum, product
}

func main() {
    // Assigning both return values
    s, p := calculate(5, 3)
    fmt.Println("Sum:", s, "Product:", p)
}
