package main

import "fmt"

/*
1903. Largest Odd Number in String ->https://leetcode.com/problems/largest-odd-number-in-string/description/
You are given a string num, representing a large integer. Return the largest-valued odd integer (as a string) that is a non-empty substring of num, or an empty string "" if no odd integer exists.
A substring is a contiguous sequence of characters within a string.

Example 1:
Input: num = "52"
Output: "5"
Explanation: The only non-empty substrings are "5", "2", and "52". "5" is the only odd number.
Example 2:
Input: num = "4206"
Output: ""
Explanation: There are no odd numbers in "4206".
Example 3:
Input: num = "35427"
Output: "35427"
Explanation: "35427" is already an odd number.
*/

func main() {
	fmt.Println(largestOddNumber("52"))
}

func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		if (num[i]-'0')%2 != 0 {
			return num[:i+1]
		}
	}

	return ""
}

// func largestOddNumber(num string) string {
//     large:=int64(math.MinInt64)
//     snumI,_:=strconv.ParseInt(string(num),10,64)
//     for snumI>0{
//         if snumI%2!=0 && snumI >large{
//             large=snumI
//         }
//         snumI/=10
//     }
//     if large == int64(math.MinInt64){
//         return ""
//     }
//     return strconv.FormatInt(large,10)
// }
