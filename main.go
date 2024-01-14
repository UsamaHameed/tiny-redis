package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/UsamaHameed/tiny-redis/server"
	"github.com/UsamaHameed/tiny-redis/storage"
)

func main() {
    fmt.Println("spawing a tcp server")
    s, err := server.SpawnServer()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    s.Start()
    store := storage.New()
    store.Init(make(map[string]string))

    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
    <-signalChan

    fmt.Println("killing the server")
    s.Stop()
    fmt.Println("gracefully shutdown the server")
}
