# 04 任务调度与分布式 ID

本应用包含两个调度基础设施组件：Snowflake ID 生成器与任务调度器。

## Snowflake ID

64 位结构：

```text
| 1 位符号 | 41 位毫秒时间戳 | 10 位机器号 | 12 位序列号 |
```

- 单机每毫秒最多生成 4096 个 ID；
- 同一毫秒内序列号递增，跨毫秒归零；
- 时钟回拨、序列号耗尽返回错误，而不是阻塞或生成重复 ID；
- 时间源可注入，便于测试。

## 任务调度

| 调度器 | 策略 | 复杂度 |
| --- | --- | --- |
| `Scheduler` | 轮询分配 | 分配 O(1) |
| `PriorityScheduler` | 优先级高者优先，同优先级 FIFO | 提交/取出 O(log n) |

优先级队列使用堆实现，通过提交序号保证相同优先级的稳定性。

## 实现

- Go：[solution.go](go/solution.go) · [solution_test.go](go/solution_test.go)
- Python：[solution.py](python/solution.py) · [test_solution.py](python/test_solution.py)

## 测试覆盖

- 同毫秒 ID 递增、跨毫秒序列号重置；
- 机器号解析、非法机器号、时钟回拨；
- 轮询顺序与空节点错误；
- 优先级顺序与同优先级 FIFO。

## 相关主题

- [06 堆](../../topics/06-heap/README.md)
- [11 贪心](../../topics/11-greedy/README.md)
- [01 限流器](../01-rate-limiter/README.md)
