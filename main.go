package main

import (
    "fmt"
    "os"
    "syscall"
    "os/signal"
    "github.com/UsamaHameed/tiny-redis/server"
)

func main() {
    fmt.Println("spawing a tcp server")
    s, err := server.SpawnServer()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    s.Start()

    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
    <-signalChan

    fmt.Println("killing the server")
    s.Stop()
    fmt.Println("gracefully shutdown the server")
}
