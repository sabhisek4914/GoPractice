package main

import (
	"fmt"
	"sort"
	"strings"
)

// Suhani Rastogi says:Group anagrams : Given an array of strings, group anagrams together. Anagrams: Two strings are anagrams if both have same characters and count of each character to be same. In other words, they are shuffled version of one another. example – eat, tea, ate are anagrams Example:
// Input: ["eat", "tea", "tan", "ate", "nat", "bat"]
// Output: [ ["ate", "eat","tea"], ["nat","tan"], ["bat"] ]

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	op_map := make(map[string][]string)
	for i := 0; i < len(input); i++ {
		c_str := input[i]

		//sorting
		chars := strings.Split(c_str, "")
		sort.Strings(chars)
		sorted_str := strings.Join(chars, "")
		// fmt.Println(sorted_str)

		//insert to map
		arr := op_map[sorted_str]
		arr = append(arr, input[i])
		op_map[sorted_str] = arr
	}
	fmt.Println(op_map)

	//convert map to 2D array
	op := make([][]string, 0)
	for _, val := range op_map {
		op = append(op, val)
	}
	fmt.Println(op)
}
