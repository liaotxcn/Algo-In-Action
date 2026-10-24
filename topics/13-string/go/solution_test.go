package solution

import (
	"reflect"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		text string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"abba", 2},
	}
	for _, tc := range cases {
		if got := LengthOfLongestSubstring(tc.text); got != tc.want {
			t.Fatalf("LengthOfLongestSubstring(%q) = %d, want %d", tc.text, got, tc.want)
		}
	}
}

func TestFindAnagrams(t *testing.T) {
	got := FindAnagrams("cbaebabacd", "abc")
	want := []int{0, 6}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("FindAnagrams = %v, want %v", got, want)
	}
	if got := FindAnagrams("ab", "abc"); len(got) != 0 {
		t.Fatalf("模式串更长时应返回空, got %v", got)
	}
}

func TestKMPSearch(t *testing.T) {
	cases := []struct {
		text    string
		pattern string
		want    []int
	}{
		{"aaaaa", "aa", []int{0, 1, 2, 3}},
		{"abababab", "abab", []int{0, 2, 4}},
		{"hello", "world", []int{}},
		{"abc", "", []int{}},
	}
	for _, tc := range cases {
		got := KMPSearch(tc.text, tc.pattern)
		if len(got) != len(tc.want) || (len(got) > 0 && !reflect.DeepEqual(got, tc.want)) {
			t.Fatalf("KMPSearch(%q, %q) = %v, want %v", tc.text, tc.pattern, got, tc.want)
		}
	}
}

func TestLongestPalindrome(t *testing.T) {
	cases := []struct {
		text string
		want string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"", ""},
	}
	for _, tc := range cases {
		if got := LongestPalindrome(tc.text); got != tc.want {
			t.Fatalf("LongestPalindrome(%q) = %q, want %q", tc.text, got, tc.want)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		text string
		want bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{"0P", false},
	}
	for _, tc := range cases {
		if got := IsPalindrome(tc.text); got != tc.want {
			t.Fatalf("IsPalindrome(%q) = %v, want %v", tc.text, got, tc.want)
		}
	}
}

func TestReverseWords(t *testing.T) {
	got := ReverseWords("  the sky   is blue  ")
	want := "blue is sky the"
	if got != want {
		t.Fatalf("ReverseWords = %q, want %q", got, want)
	}
}
