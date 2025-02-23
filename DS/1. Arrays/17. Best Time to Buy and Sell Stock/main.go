package main

import (
	"fmt"
	"math"
)

/*
leetCode: 121 https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
You are given an array prices where prices[i] is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to
sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
Example 1:
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
Example 2:
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.


Approach: 2 pointer -> l,r
r will traverse -> object of l should to hv a minimium value in comparison with r.
*/
func main() {
	arr := []int{7, 1, 5, 3, 6, 4}
	op1 := maxProfit(arr)
	op2 := maxProfitDPSolution(arr)
	fmt.Println(op1)
	fmt.Println(op2)
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	leftBuy := 0
	rightSell := 1
	maxProfit := 0
	for rightSell < len(prices) {
		if prices[leftBuy] < prices[rightSell] {
			// fmt.Println(prices[rightSell] - prices[leftBuy])
			maxProfit = int(math.Max(float64(maxProfit), float64(prices[rightSell]-prices[leftBuy])))
		} else {
			leftBuy = rightSell
		}
		rightSell++
	}
	return maxProfit

}
func maxProfitDPSolution(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0
	for _, price := range prices {
		minPrice = int(math.Min(float64(price), float64(minPrice)))
		maxProfit = int(math.Max(float64(maxProfit), float64(price-minPrice)))
	}
	return maxProfit
}
