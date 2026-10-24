package solution

import (
	"reflect"
	"testing"
)

func TestInvertedIndex(t *testing.T) {
	index := NewInvertedIndex()
	index.Add(1, "The quick brown fox")
	index.Add(2, "quick brown dog")

	if got := index.Search("quick"); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Fatalf("Search(quick) = %v, want [1 2]", got)
	}
	if got := index.Search("QUICK"); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Fatalf("查询应不区分大小写, got %v", got)
	}
	if got := index.Search("fox"); !reflect.DeepEqual(got, []int{1}) {
		t.Fatalf("Search(fox) = %v, want [1]", got)
	}
	if got := index.Search("cat"); len(got) != 0 {
		t.Fatalf("不存在的单词应返回空, got %v", got)
	}
}

func TestInvertedIndexReAdd(t *testing.T) {
	index := NewInvertedIndex()
	index.Add(1, "old words")
	index.Add(1, "new words")

	if got := index.Search("old"); len(got) != 0 {
		t.Fatalf("重新索引后旧词应被移除, got %v", got)
	}
	if got := index.Search("new"); !reflect.DeepEqual(got, []int{1}) {
		t.Fatalf("Search(new) = %v, want [1]", got)
	}
}

func TestTrie(t *testing.T) {
	trie := NewTrie()
	for _, word := range []string{"apple", "app", "apply", "banana"} {
		trie.Insert(word)
	}

	if !trie.Search("app") {
		t.Fatal("app 应存在")
	}
	if trie.Search("ap") {
		t.Fatal("ap 不是完整单词")
	}
	if !trie.Search("apply") {
		t.Fatal("apply 应存在")
	}
}

func TestTrieComplete(t *testing.T) {
	trie := NewTrie()
	for _, word := range []string{"apple", "app", "apply", "banana"} {
		trie.Insert(word)
	}

	got := trie.Complete("app", 0)
	want := []string{"app", "apple", "apply"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Complete(app) = %v, want %v", got, want)
	}

	got = trie.Complete("app", 2)
	want = []string{"app", "apple"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("带 limit 的 Complete = %v, want %v", got, want)
	}

	if got := trie.Complete("zzz", 0); len(got) != 0 {
		t.Fatalf("不存在的前缀应返回空, got %v", got)
	}
}

func TestTokenize(t *testing.T) {
	got := tokenize("Hello, world! 2026")
	want := []string{"hello", "world", "2026"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("tokenize = %v, want %v", got, want)
	}
}
