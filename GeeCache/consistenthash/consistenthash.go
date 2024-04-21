package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// Hash 采取依赖注入的方式，默认为 crc32.ChecksumIEEE 算法
// 允许用于替换成自定义的 Hash 函数，方便测试时替换
type Hash func(data []byte) uint32

// Map 一致性hash函数的核心结构
type Map struct {
	hash      Hash
	replicas  int            // 节点的倍数，一个节点可产生replicas个虚拟节点，解决数据倾斜问题
	vnodes    []int          // 虚拟节点hash值组成的hash环
	hash2node map[int]string // 虚拟节点hash值与节点的映射关系
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		hash:      fn,
		replicas:  replicas,
		vnodes:    nil,
		hash2node: make(map[int]string),
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
			vnode := strconv.Itoa(i) + node
			h := int(m.hash([]byte(vnode)))
			m.vnodes = append(m.vnodes, h)
			m.hash2node[h] = node
		}
	}

	// 排序很重要
	sort.Ints(m.vnodes)
}

// Get 给定key返回需要的节点
func (m *Map) Get(key string) string {
	if len(m.vnodes) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))

	idx := sort.Search(len(m.vnodes), func(i int) bool {
		return m.vnodes[i] >= hash
	})

	idx = idx % len(m.vnodes)

	return m.hash2node[m.vnodes[idx]]
}
