package test

import (
	_ "fmt"
	_ "net"
	"testing"

	_ "github.com/UsamaHameed/tiny-redis/server"
)

func TestServer(t *testing.T) {
    //server, err := server.SpawnServer()

    //if err != nil {
    //    t.Fatal(err)
    //}

    //server.Start()
    //conn, err := net.Dial("tcp", "localhost:6379")

    //if err != nil {
    //    t.Fatal(err)
    //}
    //defer conn.Close()

    //expected := "usama helped me write"
    //actual := make([]byte, len(expected))

    //conn.Write([]byte("hello from the test"))

    //if _, err := conn.Write(actual); err != nil {
    //    fmt.Println("error")
    //    t.Fatal(err)
    //}

    //if string(actual) != expected {
    //    t.Errorf("expected %q but got %q", expected, actual)
    //}

    //server.Stop()

}
