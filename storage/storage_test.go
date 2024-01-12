package storage

import "testing"

func setupTests(d KeyValueStore) {
    Init(d)
}

func TestStore(t *testing.T) {
    data := make(map[string]string)
    initdata := KeyValueStore{data:data}

    setupTests(initdata)

    res := Store("Name", "John")

    if !res {
        t.Fatalf("unable to store")
    }
}

func TestRetrieve(t *testing.T) {
    data := make(map[string]string)
    data["Name"] = "John"
    initdata := KeyValueStore{data:data}
    setupTests(initdata)

    res := Retrieve("Name")

    if res != "John" {
        t.Fatalf("retrieve does not work, expected=%s, got=%s", "John", res)
    }
}
