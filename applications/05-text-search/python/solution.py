"""05 文本搜索 —— Python 实现。"""

from __future__ import annotations

import re
import threading

TOKEN_PATTERN = re.compile(r"[a-z0-9]+")


def tokenize(text: str) -> list[str]:
    """把文本切分为小写字母数字单词。"""
    return TOKEN_PATTERN.findall(text.lower())


class InvertedIndex:
    """线程安全的倒排索引。"""

    def __init__(self) -> None:
        self._documents: dict[int, list[str]] = {}
        self._index: dict[str, set[int]] = {}
        self._lock = threading.RLock()

    def add(self, doc_id: int, text: str) -> None:
        """为文档建立索引；重复添加同一文档会先移除旧词条。"""
        words = tokenize(text)
        with self._lock:
            self._remove(doc_id)
            unique = list(dict.fromkeys(words))
            for word in unique:
                self._index.setdefault(word, set()).add(doc_id)
            self._documents[doc_id] = unique

    def _remove(self, doc_id: int) -> None:
        """移除文档的旧索引，调用方必须持有锁。"""
        for word in self._documents.pop(doc_id, []):
            postings = self._index.get(word)
            if postings is None:
                continue
            postings.discard(doc_id)
            if not postings:
                del self._index[word]

    def search(self, word: str) -> list[int]:
        """返回包含单词的文档 ID，按升序排列。"""
        with self._lock:
            return sorted(self._index.get(word.lower(), set()))


class _TrieNode:
    """前缀树节点。"""

    __slots__ = ("children", "terminal")

    def __init__(self) -> None:
        self.children: dict[str, _TrieNode] = {}
        self.terminal = False

    def collect(self, prefix: str, limit: int, result: list[str]) -> None:
        """深度优先收集单词，子节点按字符顺序遍历。"""
        if limit > 0 and len(result) >= limit:
            return
        if self.terminal:
            result.append(prefix)
        for char in sorted(self.children):
            self.children[char].collect(prefix + char, limit, result)
            if limit > 0 and len(result) >= limit:
                return


class Trie:
    """用于前缀检索的字典树。"""

    def __init__(self) -> None:
        self._root = _TrieNode()

    def insert(self, word: str) -> None:
        """插入单词。"""
        node = self._root
        for char in word:
            node = node.children.setdefault(char, _TrieNode())
        node.terminal = True

    def _find(self, prefix: str) -> _TrieNode | None:
        """返回前缀对应的节点。"""
        node = self._root
        for char in prefix:
            child = node.children.get(char)
            if child is None:
                return None
            node = child
        return node

    def search(self, word: str) -> bool:
        """判断单词是否完整存在。"""
        node = self._find(word)
        return node is not None and node.terminal

    def complete(self, prefix: str, limit: int = 0) -> list[str]:
        """返回以 prefix 开头的单词；limit <= 0 表示不限制数量。"""
        node = self._find(prefix)
        if node is None:
            return []
        result: list[str] = []
        node.collect(prefix, limit, result)
        return result
