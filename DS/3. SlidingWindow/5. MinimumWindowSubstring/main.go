package main

import (
	"fmt"
	"math"
)

/*
76. Minimum Window Substring -> https://leetcode.com/problems/minimum-window-substring/description/
Given two strings s and t of lengths m and n respectively, return the minimum window substring of s such that every
character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".
The testcases will be generated such that the answer is unique.
Example 1:
Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
Example 2:
Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.
Example 3:
Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.
*/

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}
func minWindow(s, t string) string {
	freqS := [256]int{}
	freqT := [256]int{}
	small := math.MaxInt32
	for _, v := range t {
		freqT[v]++
	}
	// fmt.Println(freqT)
	cnt := 0
	start := 0
	str := ""
	for i := 0; i < len(s); i++ {
		freqS[s[i]]++
		if freqT[s[i]] != 0 && freqS[s[i]] <= freqT[s[i]] {
			cnt++
		}
		// fmt.Println(cnt)
		if cnt == len(t) {
			// fmt.Println(start, freqS[s[start]], freqT[s[start]])
			for freqT[s[start]] == 0 || (freqS[s[start]] > freqT[s[start]]) {
				// fmt.Println("OOO", start, freqS[s[start]], freqT[s[start]])
				freqS[s[start]]--
				start++
			}
			length := i - start + 1
			if length < small {
				// fmt.Println(length, small, i, start)
				small = length
				str = s[start : i+1]
				// fmt.Println(str)
			}
		}
	}
	if small == math.MaxInt32 {
		return ""
	}
	return str
}
