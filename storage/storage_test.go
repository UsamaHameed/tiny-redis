package storage

import "testing"

func TestStore(t *testing.T) {
    s := New()
    s.Init(map[string]string{})

    res := s.Store("Name", "John")

    if !res {
        t.Fatalf("unable to store")
    }
}

func TestRetrieve(t *testing.T) {
    s := New()
    initData := make(map[string]string)
    initData["Name"] = "John"

    s.Init(initData)

    res := s.Retrieve("Name")

    if res != "John" {
        t.Fatalf("retrieve does not work, expected=%s, got=%s", "John", res)
    }
}
