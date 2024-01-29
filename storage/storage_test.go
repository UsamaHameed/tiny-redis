package storage

import (
	"bytes"
	"os"
	"testing"
)

func TestStore(t *testing.T) {
    s := New("/tmp/data")
    s.Init(map[string]string{})

    res := s.Store("Name", "John")

    if !res {
        t.Fatalf("unable to store")
    }
}

func TestRetrieve(t *testing.T) {
    s := New("/tmp/data")
    initData := make(map[string]string)
    initData["Name"] = "John"

    s.Init(initData)

    res := s.Retrieve("Name")

    if res != "John" {
        t.Fatalf("retrieve does not work, expected=%s, got=%s", "John", res)
    }
}

func TestEncode(t *testing.T) {
    data := make(map[string]string)
    data["key1"] = "value1"
    data["key2"] = "value2"

    encoded, err := encodeData(data)

    if err != nil {
        t.Fatalf("encodeData returned an error: %s", err)
    }

    if encoded.Len() != 43 {
        t.Fatalf("encodeData has wrong length, returned: %d", encoded.Len())
    }

    // todo make the following test work
    //returned := bytes.TrimSpace(encoded.Bytes())
    //e := string(returned)

    //if e != "" {
    //    t.Fatalf("encodeData.String() returned the wrong value, returned: %s", e)
    //}
}

func TestDecode(t *testing.T) {
    // {key: value}
    buff := bytes.NewBufferString("\r\x7f\x04\x01\x02\xff\x80\x00\x01\f\x01\f\x00\x00\x0e\xff\x80\x00\x01\x03key\x05value")
    decoded, err := decodeData(buff)

    if err != nil {
        t.Fatalf("decodeData returned an error: %s", err)
    }

    if decoded["key"] != "value" {
        t.Fatalf("decodeData returned the wrong value %s", decoded["key"])
    }

    if len(decoded) != 1 {
        t.Fatalf("decodeData returned the wrong no of keys, %d", len(decoded))
    }
}

func TestSaveToDisk(t *testing.T) {
    buff := bytes.NewBufferString("keyvalue")
    s, err := saveToDisk("/tmp/test", buff)
    if err != nil {
        t.Fatalf("saveToDisk returned an error: %s", err)
    }

    if s != 8 {
        t.Fatalf("saveToDisk returned the wrong value: %d", s)
    }

    os.Remove("tmp/test")
}

func TestReadFromDisk(t *testing.T) {
    os.WriteFile("/tmp/test", []byte("100\nhello world"), 0666)

    var buff bytes.Buffer
    s, err := readFromDisk("/tmp/test", &buff)

    if err != nil {
        t.Fatalf("readFromDisk returned an error: %s", err)
    }

    if s != 14 {
        t.Fatalf("readFromDisk returned the wrong value: %d", s)
    }

    // cleanup
    os.Remove("/tmp/test")
}

func TestSaveRedisDataToDisk(t *testing.T) {
    data := make(map[string]string)
    data["key1"] = "value1"
    data["key2"] = "value2"

    s := New("/tmp/test")
    s.Init(data)

    err := s.SaveRedisDataToDisk()

    if err != nil {
        t.Fatalf("SaveRedisDataToDisk returned an error: %s", err)
    }

    // cleanup, /tmp/data is created by saveToDisk
    os.Remove("/tmp/test")
}

func TestLoadRedisDataFromDisk(t *testing.T) {
    str := "\r\x7f\x04\x01\x02\xff\x80\x00\x01\f\x01\f\x00\x00\x0e\xff\x80\x00\x01\x03key\x05value"
    os.WriteFile("/tmp/test", []byte(str), 0666)

    s := New("/tmp/test")

    err := s.LoadRedisDataFromDisk()

    if err != nil {
        t.Fatalf("LoadRedisDataFromDisk returned an error: %s", err)
    }

    // cleanup
    os.Remove("/tmp/test")
}
