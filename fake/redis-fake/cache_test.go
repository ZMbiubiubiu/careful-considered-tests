package main

import (
	"os"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

var (
	// redis-fake server
	miniRedis *miniredis.Miniredis
)

func TestMain(m *testing.M) {
	// 启动redis-fake服务
	miniRedis, _ = miniredis.Run()

	// 连接redis-fake服务
	rdb = redis.NewClient(&redis.Options{
		Addr: miniRedis.Addr(),
	})
	os.Exit(m.Run())
}

func TestCache(t *testing.T) {
	// 测试未过期
	err := SetString("foo", "bar", 10*time.Second)
	assert.Nil(t, err)
	got, err := GetString("foo")
	assert.Nil(t, err)
	assert.Equal(t, "bar", got)

	// 测试过期
	err = SetString("foo", "bar", 10*time.Second)
	assert.Nil(t, err)
	// 快进11s
	miniRedis.FastForward(11 * time.Second)
	got, err = GetString("foo")
	assert.Nil(t, err)
	assert.Equal(t, "", got)
}
