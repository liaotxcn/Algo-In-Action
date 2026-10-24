// Package solution 提供动态规划经典问题的 Go 实现。
package solution

import "slices"

// Fib 返回第 n 个斐波那契数（F(0)=0, F(1)=1）。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func Fib(n int) int {
	if n <= 0 {
		return 0
	}
	prev, current := 0, 1
	for i := 2; i <= n; i++ {
		prev, current = current, prev+current
	}
	return current
}

// ClimbStairs 返回爬 n 级台阶的方法数（每次 1 或 2 级）。
func ClimbStairs(n int) int {
	if n <= 2 {
		return max(n, 0)
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// Knapsack01 求解 0-1 背包的最大价值。
// 时间复杂度 O(n·C)，空间复杂度 O(C)。
func Knapsack01(weights, values []int, capacity int) int {
	if capacity <= 0 || len(weights) == 0 {
		return 0
	}

	dp := make([]int, capacity+1)
	for i, weight := range weights {
		if weight > capacity {
			continue
		}
		// 逆序更新，保证每件物品最多使用一次
		for c := capacity; c >= weight; c-- {
			dp[c] = max(dp[c], dp[c-weight]+values[i])
		}
	}
	return dp[capacity]
}

// LIS 返回最长严格递增子序列的长度，时间复杂度 O(n log n)。
func LIS(nums []int) int {
	tails := make([]int, 0, len(nums))
	for _, value := range nums {
		// 找到第一个 >= value 的位置
		index, found := slices.BinarySearch(tails, value)
		if !found {
			if index == len(tails) {
				tails = append(tails, value)
			} else {
				tails[index] = value
			}
		}
	}
	return len(tails)
}

// EditDistance 返回将 word1 转换为 word2 的最少操作数。
// 时间复杂度 O(m·n)，空间复杂度 O(n)。
func EditDistance(word1, word2 string) int {
	if len(word1) < len(word2) {
		word1, word2 = word2, word1
	}

	previous := make([]int, len(word2)+1)
	current := make([]int, len(word2)+1)
	for j := range previous {
		previous[j] = j
	}

	for i := 1; i <= len(word1); i++ {
		current[0] = i
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				current[j] = previous[j-1]
			} else {
				current[j] = 1 + min(previous[j-1], previous[j], current[j-1])
			}
		}
		previous, current = current, previous
	}
	return previous[len(word2)]
}

// CoinChange 返回凑出 amount 所需的最少硬币数；无法凑出返回 -1。
// 时间复杂度 O(n·amount)，空间复杂度 O(amount)。
func CoinChange(coins []int, amount int) int {
	unreachable := mathMaxInt()

	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = unreachable
	}
	for _, coin := range coins {
		if coin <= 0 {
			continue
		}
		for value := coin; value <= amount; value++ {
			if dp[value-coin] != unreachable {
				dp[value] = min(dp[value], dp[value-coin]+1)
			}
		}
	}
	if dp[amount] == unreachable {
		return -1
	}
	return dp[amount]
}

// mathMaxInt 返回 int 的最大值，避免直接依赖 math 包中的浮点常量。
func mathMaxInt() int {
	return int(^uint(0) >> 1)
}
