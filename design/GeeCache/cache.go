// 拥有并发控制能力的LRU，缓存值用抽象的ByteView

package geecache

import (
	"sync"

	"geecache/lru"
)

type cache struct {
	mu         sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}

	v, ok := c.lru.Get(key)
	if !ok {
		return
	}

	return v.(ByteView), true
}

func (c *cache) set(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// 懒加载
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}
	c.lru.Set(key, value)
}
