package main

import "fmt"

/*
392. https://leetcode.com/problems/is-subsequence/description/
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the
characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde"
while "aec" is not).

Example 1:
Input: s = "abc", t = "ahbgdc"
Output: true
*/

func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	sLen, tLen := len(s)-1, len(t)-1
	for tLen >= 0 {
		if s[sLen] == t[tLen] {
			sLen--
			tLen--
		} else {
			tLen--
		}
		if sLen == -1 {
			return true
		}
	}
	return false
}
