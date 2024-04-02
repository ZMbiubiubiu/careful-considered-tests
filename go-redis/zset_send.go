package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var RedisCache = &redis.Client{}

func init() {
	RedisCache = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	//ping
	pong, err := RedisCache.Ping().Result()
	if err != nil {
		fmt.Println("ping error", err.Error())
		return
	}
	fmt.Println("ping result:", pong)
}

func main() {
	key := "zset-test"
	nowTs := time.Now().Unix()
	for i := 0; i < 10_0000; i++ {
		err := RedisCache.ZAdd(key, redis.Z{Score: float64(nowTs), Member: i}).Err()
		if err != nil {
			fmt.Println("zadd error", err.Error())
		}
	}
}
