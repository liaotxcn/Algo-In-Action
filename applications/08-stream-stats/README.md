# 08 流式统计

数据流场景无法保存全部历史，本应用演示三种固定内存的统计结构。

## Count-Min Sketch

- 使用 `depth` 行、`width` 列的计数器矩阵；
- 每个元素在每行命中一个桶并加一；
- 查询取所有行中的最小值；
- **只会高估，不会低估**，宽度越大误差越小；
- 适合统计热点 Key、检测流量异常。

## 滑动平均

- 固定长度环形缓冲保存最近 `capacity` 个值；
- 同时维护窗口和 `sum`，`Add` 与 `Average` 均为 O(1)；
- 窗口满后新值覆盖最旧值。

## 蓄水池抽样

- 前 `capacity` 个元素直接进入样本；
- 第 `i` 个元素以 `capacity / i` 的概率替换样本中的随机位置；
- 每个元素被选中的概率相等；
- 注入 `seed` 后结果可复现，便于测试。

## 复杂度

| 结构 | Add | 查询 | 空间 |
| --- | --- | --- | --- |
| Count-Min Sketch | O(depth) | O(depth) | O(width × depth) |
| MovingAverage | O(1) | O(1) | O(capacity) |
| ReservoirSampler | O(1) | O(capacity) | O(capacity) |

## 实现

- Go：[solution.go](go/solution.go) · [solution_test.go](go/solution_test.go)
- Python：[solution.py](python/solution.py) · [test_solution.py](python/test_solution.py)

## 测试覆盖

- 频次估计不低于真实值、未出现元素估计为 0；
- 增量更新单调；
- 滑动窗口滚动与空窗口；
- 相同种子的抽样结果一致、样本大小与取值范围。

## 相关主题

- [04 哈希表](../../topics/04-hash-table/README.md)
- [03 TopK 排行](../03-topk-ranking/README.md)
- [01 限流器](../01-rate-limiter/README.md)
