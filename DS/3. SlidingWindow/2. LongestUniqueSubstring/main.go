package main

import "fmt"

/*
3.  Longest Substring Without Repeating Characters -> https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
Given a string s, find the length of the longest substring without duplicate characters.
Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Time Complexity Analysis
Best Case (All Unique Characters)
Each character is processed once by j, moving from 0 to len(s).
Since s[j] is never duplicated, the inner loop never runs.
The complexity is O(n).
Worst Case (All Characters Repeated)
Every character might be removed from uniqueMap when a duplicate is found.
Each character is inserted and removed at most once.
Each pointer (i and j) moves forward at most n times.
The inner loop runs at most n times in total, not n times per iteration.
The total operations are O(2n) = O(n).
*/

func main() {
	str := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(s string) int {
	uniqueMap := make(map[string]bool)
	i, j := 0, 0
	large := 0
	for j < len(s) {
		for uniqueMap[string(s[j])] {
			delete(uniqueMap, string(s[i]))
			i++
		}
		uniqueMap[string(s[j])] = true
		if len(uniqueMap) > large {
			large = len(uniqueMap)
		}
		j++
	}
	return large
}
