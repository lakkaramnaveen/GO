package main

import "fmt"

func measure(s Shape) {
    fmt.Println("Shape:", s)
    fmt.Println("Area:", s.Area())
    fmt.Println("Perimeter:", s.Perimeter())
}

func main() {
    r := Rectangle{Width: 3, Height: 4}

    // A Rectangle is passed to a function expecting a Shape interface
    measure(r)
}
