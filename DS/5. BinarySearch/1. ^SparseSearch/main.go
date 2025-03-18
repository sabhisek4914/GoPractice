package main

import "fmt"

/*
sorted []strings with empty spaces -> search an element in the array.
linear -> O(n)
binary -> O(logn)
*/
func main() {
	str := []string{"ai", "", "", "bat", "", "car", "cat", "", "", "dog", ""}
	k := "bat" //op: 4
	s, e := 0, len(str)-1
	for s <= e {
		mid := (s + e) / 2
		pos := -1
		if str[mid] == "" {
			i, j := mid-1, mid+1
			for i >= s || j <= e {
				if i < s || j > e {
					break
				}
				if str[i] != "" {
					pos = i
					break
				}
				if str[j] != "" {
					pos = j
					break
				}
				i--
				j++
			}
		} else {
			pos = mid
		}
		if pos == -1 {
			fmt.Println("ELE NOT FOUND")
		}
		if str[pos] == k {
			fmt.Println("Element Found at pos:", pos+1)
			break
		} else if str[pos] > k {
			e = pos
		} else {
			s = pos
		}
	}
}
