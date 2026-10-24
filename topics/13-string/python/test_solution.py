"""13 字符串的 pytest 测试。"""

import pytest
from solution_loader import load_solution

solution = load_solution(__file__)


@pytest.mark.parametrize(
    ("text", "expected"),
    [("abcabcbb", 3), ("bbbbb", 1), ("pwwkew", 3), ("", 0), ("abba", 2)],
)
def test_length_of_longest_substring(text: str, expected: int) -> None:
    assert solution.length_of_longest_substring(text) == expected


def test_find_anagrams() -> None:
    assert solution.find_anagrams("cbaebabacd", "abc") == [0, 6]
    assert solution.find_anagrams("ab", "abc") == []


@pytest.mark.parametrize(
    ("text", "pattern", "expected"),
    [
        ("aaaaa", "aa", [0, 1, 2, 3]),
        ("abababab", "abab", [0, 2, 4]),
        ("hello", "world", []),
        ("abc", "", []),
    ],
)
def test_kmp_search(text: str, pattern: str, expected: list[int]) -> None:
    assert solution.kmp_search(text, pattern) == expected


@pytest.mark.parametrize(
    ("text", "expected"),
    [("babad", "bab"), ("cbbd", "bb"), ("a", "a"), ("", "")],
)
def test_longest_palindrome(text: str, expected: str) -> None:
    assert solution.longest_palindrome(text) == expected


@pytest.mark.parametrize(
    ("text", "expected"),
    [
        ("A man, a plan, a canal: Panama", True),
        ("race a car", False),
        (" ", True),
        ("0P", False),
    ],
)
def test_is_palindrome(text: str, expected: bool) -> None:
    assert solution.is_palindrome(text) is expected


def test_reverse_words() -> None:
    assert solution.reverse_words("  the sky   is blue  ") == "blue is sky the"
