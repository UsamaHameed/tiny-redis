package storage

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
)

type Storage struct {
    redisData KeyValueStore
    filename string

    // just to make decodeRedisData work, this should be removed
    encodedData bytes.Buffer
}

func New(filename string) *Storage {
    return &Storage{ filename: filename }
}

type KeyValueStore struct {
    Data map[string] string
}

func (s *Storage) Init(initData map[string]string) {
    s.redisData.Data = make(map[string]string)
    for k, v := range initData {
        s.redisData.Data[k] = v
    }
}

func (s *Storage) Store(key string, value string) bool {
    s.redisData.Data[key] = value
    return true
}

func (s *Storage) Retrieve(key string) string {
    return s.redisData.Data[key]
}

func (s *Storage) encodeRedisData() (*bytes.Buffer, error) {
    encoded, err := encodeData(s.redisData.Data)

    // just to make decodeRedisData work, this should be removed
    s.encodedData = *encoded

    return encoded, err
}



func (s *Storage) decodeRedisData(buff *bytes.Buffer) (map[string]string, error) {
    return decodeData(buff)
}

func (s *Storage) SaveRedisDataToDisk() error {
    data, err := s.encodeRedisData()

    if err != nil {
        return err
    }

    savedBytes, saveErr := saveToDisk(s.filename, data)

    if saveErr != nil {
        if savedBytes == 0 {
            e := errors.New("unable to save data to file")
            return errors.Join(e, saveErr)
        }
        return saveErr
    }

    return nil // success
}

func (s *Storage) LoadRedisDataFromDisk() error {
    var buff bytes.Buffer

    readBytes, readErr := readFromDisk(s.filename, &buff)

    fmt.Println(readBytes, readErr)
    data, err := s.decodeRedisData(&buff)
    fmt.Println(data, err)
    fmt.Println(data)

    if err != nil {
        return err
    }

    if readErr != nil {
        if readBytes == 0 {
            e := errors.New("unable to read data from the file")
            return errors.Join(e, readErr)
        }
        return readErr
    }

    return nil // success
}

func saveToDisk(filename string, data *bytes.Buffer) (int, error) {
    file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

    defer file.Close()

    if err != nil {
        panic(err)
    }

    writeBuff := data.Bytes()

    //file.WriteString(fmt.Sprint(data.Len(), "\n"))
    return file.Write(writeBuff)
}

func readFromDisk(filename string, buff *bytes.Buffer) (int, error) {
    file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666)
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)

    count := 0
    for scanner.Scan() {
        for _, b := range scanner.Bytes() {
            count++
            buff.WriteByte(b)
        }
    }

    if count == 0 {
        return -1, errors.New("unable to read file")
    }

    return count, nil
}

func encodeData(data map[string]string) (*bytes.Buffer, error) {
    buffer := new(bytes.Buffer)
    encoder := gob.NewEncoder(buffer)
    err := encoder.Encode(data)

    //fmt.Printf("buff %#v\n", buffer)
    // todo: properly log this

    return buffer, err
}

func decodeData(buff *bytes.Buffer) (map[string]string, error) {
    var decodedMap map[string]string

    decoder := gob.NewDecoder(buff)
    err := decoder.Decode(&decodedMap)

    return decodedMap, err
}
