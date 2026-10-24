# 05 文本搜索

本应用演示搜索引擎的两个基础组件：倒排索引与前缀树。

## 倒排索引

```text
单词 -> {文档 ID 集合}
```

- 写入时先分词（转小写、按非字母数字切分），再更新词到文档的映射；
- 重复写入同一文档会先移除旧词条，避免过期索引；
- 查询结果按文档 ID 升序返回，读写由 `sync.RWMutex` / `RLock` 保护。

## 前缀树 Trie

- 每个节点保存子节点映射与 `terminal` 标记；
- `Complete` 深度优先遍历，子节点按字符排序，因此结果天然字典序；
- 支持 `limit` 参数限制返回数量，适用于搜索框自动补全。

## 复杂度

| 操作 | 复杂度 |
| --- | --- |
| 建索引 | O(文档总词数) |
| 单词查询 | O(命中文档数 log 命中文档数) |
| Trie 插入/精确查找 | O(单词长度) |
| 前缀补全 | O(前缀长度 + 结果数 × 单词长度) |

## 实现

- Go：[solution.go](go/solution.go) · [solution_test.go](go/solution_test.go)
- Python：[solution.py](python/solution.py) · [test_solution.py](python/test_solution.py)

## 测试覆盖

- 多文档检索、大小写不敏感、未命中单词；
- 同一文档重新索引后旧词被清理；
- Trie 精确匹配、前缀补全、字典序与 limit。

## 相关主题

- [04 哈希表](../../topics/04-hash-table/README.md)
- [13 字符串](../../topics/13-string/README.md)
