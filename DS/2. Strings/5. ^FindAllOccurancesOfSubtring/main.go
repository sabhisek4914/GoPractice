package main

import (
	"fmt"
	"regexp"
	"strings"
)

/*
Find all the occurances of a substring in a given string
inp: "abc abc abc pqr abc"
key: "abc"
OP: 4
*/

func main() {
	str := "abc abc abc pqr abc"
	key := "abc"
	cnt := FindSubstringIndex(str, key)
	fmt.Println(cnt)
	// fmt.Println(strings.Index(str[0:], key))
	// fmt.Println(strings.Index(str[4:], key))
	x := strings.Split(str, key)
	fmt.Println(x, len(x)-1) //O(n)
	cnt = findAllRegex(str, key)
	fmt.Println(cnt)
	cnt1 := countOccurrences(str, key)
	fmt.Println(cnt1)
}

//Time COmplexity is O(n*m) where n is length of string, m length of key
/*
Overall Time Complexity
Worst Case: O(n^2) if key is a single character and found frequently (strings.Index() runs O(n), and the loop runs O(n) times).
Best Case: O(n / m) if key has length m, reducing the number of iterations.
Average Case: O(n) for practical cases where m is significant.
✅ Final Complexity: O(n) to O(n^2), depending on key length and occurrences.

strings.Index(str[start:], key) -> O(n)
for loop O(n)
*/
func FindSubstringIndex(str, key string) []int {
	var indices []int
	start := 0
	for {
		x := strings.Index(str[start:], key)
		// fmt.Println(x, start)
		if x == -1 {
			break
		} else {
			indices = append(indices, start+x)
			start += x + len(key)
		}
	}
	// or
	// for {
	// 	x := strings.Index(str[start:], key)
	// 	fmt.Println(x, start)
	// 	if x == -1 {
	// 		break
	// 	} else {
	// 		indices = append(indices, start+x)
	// 		start += x + +
	// 	}
	// }
	return indices
}

//O(n) -> More efficient for large strings
func findAllRegex(text, pattern string) []int {
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringIndex(text, -1)
	fmt.Println(matches)
	var indices []int
	for _, match := range matches {
		indices = append(indices, match[0]) // Only store the start index
	}

	return indices
}

// O(n) -> Can return cnt
func countOccurrences(text, sub string) int {
	parts := strings.Split(text, sub)
	return len(parts) - 1
}
