// Package solution 提供字符串经典问题的 Go 实现。
package solution

import "strings"

// LengthOfLongestSubstring 返回不含重复字符的最长子串长度。
// 使用滑动窗口记录每个字符最近一次出现的位置，时间复杂度 O(n)。
func LengthOfLongestSubstring(text string) int {
	lastSeen := make(map[byte]int, len(text))
	left, longest := 0, 0
	for right := 0; right < len(text); right++ {
		if index, ok := lastSeen[text[right]]; ok && index >= left {
			left = index + 1
		}
		lastSeen[text[right]] = right
		longest = max(longest, right-left+1)
	}
	return longest
}

// FindAnagrams 返回 s 中所有 p 的异位词起始下标，按升序排列。
// 使用固定长度滑动窗口 + 字符计数，时间复杂度 O(n + m)。
func FindAnagrams(text, pattern string) []int {
	result := make([]int, 0)
	if len(pattern) == 0 || len(pattern) > len(text) {
		return result
	}

	// 小写字母计数表，下标由字符偏移得到
	const alphabetSize = 26
	window := make([]int, alphabetSize)
	target := make([]int, alphabetSize)
	for i := 0; i < len(pattern); i++ {
		target[pattern[i]-'a']++
	}

	for right := 0; right < len(text); right++ {
		window[text[right]-'a']++
		if right >= len(pattern) {
			window[text[right-len(pattern)]-'a']--
		}
		if right >= len(pattern)-1 && equalCounts(window, target) {
			result = append(result, right-len(pattern)+1)
		}
	}
	return result
}

// equalCounts 判断两个计数表是否完全一致。
func equalCounts(left, right []int) bool {
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}

// KMPSearch 返回 pattern 在 text 中所有匹配的起始下标，时间复杂度 O(n + m)。
func KMPSearch(text, pattern string) []int {
	result := make([]int, 0)
	if pattern == "" {
		return result
	}

	next := buildLPS(pattern)
	matched := 0
	for i := 0; i < len(text); i++ {
		for matched > 0 && text[i] != pattern[matched] {
			matched = next[matched-1]
		}
		if text[i] == pattern[matched] {
			matched++
		}
		if matched == len(pattern) {
			result = append(result, i-len(pattern)+1)
			matched = next[matched-1]
		}
	}
	return result
}

// buildLPS 构造 KMP 的最长公共前后缀数组。
func buildLPS(pattern string) []int {
	next := make([]int, len(pattern))
	length := 0
	for i := 1; i < len(pattern); i++ {
		for length > 0 && pattern[i] != pattern[length] {
			length = next[length-1]
		}
		if pattern[i] == pattern[length] {
			length++
			next[i] = length
		}
	}
	return next
}

// LongestPalindrome 返回最长回文子串，中心扩展法，时间复杂度 O(n^2)。
func LongestPalindrome(text string) string {
	if len(text) < 2 {
		return text
	}

	start, end := 0, 0
	for center := 0; center < len(text); center++ {
		oddStart, oddEnd := expandAroundCenter(text, center, center)
		evenStart, evenEnd := expandAroundCenter(text, center, center+1)
		if oddEnd-oddStart > end-start {
			start, end = oddStart, oddEnd
		}
		if evenEnd-evenStart > end-start {
			start, end = evenStart, evenEnd
		}
	}
	return text[start : end+1]
}

// expandAroundCenter 从中心向两侧扩展，返回回文区间的闭区间下标。
func expandAroundCenter(text string, left, right int) (int, int) {
	for left >= 0 && right < len(text) && text[left] == text[right] {
		left--
		right++
	}
	return left + 1, right - 1
}

// IsPalindrome 判断字符串是否为回文，忽略非字母数字字符与大小写。
func IsPalindrome(text string) bool {
	left, right := 0, len(text)-1
	for left < right {
		if !isAlphanumeric(text[left]) {
			left++
			continue
		}
		if !isAlphanumeric(text[right]) {
			right--
			continue
		}
		if toLower(text[left]) != toLower(text[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

// isAlphanumeric 判断字符是否为数字或英文字母。
func isAlphanumeric(char byte) bool {
	return (char >= '0' && char <= '9') ||
		(char >= 'a' && char <= 'z') ||
		(char >= 'A' && char <= 'Z')
}

// toLower 将大写英文字母转为小写。
func toLower(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		return char + ('a' - 'A')
	}
	return char
}

// ReverseWords 反转字符串中的单词顺序，并合并多余空格。
func ReverseWords(text string) string {
	words := strings.Fields(text)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}
