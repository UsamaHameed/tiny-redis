package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/UsamaHameed/tiny-redis/utils"
)

// had to escape the \r and \n chars so that I could test using
// nc on the command line
const DEFAULT_DELIMITER = "\r\n"

type Parser struct {
    currPos int
    input string
    delimiter string
    Commands []string
}

func New(str string, delimiter string) *Parser {
    p := Parser{ currPos: 0, input: str, delimiter: delimiter }
    return &p
}

func (p *Parser) advancePointer() {
    str := p.input[p.currPos:]
    index := strings.Index(str, p.delimiter)

    if index != -1 {
        p.currPos = index + p.currPos + len(p.delimiter)
    } else {
        p.currPos = len(p.input)
    }
}

func (p *Parser) incrementPointer() {
    p.currPos += 1
}

func (p *Parser) appendParsedString(str string) {
    p.Commands = append(p.Commands, str)
}

func (p *Parser) trimQuotes() {
    length := len(p.input)
    if (p.input[0] == '"' && p.input[length - 1] == '"') {
        p.input = p.input[1:length - 1]
    }
}

func (p *Parser) ParseRespString() error {
    //fmt.Println("input to ParseRespString", p.input)
    if len(p.input) > 0 {
        p.trimQuotes()
        respType := p.input[p.currPos]
        str := p.input[1:]

        switch respType {
        case '+':
            return p.parseSimpleString()
        case '$':
            return p.ParseBulkString()
        case ':':
            return p.ParseInt()
        case '*':
            return p.ParseArray()
        case '-':
            return p.parseError(str)
        }

        return  errors.New(
            fmt.Sprintf("unsupported redis command type: %s", string(respType)),
        )
    }
    return errors.New(fmt.Sprintf("zero length"))
}

// simple strings can only be ok, ping and pong
func (p *Parser) parseSimpleString() error {
    p.incrementPointer() // skip the type
    start := p.currPos
    end := strings.Index(p.input[start:], p.delimiter)
    str := strings.ToLower(p.input[p.currPos:end + 1])

    if str == "ping" {
        p.appendParsedString("PING")
        return nil
    } else if str == "echo" {
        p.appendParsedString("ECHO")
        return nil
    } else if str == "ok" {
        p.appendParsedString("OK")
        return nil
    }

    return errors.New(fmt.Sprintf("unsupported simple string: %s", p.input[start:]))
}

func (p *Parser) ParseSize() int {
    current := p.input[p.currPos:]
    start := 0
    end := strings.Index(current, p.delimiter)
    input := current[start:end]
    size, err := utils.ParseByteToInt([]rune(input))

    if err != nil {
        e := errors.New(fmt.Sprint("unable to find size", input))
        joinedErr := errors.Join(e, err)
        panic(joinedErr)
    }
    return size
}

func (p *Parser) ParseBulkString() error {
    p.incrementPointer() // skip the type
    input := p.input[p.currPos:]
    size := p.ParseSize()
    index := strings.Index(input, p.delimiter)

    start := index + len(p.delimiter)
    end := start + size

    comm := input[start:end]
    if size != len(input[start:end]) {
        panic(fmt.Sprintf("invalid bulk string=%s", input))
    }

    //fmt.Println("parsing bulk string", comm)

    p.advancePointer()
    p.appendParsedString(comm)

    return nil // no error
}

func (p *Parser) ParseInt() error {
    start := 1 // skip the type
    end := strings.Index(p.input[1:], p.delimiter)

    input := p.input[start:end + 1]
    i, e := strconv.ParseInt(input, 10, 64)

    if e != nil {
        err := errors.New(fmt.Sprintf("unable to parse %s to int", input))
        joinedErr := errors.Join(e, err)
        return joinedErr
    }

    p.appendParsedString(fmt.Sprint(i))
    return nil
}

func (p *Parser) ParseArray() error {
    //fmt.Println("parsing array", p.input)
    p.incrementPointer() // skip the type
    size := p.ParseSize()

    // offset size
    p.advancePointer()
    for i := 0; i < int(size); i++ {
        p.ParseRespString()
        p.advancePointer()
    }

    return nil
}

func (p *Parser) parseError(input string) error {
    return errors.New(fmt.Sprint("method not implemented", input))
}
