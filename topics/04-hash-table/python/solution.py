"""04 哈希表 —— Python 实现。"""

from __future__ import annotations

from collections import defaultdict


def two_sum(nums: list[int], target: int) -> list[int] | None:
    """返回和为目标值的两个下标，时间复杂度 O(n)。"""
    seen: dict[int, int] = {}
    for index, value in enumerate(nums):
        other = seen.get(target - value)
        if other is not None:
            return [other, index]
        seen[value] = index
    return None


def longest_consecutive(nums: list[int]) -> int:
    """最长连续序列，时间复杂度 O(n)。"""
    values = set(nums)
    ans = 0
    for value in values:
        if value - 1 in values:
            continue
        length = 1
        while value + length in values:
            length += 1
        ans = max(ans, length)
    return ans


def group_anagrams(strs: list[str]) -> list[list[str]]:
    """字母异位词分组，输出按字典序稳定排列。"""
    groups: dict[str, list[str]] = defaultdict(list)
    for word in strs:
        groups["".join(sorted(word))].append(word)
    result = [sorted(group) for group in groups.values()]
    result.sort(key=lambda group: group[0])
    return result


def subarray_sum(nums: list[int], k: int) -> int:
    """和为 k 的连续子数组个数，时间复杂度 O(n)。"""
    prefix_count: dict[int, int] = {0: 1}
    prefix = ans = 0
    for value in nums:
        prefix += value
        ans += prefix_count.get(prefix - k, 0)
        prefix_count[prefix] = prefix_count.get(prefix, 0) + 1
    return ans


def is_anagram(a: str, b: str) -> bool:
    """判断两个字符串是否互为字母异位词。"""
    if len(a) != len(b):
        return False
    counts: dict[str, int] = {}
    for left, right in zip(a, b):
        counts[left] = counts.get(left, 0) + 1
        counts[right] = counts.get(right, 0) - 1
    return all(count == 0 for count in counts.values())


def normalize_word(word: str) -> str:
    """返回按字母排序后的字符串。"""
    return "".join(sorted(word))
