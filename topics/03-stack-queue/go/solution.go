// Package solution 提供栈与队列相关算法的 Go 实现。
package solution

// IsValid 判断括号字符串是否有效。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func IsValid(s string) bool {
	stack := make([]rune, 0, len(s))
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// MinStack 是支持 O(1) 获取最小值的栈。
type MinStack struct {
	data []int
	mins []int
}

// NewMinStack 创建空的最小栈。
func NewMinStack() *MinStack {
	return &MinStack{}
}

// Push 入栈。
func (s *MinStack) Push(value int) {
	s.data = append(s.data, value)
	if len(s.mins) == 0 || value <= s.mins[len(s.mins)-1] {
		s.mins = append(s.mins, value)
	}
}

// Pop 出栈并返回栈顶元素；空栈返回 0。
func (s *MinStack) Pop() int {
	if len(s.data) == 0 {
		return 0
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	if value == s.mins[len(s.mins)-1] {
		s.mins = s.mins[:len(s.mins)-1]
	}
	return value
}

// Top 返回栈顶元素；空栈返回 0。
func (s *MinStack) Top() int {
	if len(s.data) == 0 {
		return 0
	}
	return s.data[len(s.data)-1]
}

// GetMin 返回栈中最小元素；空栈返回 0。
func (s *MinStack) GetMin() int {
	if len(s.mins) == 0 {
		return 0
	}
	return s.mins[len(s.mins)-1]
}

// MyQueue 使用两个栈实现 FIFO 队列。
type MyQueue struct {
	inStack  []int
	outStack []int
}

// NewMyQueue 创建空队列。
func NewMyQueue() *MyQueue {
	return &MyQueue{}
}

// Push 入队。
func (q *MyQueue) Push(value int) {
	q.inStack = append(q.inStack, value)
}

// Pop 出队并返回队首元素；空队列返回 0。
func (q *MyQueue) Pop() int {
	q.move()
	if len(q.outStack) == 0 {
		return 0
	}
	value := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return value
}

// Peek 返回队首元素；空队列返回 0。
func (q *MyQueue) Peek() int {
	q.move()
	if len(q.outStack) == 0 {
		return 0
	}
	return q.outStack[len(q.outStack)-1]
}

// Empty 判断队列是否为空。
func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}

// move 在需要出队时把入栈元素整体转移到出栈。
func (q *MyQueue) move() {
	if len(q.outStack) > 0 {
		return
	}
	for len(q.inStack) > 0 {
		last := len(q.inStack) - 1
		q.outStack = append(q.outStack, q.inStack[last])
		q.inStack = q.inStack[:last]
	}
}

// MaxSlidingWindow 返回每个大小为 k 的滑动窗口的最大值。
// 使用单调递减队列，时间复杂度 O(n)，空间复杂度 O(k)。
func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 {
		return []int{}
	}
	if k > len(nums) {
		k = len(nums)
	}

	// deque 存储下标，对应的值从队首到队尾单调递减
	deque := make([]int, 0, k)
	ans := make([]int, 0, len(nums)-k+1)

	for i, value := range nums {
		// 移除已经滑出窗口的下标
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		// 保持单调递减
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= value {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)

		if i >= k-1 {
			ans = append(ans, nums[deque[0]])
		}
	}
	return ans
}
