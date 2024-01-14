package commands

import (
	"fmt"
	"testing"

	"github.com/UsamaHameed/tiny-redis/storage"
)

func TestEcho(t *testing.T) {
    c := []string{"echo", "hello"}
    res := Echo(c)

    if !res.Success {
        errs := ""
        for _, e := range res.Errors {
            errs += e
        }
        t.Fatalf("Echo command failed with errors=%s", errs)
    }

    if res.Response != "hello" {
        errs := ""
        for _, e := range res.Errors {
            errs += e
        }
        t.Fatalf("Echo command did not return correct response, expected=%s, got=%s",
            "hello", res.Response)
    }

}
func TestSet(t *testing.T) {
    c := []string{"set", "key", "value"}
    s := storage.New()
    s.Init(map[string]string{})

    res := Set(s, c)

    if !res.Success {
        errs := ""
        for _, e := range res.Errors {
            errs += e
        }
        t.Fatalf("Set command failed with errors=%s", errs)
    }

    if res.Response != "OK" {
        errs := ""
        for _, e := range res.Errors {
            errs += e
        }
        t.Fatalf("Set command did not return correct response, expected=%s, got=%s",
            "value", res.Response)
    }

}
func TestGet(t *testing.T) {
    c := []string{"get", "key"}
    s := storage.New()
    s.Init(map[string]string{"key": "value"})

    res := Get(s, c)
    fmt.Println("res", res)

    if !res.Success {
        errs := ""
        for _, e := range res.Errors {
            errs += e
        }
        t.Fatalf("Get command failed with errors=%s", errs)
    }

    if res.Response != "value" {
        errs := ""
        for _, e := range res.Errors {
            errs += e
        }
        t.Fatalf("Get command did not return correct response, expected=%s, got=%s",
            "value", res.Response)
    }

}

func TestExists(t *testing.T) {

}

func TestDel(t *testing.T) {

}

func TestIncr(t *testing.T) {

}

func TestLPush(t *testing.T) {

}

func TestRPush(t *testing.T) {

}

func TestSave(t *testing.T) {

}
