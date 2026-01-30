package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {

    signals := make(chan os.Signal, 1)

    signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

    done := make(chan bool, 1)

    go func() {

        sig := <-signals
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()

    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}