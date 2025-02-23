package main

import "fmt"

/*
2149. Rearrange Array Elements by Sign https://leetcode.com/problems/rearrange-array-elements-by-sign/description/
You are given a 0-indexed integer array nums of even length consisting of an equal number of positive and negative integers.
You should return the array of nums such that the the array follows the given conditions:
Every consecutive pair of integers have opposite signs.
For all integers with the same sign, the order in which they were present in nums is preserved.
The rearranged array begins with a positive integer.
Return the modified array after rearranging the elements to satisfy the aforementioned conditions.
Example 1:
Input: nums = [3,1,-2,-5,2,-4]
Output: [3,-2,1,-5,2,-4]
Explanation:
The positive integers in nums are [3,1,2]. The negative integers are [-2,-5,-4].
The only possible way to rearrange them such that they satisfy all conditions is [3,-2,1,-5,2,-4].
Other ways such as [1,-2,2,-5,3,-4], [3,1,2,-2,-5,-4], [-2,3,-5,1,-4,2] are incorrect because they do not satisfy one or more conditions.
Example 2:
Input: nums = [-1,1]
Output: [1,-1]
Explanation:
1 is the only positive integer and -1 the only negative integer in nums.
So nums is rearranged to [1,-1].

Here time Complexity is: O(n) -> Analyze how?
for i := 0; i < n; i += 2 -> O(n/2)
	for every value of i
		for pos < n && nums[pos] < 0
		for neg < n && nums[neg] >= 0
			pos & neg movement is at mostonce bcz count of +ve == -ve
*/

func main() {
	nums := []int{3, 1, -2, -5, 2, -4}
	x := rearrangeArray(nums)
	fmt.Println(x)

}
func rearrangeArray(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	pos, neg := 0, 0
	for i := 0; i < n; i += 2 {
		for pos < n && nums[pos] < 0 {
			pos++
		}
		for neg < n && nums[neg] >= 0 {
			neg++
		}
		ans[i], ans[i+1] = nums[pos], nums[neg]
		pos++
		neg++
	}
	return ans

	// positiveArr:=make([]int,0)
	// negativeArr:=make([]int,0)
	// opArr:=make([]int,0)
	// for _,num:=range nums{
	//     if num>=0{
	//         positiveArr=append(positiveArr,num)
	//     }else{
	//         negativeArr=append(negativeArr,num)
	//     }
	// }
	// for i:=0;i<len(nums)/2;i++{
	//     opArr=append(opArr,positiveArr[i])
	//     opArr=append(opArr,negativeArr[i])
	// }
	// return opArr
}
