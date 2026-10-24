// Package solution 提供倒排索引与前缀树（Trie）的 Go 实现。
package solution

import (
	"slices"
	"strings"
	"sync"
	"unicode"
)

// tokenize 把文本切分为小写字母数字单词。
func tokenize(text string) []string {
	words := make([]string, 0)
	var builder strings.Builder
	for _, char := range text {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			builder.WriteRune(unicode.ToLower(char))
			continue
		}
		if builder.Len() > 0 {
			words = append(words, builder.String())
			builder.Reset()
		}
	}
	if builder.Len() > 0 {
		words = append(words, builder.String())
	}
	return words
}

// InvertedIndex 是线程安全的倒排索引：单词 -> 文档 ID 集合。
type InvertedIndex struct {
	mu        sync.RWMutex
	documents map[int][]string
	index     map[string]map[int]struct{}
}

// NewInvertedIndex 创建空倒排索引。
func NewInvertedIndex() *InvertedIndex {
	return &InvertedIndex{
		documents: make(map[int][]string),
		index:     make(map[string]map[int]struct{}),
	}
}

// Add 为文档建立索引；重复添加同一文档会先移除旧词条。
func (idx *InvertedIndex) Add(docID int, text string) {
	words := tokenize(text)

	idx.mu.Lock()
	defer idx.mu.Unlock()

	idx.removeLocked(docID)
	unique := make([]string, 0, len(words))
	seen := make(map[string]struct{}, len(words))
	for _, word := range words {
		if _, ok := seen[word]; ok {
			continue
		}
		seen[word] = struct{}{}
		unique = append(unique, word)
		if idx.index[word] == nil {
			idx.index[word] = make(map[int]struct{})
		}
		idx.index[word][docID] = struct{}{}
	}
	idx.documents[docID] = unique
}

// removeLocked 移除文档的旧索引，调用方必须持有写锁。
func (idx *InvertedIndex) removeLocked(docID int) {
	words, ok := idx.documents[docID]
	if !ok {
		return
	}
	for _, word := range words {
		delete(idx.index[word], docID)
		if len(idx.index[word]) == 0 {
			delete(idx.index, word)
		}
	}
	delete(idx.documents, docID)
}

// Search 返回包含单词的文档 ID，按升序排列；查询不区分大小写。
func (idx *InvertedIndex) Search(word string) []int {
	key := strings.ToLower(word)

	idx.mu.RLock()
	defer idx.mu.RUnlock()

	postings := idx.index[key]
	result := make([]int, 0, len(postings))
	for docID := range postings {
		result = append(result, docID)
	}
	slices.Sort(result)
	return result
}

// trieNode 是前缀树节点。
type trieNode struct {
	children map[rune]*trieNode
	terminal bool
}

// Trie 是用于前缀检索的字典树。
type Trie struct {
	root *trieNode
}

// NewTrie 创建空前缀树。
func NewTrie() *Trie {
	return &Trie{root: &trieNode{children: make(map[rune]*trieNode)}}
}

// Insert 插入单词。
func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		child, ok := node.children[char]
		if !ok {
			child = &trieNode{children: make(map[rune]*trieNode)}
			node.children[char] = child
		}
		node = child
	}
	node.terminal = true
}

// Search 判断单词是否完整存在。
func (t *Trie) Search(word string) bool {
	node := t.find(word)
	return node != nil && node.terminal
}

// find 返回前缀对应的节点，不存在时返回 nil。
func (t *Trie) find(prefix string) *trieNode {
	node := t.root
	for _, char := range prefix {
		child, ok := node.children[char]
		if !ok {
			return nil
		}
		node = child
	}
	return node
}

// Complete 返回以 prefix 开头的单词，按字典序排列；limit <= 0 表示不限制数量。
func (t *Trie) Complete(prefix string, limit int) []string {
	node := t.find(prefix)
	if node == nil {
		return []string{}
	}

	result := make([]string, 0)
	node.collect(prefix, limit, &result)
	return result
}

// collect 深度优先收集单词，子节点按字符顺序遍历以保证字典序。
func (n *trieNode) collect(prefix string, limit int, result *[]string) {
	if limit > 0 && len(*result) >= limit {
		return
	}
	if n.terminal {
		*result = append(*result, prefix)
	}

	keys := make([]rune, 0, len(n.children))
	for char := range n.children {
		keys = append(keys, char)
	}
	slices.Sort(keys)

	for _, char := range keys {
		n.children[char].collect(prefix+string(char), limit, result)
		if limit > 0 && len(*result) >= limit {
			return
		}
	}
}
