package main

import (
    "fmt"
    "os"
)

func main() {
    fileInfo, err := os.Stat("book_data.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("File Name:", fileInfo.Name())
    fmt.Println("Size (bytes):", fileInfo.Size())
    fmt.Println("Last Modified:", fileInfo.ModTime())
    fmt.Println("Is Directory:", fileInfo.IsDir())
}
