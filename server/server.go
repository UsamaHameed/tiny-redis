package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

type server struct {
    waitGroup   sync.WaitGroup
    listener    net.Listener
    connection  chan net.Conn
    shutdown    chan struct{}
}
func (s *server) handleConnection(conn net.Conn) {
    fmt.Println("connected to", conn.RemoteAddr().String())

    // serve the connection as long as the TCP client desires
    // todo: add a timeout?
    for {
        data, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            fmt.Println("unable to create a read buffer")
            //panic(err)
        }

        str := strings.TrimSpace(string(data))
        fmt.Println("received tcp packet", str)

        if str == "STOP" {
            fmt.Println("closed connection with", conn.RemoteAddr().String())
            break
        }

        res := "hello world from the server"

        fmt.Println("responding with", res, "to", conn.RemoteAddr().String())
        conn.Write([]byte(res))

    }
    conn.Close()

}

func (s *server) acceptConnections() {
    defer s.waitGroup.Done()

    for {
        select {
        case <- s.shutdown:
            return
        default:
            conn, err := s.listener.Accept()
            if err != nil {
                continue
            }
            s.connection <- conn
        }
    }
}

func (s *server) handleConnections() {
    defer s.waitGroup.Done()

    for {
        select {
        case <- s.shutdown:
            return
        case conn := <- s.connection:
            go s.handleConnection(conn)
        }
    }
}

func (s *server) Start() {
    s.waitGroup.Add(2)
    go s.acceptConnections()
    go s.handleConnections()
}

func (s *server) Stop() {
    close(s.shutdown)
    s.listener.Close()
}

func SpawnServer() (*server, error) {
    // default port in the redis server is 6379
    ln, err := net.Listen("tcp", ":6379")

    if err != nil {
        return nil, fmt.Errorf("unable to bind to port 6379")
    }

    return &server {
        listener: ln,
        shutdown: make(chan struct{}),
        connection: make(chan net.Conn),
    }, nil
}
