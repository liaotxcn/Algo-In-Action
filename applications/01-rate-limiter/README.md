# 01 限流器

限流器用于保护下游服务，常见策略有令牌桶、漏桶、固定窗口与滑动窗口计数。

## 实现内容

| 实现 | 行为 | 突发流量 | 时间注入 |
| --- | --- | --- | --- |
| 令牌桶 TokenBucket | 按固定速率补充令牌 | 允许（受桶容量限制） | `Clock` 函数 |
| 滑动窗口 SlidingWindow | 统计最近一个窗口内的请求 | 不允许 | `Clock` 函数 |

## 设计要点

- 通过注入 `Clock`（Go）/ 可调用 `clock`（Python）让测试不依赖真实时间；
- 令牌桶的补充量按 `elapsed * refillPerSecond` 计算，并封顶到桶容量；
- 滑动窗口在每次判定前清理过期事件，使用原地过滤避免频繁分配；
- 所有状态由互斥锁保护，可直接用于并发场景。

## 复杂度

| 操作 | 令牌桶 | 滑动窗口 |
| --- | --- | --- |
| Allow | O(1) | O(window 内事件数) |
| 空间 | O(1) | O(limit) |

## 实现

- Go：[solution.go](go/solution.go) · [solution_test.go](go/solution_test.go)
- Python：[solution.py](python/solution.py) · [test_solution.py](python/test_solution.py)

## 运行

```bash
go test ./applications/01-rate-limiter/go
python -m pytest applications/01-rate-limiter
```

## 相关主题

- [03 栈与队列](../../topics/03-stack-queue/README.md)
- [08 流式统计](../08-stream-stats/README.md)
