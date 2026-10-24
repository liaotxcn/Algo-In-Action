// Package solution 提供堆与优先队列相关算法的 Go 实现。
package solution

import "slices"

// MaxHeap 是大顶堆。
type MaxHeap struct {
	data []int
}

// NewMaxHeap 创建空的大顶堆。
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

// Len 返回堆中元素个数。
func (h *MaxHeap) Len() int {
	return len(h.data)
}

// Peek 返回堆顶元素。
func (h *MaxHeap) Peek() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	return h.data[0], true
}

// Push 插入元素，时间复杂度 O(log n)。
func (h *MaxHeap) Push(value int) {
	h.data = append(h.data, value)
	h.siftUp(len(h.data) - 1)
}

// Pop 弹出堆顶元素，时间复杂度 O(log n)。
func (h *MaxHeap) Pop() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}

	top := h.data[0]
	last := len(h.data) - 1
	h.data[0] = h.data[last]
	h.data = h.data[:last]
	if len(h.data) > 0 {
		h.siftDown(0)
	}
	return top, true
}

// siftUp 从 index 向上调整，恢复大顶堆性质。
func (h *MaxHeap) siftUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.data[parent] >= h.data[index] {
			return
		}
		h.data[parent], h.data[index] = h.data[index], h.data[parent]
		index = parent
	}
}

// siftDown 从 index 向下调整，恢复大顶堆性质。
func (h *MaxHeap) siftDown(index int) {
	n := len(h.data)
	for {
		largest := index
		left, right := 2*index+1, 2*index+2
		if left < n && h.data[left] > h.data[largest] {
			largest = left
		}
		if right < n && h.data[right] > h.data[largest] {
			largest = right
		}
		if largest == index {
			return
		}
		h.data[index], h.data[largest] = h.data[largest], h.data[index]
		index = largest
	}
}

// HeapSort 原地升序排序，时间复杂度 O(n log n)，空间复杂度 O(1)。
func HeapSort(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}

	// 建堆：从最后一个非叶子节点开始向下调整
	for i := n/2 - 1; i >= 0; i-- {
		siftDownRange(nums, i, n)
	}

	// 依次把堆顶最大值交换到末尾
	for end := n - 1; end > 0; end-- {
		nums[0], nums[end] = nums[end], nums[0]
		siftDownRange(nums, 0, end)
	}
}

// siftDownRange 在 [0, size) 范围内向下调整大顶堆。
func siftDownRange(nums []int, index, size int) {
	for {
		largest := index
		left, right := 2*index+1, 2*index+2
		if left < size && nums[left] > nums[largest] {
			largest = left
		}
		if right < size && nums[right] > nums[largest] {
			largest = right
		}
		if largest == index {
			return
		}
		nums[index], nums[largest] = nums[largest], nums[index]
		index = largest
	}
}

// TopK 返回数组中最大的 k 个元素，按从大到小排列。
// 使用大小为 k 的小顶堆，时间复杂度 O(n log k)，空间复杂度 O(k)。
func TopK(nums []int, k int) []int {
	if k <= 0 || len(nums) == 0 {
		return []int{}
	}
	if k > len(nums) {
		k = len(nums)
	}

	heap := make([]int, 0, k)
	for _, value := range nums {
		if len(heap) < k {
			heap = append(heap, value)
			minSiftUp(heap, len(heap)-1)
			continue
		}
		if value > heap[0] {
			heap[0] = value
			minSiftDown(heap, 0)
		}
	}

	ans := slices.Clone(heap)
	slices.SortFunc(ans, func(a, b int) int { return b - a })
	return ans
}

// minSiftUp 在小顶堆中向上调整。
func minSiftUp(heap []int, index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if heap[parent] <= heap[index] {
			return
		}
		heap[parent], heap[index] = heap[index], heap[parent]
		index = parent
	}
}

// minSiftDown 在小顶堆中向下调整。
func minSiftDown(heap []int, index int) {
	n := len(heap)
	for {
		smallest := index
		left, right := 2*index+1, 2*index+2
		if left < n && heap[left] < heap[smallest] {
			smallest = left
		}
		if right < n && heap[right] < heap[smallest] {
			smallest = right
		}
		if smallest == index {
			return
		}
		heap[index], heap[smallest] = heap[smallest], heap[index]
		index = smallest
	}
}
