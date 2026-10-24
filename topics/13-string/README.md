# 13 字符串

字符串问题的核心是「窗口」与「匹配」。滑动窗口处理子串统计，KMP 处理模式匹配，中心扩展处理回文。

## 实现内容

| 问题 | 算法 | 时间复杂度 | 空间复杂度 |
| --- | --- | --- | --- |
| 无重复字符的最长子串 | 滑动窗口 + 最近位置表 | O(n) | O(字符集) |
| 找到所有异位词 | 定长窗口 + 计数表 | O(n + m) | O(字符集) |
| KMP 模式匹配 | 前缀函数 next 数组 | O(n + m) | O(m) |
| 最长回文子串 | 中心扩展 | O(n²) | O(1) |
| 验证回文串 | 双指针 + 字符过滤 | O(n) | O(1)（Go）/ O(n)（Python） |
| 反转单词顺序 | 按空白切分后反转 | O(n) | O(n) |

## 实现

- Go：[solution.go](go/solution.go) · [solution_test.go](go/solution_test.go)
- Python：[solution.py](python/solution.py) · [test_solution.py](python/test_solution.py)

## 关键点

- 滑动窗口的 `left` 只能右移，遇到重复字符时用 `max(left, lastSeen+1)` 避免回退；
- KMP 的 next 数组含义是「当前前缀的最长公共前后缀长度」，失配时跳转 `next[j-1]`；
- 中心扩展要分别处理奇数长度与偶数长度两种中心；
- 异位词计数窗口长度固定，进入新字符的同时移出旧字符；
- 中文等多字节字符需要先 `[]rune` 转换，本仓库示例按 ASCII 处理。

## 相关应用

- [05 文本搜索](../../applications/05-text-search/README.md)

## 相关主题

- [03 栈与队列](../03-stack-queue/README.md)
- [04 哈希表](../04-hash-table/README.md)
