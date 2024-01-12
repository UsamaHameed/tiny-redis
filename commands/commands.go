package commands

import (
	"fmt"
	"strings"

	"github.com/UsamaHameed/tiny-redis/storage"
)

type Commands struct {
    count int
}

type ParseCommandResponse struct {
    Response string
    Success  bool
    Errors   []string
}

func ParseCommand(input string) *ParseCommandResponse {
    chunks := strings.Split(input, " ")
    comm := chunks[0]

    if comm == "PING" {
        return &ParseCommandResponse{ Response: Ping(), Success: true}
    } else if comm == "ECHO" {
        return &ParseCommandResponse{ Response: Echo(chunks[1]), Success: true}
    } else if comm == "SET" {
        return &ParseCommandResponse{
            Response: "", Success: Set(chunks[1], chunks[2]),
        }
    } else {
        return &ParseCommandResponse{
            Success: false, Response: "", Errors: []string{
                fmt.Sprintf("unknown command %s", comm),
            },
        }
    }
}

const (
    PING = iota
    ECHO
    SET
    GET
)

func Ping() string {
    return "PONG"
}

func Echo(input string) string {
    return input
}

func Set(key string, value string) bool {
    storage.Store(key, value)
    return true
}

func Get(key string) string {
    return storage.Retrieve(key)
}

