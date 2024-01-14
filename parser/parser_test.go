package parser

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
    tests := []struct{
        input               string
        error               bool
        expectedLength      int
        expectedCommands    []string
    }{
        {
            input: "*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n",
            expectedLength: 3, expectedCommands: []string{"SET", "mykey", "myvalue"},
        },
        {
            input: "*2\r\n$3\r\nGET\r\n$5\r\nmykey\r\n",
            expectedLength: 2, expectedCommands: []string{"GET", "mykey"},
        },
        {
            input: "*1\r\n$4\r\nping\r\n",
            expectedLength: 1, expectedCommands: []string{"ping"},
        },
        {
            input: "*2\r\n$4\r\necho\r\n$11\r\nhello world\r\n",
            expectedLength: 2, expectedCommands: []string{"echo", "hello world"},
        },
        {
            input: "+OK\r\n",
            expectedLength: 1, expectedCommands: []string{"OK"},
        },
        {
            input: "$0\r\n\r\n",
            expectedLength: 1, expectedCommands: []string{""},
        },
        //{
        //    input: "+hello world\r\n",
        //    error: true,
        //    expectedLength: 0, expectedCommands: []string{""},
        //},
    }

    for _, test := range tests {
        p := New(test.input)
        err := p.ParseRespString()

        if err != nil {
            t.Fatalf("ParseRespString returned an error, %s", err)
        }

        if len(p.commands) != test.expectedLength {
            t.Fatalf("Parser returned wrong length, expected=%d, got=%d",
            test.expectedLength, len(p.commands))
        }

        if !reflect.DeepEqual(p.commands, test.expectedCommands) {
            t.Fatalf("Parser returned wrong commands, expected=%v, got=%v",
            test.expectedCommands, p.commands)
        }
    }
}

func TestAdvancePointer(t *testing.T) {
    input := "*3\r\n$3\r\nSET\r\n"
    p := New(input)

    if p.currPos != 0 {
        t.Fatalf("incorrect initial currPos, expected=%d, got=%d", 0, p.currPos)
    }

    p.advancePointer()
    if p.currPos != 4 {
        t.Fatalf("incorrect currPos, expected=%d, got=%d", 4, p.currPos)
    }

    p.advancePointer()
    if p.currPos != 8 {
        t.Fatalf("incorrect currPos, expected=%d, got=%d", 8, p.currPos)
    }
}

func TestParseSimpleString(t *testing.T) {
    input := "+PING\r\n"

    p := New(input)
    err := p.parseSimpleString()

    if err != nil { // command.PING is 0
        t.Fatalf("ParseSimpleString returned an error: %v", err)
    }
}

func TestParseBulkString(t *testing.T) {
    input := "$3\r\nSETxx"

    p := New(input)
    err := p.ParseBulkString()

    if err != nil {
        t.Fatalf("ParseBulkString returned an error, %s", err)
    }
    if p.commands[0] != "SET" {
        t.Fatalf("ParseBulkString did not parse correctly, expected=%s, got=%s",
        "SET", p.commands[0])
    }
}


func TestParseArray(t *testing.T) {
    input := "*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n"

    p := New(input)
    err := p.ParseArray()

    if err != nil {
        t.Fatalf("ParseArray returned an error, %s", err)
    }

    if len(p.commands) != 3 {
        t.Fatalf("ParseArray returned wrong length, expected=%d, got=%d",
        3, len(p.commands))
    }

    expected := []string{"SET", "mykey", "myvalue"}
    if !reflect.DeepEqual(p.commands, expected) {
        t.Fatalf("ParseArray returned wrong length, expected=%d, got=%d",
        3, len(p.commands))
    }
}

func TestParseInt(t *testing.T) {
    input := ":10\r\n"
    p := New(input)
    err := p.ParseInt()

    if err != nil {
        t.Fatalf("ParseInt returned an error, %s", err)
    }

    if p.commands[0] != "10" {
        t.Fatalf("ParseInt returned wrong int, expected=%s, got=%s",
        "10", p.commands[0])
    }
}

func TestParseSize(t *testing.T) {
    input := "*10\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n"
    p := New(input)
    p.currPos = 1
    res := p.ParseSize()

    if res != 10 {
        t.Fatalf("ParseSize returned wrong int, expected=%d, got=%d",
        10, res)
    }
}
