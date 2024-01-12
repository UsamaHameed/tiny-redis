package storage

// todo:
// handle saving to memory,
// save to file
// handle expiry time

type KeyValueStore struct {
    data map[string] string
}

var redisData KeyValueStore

func Init(initData KeyValueStore) {
    redisData.data = make(map[string]string)
    for k, v := range initData.data {
        redisData.data[k] = v
    }
}

func Store(key string, value string) bool {
    redisData.data[key] = value
    return true
}

func Retrieve(key string) string {
    return redisData.data[key]
}

