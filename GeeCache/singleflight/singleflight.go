package singleflight

import "sync"

type call struct {
	wg    sync.WaitGroup
	value interface{}
	err   error
}

type Group struct {
	mu sync.Mutex // protects m
	m  map[string]*call
}

// Do 请求key期间，无论过来多少请求，只实际调用依此fn
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	// 懒加载
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	c, ok := g.m[key]
	// 请求已经存在，等待请求结果
	if ok {
		g.mu.Unlock()
		c.wg.Wait() // 已经创建了call，等待执行结果
		return c.value, c.err
	}

	c = new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	// 发起请求
	c.value, c.err = fn()
	// 请求结束
	c.wg.Done()

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.value, c.err
}
