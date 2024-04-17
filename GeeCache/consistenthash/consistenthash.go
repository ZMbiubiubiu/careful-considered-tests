package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// Hash 采取依赖注入的方式，允许用于替换成自定义的 Hash 函数，也方便测试时替换，默认为 crc32.ChecksumIEEE 算法。
type Hash func(data []byte) uint32

// Map 一致性hash函数的核心结构
type Map struct {
	hash     Hash
	replicas int            // 虚拟节点的倍数
	keys     []int          // 虚拟节点hash值组成的hash环
	hashMap  map[int]string // 虚拟节点hash值与节点的映射关系
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		hash:     fn,
		replicas: replicas,
		keys:     nil,
		hashMap:  make(map[int]string),
	}

	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}

	return m
}

// Add 添加真实节点
func (m *Map) Add(nodes ...string) {
	for _, node := range nodes {
		// 每个真实节点对应replicas个虚拟节点
		for i := 0; i < m.replicas; i++ {
			h := int(m.hash([]byte(strconv.Itoa(i) + node)))
			m.keys = append(m.keys, h)
			m.hashMap[h] = node
		}
	}
	sort.Ints(m.keys)
}

// Get 给定key返回需要的节点
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))

	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	idx = idx % len(m.keys)

	return m.hashMap[m.keys[idx]]
}
