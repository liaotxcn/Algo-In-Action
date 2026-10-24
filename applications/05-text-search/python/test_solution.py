"""05 文本搜索的 pytest 测试。"""

from solution_loader import load_solution

solution = load_solution(__file__)


def test_inverted_index() -> None:
    index = solution.InvertedIndex()
    index.add(1, "The quick brown fox")
    index.add(2, "quick brown dog")

    assert index.search("quick") == [1, 2]
    assert index.search("QUICK") == [1, 2]
    assert index.search("fox") == [1]
    assert index.search("cat") == []


def test_inverted_index_re_add() -> None:
    index = solution.InvertedIndex()
    index.add(1, "old words")
    index.add(1, "new words")

    assert index.search("old") == []
    assert index.search("new") == [1]


def test_trie() -> None:
    trie = solution.Trie()
    for word in ["apple", "app", "apply", "banana"]:
        trie.insert(word)

    assert trie.search("app")
    assert not trie.search("ap")
    assert trie.search("apply")


def test_trie_complete() -> None:
    trie = solution.Trie()
    for word in ["apple", "app", "apply", "banana"]:
        trie.insert(word)

    assert trie.complete("app") == ["app", "apple", "apply"]
    assert trie.complete("app", 2) == ["app", "apple"]
    assert trie.complete("zzz") == []


def test_tokenize() -> None:
    assert solution.tokenize("Hello, world! 2026") == ["hello", "world", "2026"]
