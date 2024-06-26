# 序言
> 商业世界里，现金为王；架构世界里，缓存为王。

本项目的分布式缓存，涉及并发控制、淘汰策略、分布式节点通信等方面。

* [x] 单机缓存和基于 HTTP 的分布式缓存 
* [x] 最近最少访问(Least Recently Used, LRU) 缓存策略
* [x] 使用 Go 锁机制防止缓存击穿
* [x] 使用一致性哈希选择节点，实现负载均衡
* [x] 分布式节点通信
* [x] singleflight防止缓存击穿
* [ ] 使用 protobuf 优化节点间二进制通信

# 代码结构

```shell
geecache/
    |--consistenthash/
        |--consistenthash.go // 一致性hash实现
    |--lru/
        |--lru.go  // lru 缓存淘汰策略
    |--byteview.go // 缓存值的抽象与封装
    |--cache.go    // 并发控制
    |--geecache.go // 负责与外部交互,控制缓存存储和获取的主流程
    |--http.go     // 提供被其他节点访问的能力(http)
```

# 业务逻辑
```shell
                            是
接收 key --> 检查是否被缓存 -----> 返回缓存值
                |  否                         是
                |-----> 是否应当从远程节点获取 -----> 与远程节点交互 --> 返回缓存值 ⑵
                            |  否
                            |-----> 调用`回调函数`,获取值并添加到缓存 --> 返回缓存值
```