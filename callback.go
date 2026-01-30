package main
import "fmt"

// Function that accepts a callback
func fetchData(callback func(string)) {
    fmt.Println("Fetching data...")
    callback("Data received") // Invoke the callback
}

func main() {
    // Passing an anonymous function as a callback
    fetchData(func(result string) {
        fmt.Println("Callback:", result)
    })
}
