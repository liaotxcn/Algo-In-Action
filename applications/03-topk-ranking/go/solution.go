// Package solution 提供 TopK 排行的 Go 实现。
package solution

import (
	"container/heap"
	"slices"
)

// intMinHeap 是求最大 K 个数时使用的小顶堆。
type intMinHeap []int

func (h intMinHeap) Len() int           { return len(h) }
func (h intMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intMinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *intMinHeap) Pop() any {
	old := *h
	last := len(old) - 1
	value := old[last]
	*h = old[:last]
	return value
}

// TopK 返回数组中最大的 k 个数，按降序排列。
// 使用大小为 k 的小顶堆，时间复杂度 O(n log k)，空间复杂度 O(k)。
func TopK(nums []int, k int) []int {
	if k <= 0 || len(nums) == 0 {
		return []int{}
	}
	if k > len(nums) {
		k = len(nums)
	}

	top := make(intMinHeap, 0, k)
	for _, value := range nums {
		if len(top) < k {
			heap.Push(&top, value)
			continue
		}
		if value > top[0] {
			heap.Pop(&top)
			heap.Push(&top, value)
		}
	}

	result := make([]int, len(top))
	copy(result, top)
	slices.SortFunc(result, func(a, b int) int { return b - a })
	return result
}

// wordCount 记录单词及其出现次数。
type wordCount struct {
	word  string
	count int
}

// wordMinHeap 把「更差」的词放在堆顶：次数少更差，次数相同则字典序大更差。
type wordMinHeap []wordCount

func (h wordMinHeap) Len() int { return len(h) }
func (h wordMinHeap) Less(i, j int) bool {
	if h[i].count != h[j].count {
		return h[i].count < h[j].count
	}
	return h[i].word > h[j].word
}
func (h wordMinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *wordMinHeap) Push(x any)   { *h = append(*h, x.(wordCount)) }
func (h *wordMinHeap) Pop() any {
	old := *h
	last := len(old) - 1
	value := old[last]
	*h = old[:last]
	return value
}

// TopKFrequent 返回出现次数最多的 k 个单词；次数相同时字典序小的优先。
func TopKFrequent(words []string, k int) []string {
	if k <= 0 || len(words) == 0 {
		return []string{}
	}

	counts := make(map[string]int, len(words))
	for _, word := range words {
		counts[word]++
	}

	top := make(wordMinHeap, 0, k)
	for word, count := range counts {
		candidate := wordCount{word: word, count: count}
		if len(top) < k {
			heap.Push(&top, candidate)
			continue
		}
		if betterWord(candidate, top[0]) {
			heap.Pop(&top)
			heap.Push(&top, candidate)
		}
	}

	result := make([]wordCount, len(top))
	copy(result, top)
	slices.SortFunc(result, func(a, b wordCount) int {
		if a.count != b.count {
			return b.count - a.count
		}
		switch {
		case a.word < b.word:
			return -1
		case a.word > b.word:
			return 1
		default:
			return 0
		}
	})

	ordered := make([]string, 0, len(result))
	for _, item := range result {
		ordered = append(ordered, item.word)
	}
	return ordered
}

// betterWord 判断 candidate 是否优于 current。
func betterWord(candidate, current wordCount) bool {
	if candidate.count != current.count {
		return candidate.count > current.count
	}
	return candidate.word < current.word
}
