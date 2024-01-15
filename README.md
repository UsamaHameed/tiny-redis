# tiny-redis

## Introduction
This is my solution to the redis server implementation challenge from
https://codingchallenges.fyi/challenges/challenge-redis/

## Redis
Redis is an open source in memory key-value data store that can be used
in the following ways:
- caching
- real time chat apps
- queues
- leaderboards
- session store
- geospatial data
...and many more!

## How does Redis work?
Redis uses the RESP protocol for communication.

### What is RESP?
Redis clients use a protocol called Redis Serialization Protocol (RESP).
RESP can serialize different data types including integers, strings, and
arrays. It also features an error-specific type. A client sends a request
to the Redis server as an array of strings. The array’s contents are the
command and its arguments that the server should execute. The server’s
reply type is command-specific. RESP is binary-safe and uses prefixed
length to transfer bulk data so it does not require processing bulk data
transferred from one process to another.

# Installation

## with curl
```bash
curl -OL https://github.com/UsamaHameed/tiny-redis/releases/latest/download/tiny-redis
```

## build from source
```bash
git clone git@github.com:UsamaHameed/tiny-redis.git
go build
```

# How to test

```bash
# start the server
go run main.go

# connect to redis via tcp and then try sending any of the resp strings
nc localhost 6379
# resp strings to try
+PING\r\n
$12\r\nRedis\r\n
:10\r\n
*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n
*2\r\n$3\r\nGET\r\n$5\r\nmykey\r\n
```

# TODO
- a cron job to build since artifacts get deleted in 90 days?
- add testing guide with files `nc localhost 6379 < file`
