package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
14. Longest Common Prefix -> https://leetcode.com/problems/longest-common-prefix/description/
Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".

Example 1:
Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
*/

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	small := strs[0]
	large := strs[len(strs)-1]
	var op strings.Builder
	for i := 0; i < len(small); i++ {
		if string(small[i]) == string(large[i]) {
			op.WriteByte(small[i])
		} else {
			return op.String()
		}
	}
	return op.String()
}
