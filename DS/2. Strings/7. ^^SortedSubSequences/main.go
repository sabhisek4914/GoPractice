package main

import (
	"fmt"
	"sort"
)

/*
Sorted Subsequences in loxographical order
inpuT: "abcd"
op:= null,a,b,c,d,ab,ac,ad,bc,bd,cd,abc,abd,acd,bcd,abcd
-> See copy
*/

func main() {
	str := "abcd"
	var subsequences []string

	generateSubsequences(str, "", &subsequences) // Pass slice as a pointer

	sort.Slice(subsequences, func(i, j int) bool {
		if len(subsequences[i]) == len(subsequences[j]) {
			return subsequences[i] < subsequences[j]
		}
		return len(subsequences[i]) < len(subsequences[j])
	})
	fmt.Println("All Subsequences:", subsequences)
}

// Recursive function to generate subsequences
func generateSubsequences(str, current string, subsequences *[]string) {
	// Base case: When the string is empty, store the accumulated subsequence
	if len(str) == 0 {
		*subsequences = append(*subsequences, current)
		return
	}

	// Take the first character
	ch := string(str[0])
	remain := str[1:]

	// Recursive calls:
	generateSubsequences(remain, current+ch, subsequences) // Include the character
	generateSubsequences(remain, current, subsequences)    // Exclude the character
}
