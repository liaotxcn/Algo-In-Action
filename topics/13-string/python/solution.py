"""13 字符串 —— Python 实现。"""

from __future__ import annotations


def length_of_longest_substring(text: str) -> int:
    """不含重复字符的最长子串长度。"""
    last_seen: dict[str, int] = {}
    left = longest = 0
    for right, char in enumerate(text):
        if char in last_seen and last_seen[char] >= left:
            left = last_seen[char] + 1
        last_seen[char] = right
        longest = max(longest, right - left + 1)
    return longest


def find_anagrams(text: str, pattern: str) -> list[int]:
    """返回 text 中所有 pattern 异位词的起始下标。"""
    if not pattern or len(pattern) > len(text):
        return []

    target = [0] * 26
    window = [0] * 26
    for char in pattern:
        target[ord(char) - ord("a")] += 1

    result: list[int] = []
    for right, char in enumerate(text):
        window[ord(char) - ord("a")] += 1
        if right >= len(pattern):
            window[ord(text[right - len(pattern)]) - ord("a")] -= 1
        if right >= len(pattern) - 1 and window == target:
            result.append(right - len(pattern) + 1)
    return result


def build_lps(pattern: str) -> list[int]:
    """构造 KMP 的最长公共前后缀数组。"""
    lps = [0] * len(pattern)
    length = 0
    for index in range(1, len(pattern)):
        while length > 0 and pattern[index] != pattern[length]:
            length = lps[length - 1]
        if pattern[index] == pattern[length]:
            length += 1
            lps[index] = length
    return lps


def kmp_search(text: str, pattern: str) -> list[int]:
    """KMP 算法返回所有匹配的起始下标。"""
    if not pattern:
        return []

    lps = build_lps(pattern)
    result: list[int] = []
    matched = 0
    for index, char in enumerate(text):
        while matched > 0 and char != pattern[matched]:
            matched = lps[matched - 1]
        if char == pattern[matched]:
            matched += 1
        if matched == len(pattern):
            result.append(index - len(pattern) + 1)
            matched = lps[matched - 1]
    return result


def expand_around_center(text: str, left: int, right: int) -> tuple[int, int]:
    """从中心向两侧扩展，返回回文闭区间。"""
    while left >= 0 and right < len(text) and text[left] == text[right]:
        left -= 1
        right += 1
    return left + 1, right - 1


def longest_palindrome(text: str) -> str:
    """中心扩展法求最长回文子串。"""
    if len(text) < 2:
        return text

    start = end = 0
    for center in range(len(text)):
        odd_start, odd_end = expand_around_center(text, center, center)
        even_start, even_end = expand_around_center(text, center, center + 1)
        if odd_end - odd_start > end - start:
            start, end = odd_start, odd_end
        if even_end - even_start > end - start:
            start, end = even_start, even_end
    return text[start : end + 1]


def is_palindrome(text: str) -> bool:
    """忽略非字母数字字符与大小写判断回文。"""
    filtered = [char.lower() for char in text if char.isalnum()]
    return filtered == filtered[::-1]


def reverse_words(text: str) -> str:
    """反转单词顺序，并合并多余空格。"""
    return " ".join(reversed(text.split()))
