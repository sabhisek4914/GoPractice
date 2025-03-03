package main

import (
	"fmt"
	"strings"
)

/*
1328. Break a Pallindrome -> https://leetcode.com/problems/break-a-palindrome/description/
Given a palindromic string of lowercase English letters palindrome, replace exactly one character
with any lowercase English letter so that the resulting string is not a palindrome and that it is the
lexicographically smallest one possible.
Return the resulting string. If there is no way to replace a character to make it not a palindrome, return an empty string.
A string a is lexicographically smaller than a string b (of the same length) if in the first position where a and b differ,
a has a character strictly smaller than the corresponding character in b. For example, "abcc" is lexicographically smaller
than "abcd" because the first position they differ is at the fourth character, and 'c' is smaller than 'd'.

Example 1:
Input: palindrome = "abccba"
Output: "aaccba"
Explanation: There are many ways to make "abccba" not a palindrome, such as "zbccba", "aaccba", and "abacba".
Of all the ways, "aaccba" is the lexicographically smallest.
Example 2:
Input: palindrome = "a"
Output: ""
Explanation: There is no way to replace a single character to make "a" not a palindrome, so return an empty string.
Example 3:
Input: palindrome = "aaa"
Output: "aab"

Approach:
1. We need a lexographically smallest string.
2. 3 Scenarios
	1. a,b,c,b,a -> a,a,c,b,a
	2. a,a,a,a -> a,a,a,b
	3. a- > ""
*/

func main() {
	str := "aaaa"
	fmt.Println(BreakAPAllindrome(str))
	fmt.Println(BreakAPalindromeRune(str))
}

func BreakAPAllindrome(pallindrome string) string {
	str := strings.Split(pallindrome, "")
	if len(str) > 1 {
		i := 0
		j := len(str) - 1
		for i < j {
			if str[i] != "a" {
				str[i] = "a"
				return strings.Join(str, "")
			}
			i++
			j--
		}
		str[len(str)-1] = "b"
	} else {
		return ""
	}
	return strings.Join(str, "")
}

func BreakAPalindromeRune(palindrome string) string {
	if len(palindrome) <= 1 {
		return ""
	}

	// Convert to a mutable slice of runes
	chars := []rune(palindrome)

	// Try to replace the first non-'a' character with 'a'
	for i := 0; i < len(chars)/2; i++ {
		if chars[i] != 'a' {
			chars[i] = 'a'
			return string(chars)
		}
	}

	// If all characters are 'a', change the last character to 'b'
	chars[len(chars)-1] = 'b'
	return string(chars)
}
