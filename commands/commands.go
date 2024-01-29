package commands

import (
	"strings"

	"github.com/UsamaHameed/tiny-redis/parser"
	"github.com/UsamaHameed/tiny-redis/storage"
)

type Commands struct {
    count int
}

type RedisResponse struct {
    Response string
    Success  bool
    Errors   []string
}

func ProcessCommand(input string) RedisResponse {
    s := storage.New("/tmp/redis-server/data")
    p := parser.New(input, "\\r\\n")
    p.ParseRespString()

    if (len(p.Commands)) > 0 {
        command := strings.ToLower(p.Commands[0])

        switch command {
        case "ping":
            return Ping()
        case "echo":
            return Echo(p.Commands)
        case "set":
            return Set(s, p.Commands)
        case "get":
            return Get(s, p.Commands)
        }
    }

    msg := "no command provided"
    return RedisResponse{ Success: false, Errors: []string{msg} }
}

func Ping() RedisResponse {
    return RedisResponse{
        Success: true,
        Response: "PONG",
    }
}

func Echo(c []string) RedisResponse {
    if len(c) > 1 {
        return RedisResponse{ Success: true, Response: c[1] }
    }

    err := []string{"not enough args for the echo command"}
    return RedisResponse{ Success: false, Errors: err }
}

func Set(s *storage.Storage, c []string) RedisResponse {
    if len(c) > 2 {
        key := c[1]
        value := c[2]
        s.Store(key, value)

        return RedisResponse{ Success: true, Response: "OK"}
    }

    return RedisResponse{ Success: false, Response: "NOT_OK"}
}

func Get(s *storage.Storage, c []string) RedisResponse {
    if len(c) > 1 {
        key := c[1]

        value := s.Retrieve(key)
        return RedisResponse{ Success: true, Response: value }
    }

    return RedisResponse{ Success: false, Response: "" }
}

