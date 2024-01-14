package commands

import (
	"fmt"

	"github.com/UsamaHameed/tiny-redis/parser"
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

func ParseCommand(input string) []string {
    //chunks := strings.Split(input, " ")
    //comm := chunks[0]

    p := parser.New(input)
    p.ParseRespString()
    fmt.Println("commands", p.Commands)

    return p.Commands
}

type CommandType string
const (
    PING    CommandType = "PING"
    ECHO    CommandType = "ECHO"
    SET     CommandType = "SET"
    GET     CommandType = "GET"
    EXISTS  CommandType = "EXISTS"
    OK      CommandType = "OK"
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

