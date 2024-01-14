package storage

// todo:
// handle saving to memory,
// save to file
// handle expiry time

type Storage struct {}

func New() *Storage {
    return &Storage{}
}

type KeyValueStore struct {
    Data map[string] string
}

var redisData KeyValueStore

func (s *Storage) Init(initData map[string]string) {
    redisData.Data = make(map[string]string)
    for k, v := range initData {
        redisData.Data[k] = v
    }
}

func (s *Storage) Store(key string, value string) bool {
    redisData.Data[key] = value
    return true
}

func (s *Storage) Retrieve(key string) string {
    return redisData.Data[key]
}

