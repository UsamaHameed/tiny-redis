package utils

import (
	"strconv"
)

func ParseStringToInt(s string) (int, error) {
    i, err := strconv.ParseInt(s, 10, 64)
    return int(i), err
}

func ParseByteToInt(b []rune) (int, error) {
    i, err := strconv.ParseInt(string(b), 10, 64)
    return int(i), err
}
