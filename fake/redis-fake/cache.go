package main

import (
	"context"
	"sync"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/go-redis/redis/v8"
)

var (
	rdb  *redis.Client
	once sync.Once
	ctx  = context.Background()
)

func InitRedisClient() {
	once.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr: beego.AppConfig.String("redis_addr"),
		})
	})
}

func GetString(key string) (value string, err error) {
	v, err := rdb.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return v, nil
}

func SetString(key, value string, time time.Duration) (err error) {
	return rdb.Set(ctx, key, value, time).Err()
}
