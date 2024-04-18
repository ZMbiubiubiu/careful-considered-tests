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
接收 key --> 检查是否被缓存 -----> 返回缓存值 ⑴
                |  否                         是
                |-----> 是否应当从远程节点获取 -----> 与远程节点交互 --> 返回缓存值 ⑵
                            |  否
                            |-----> 调用`回调函数`，获取值并添加到缓存 --> 返回缓存值 ⑶
```