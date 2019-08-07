/*
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like
(ie, buy one and sell one share of the stock multiple times) with the following restrictions:

You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1 day)
Example:

Input: [1,2,3,0,2]
Output: 3
Explanation: transactions = [buy, sell, cooldown, buy, sell]
 */

package main

import "fmt"

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

/*
Input: [1,2,3,0,2]
Output: 3

 */

func maxProfit(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}
	buy := -prices[0]
	buy1 := buy
	sell := 0
	sell1 := 0
	sell2 := 0

	for i := 1; i < len(prices); i++ {
		buy = max(buy1, sell2 - prices[i])
		sell = max(sell1, buy1 + prices[i])
		buy1 = buy
		sell2 = sell1
		sell1 = sell
	}
	return sell
}

func main() {
	fmt.Println(maxProfit([]int{1,2,3,0,2}))
}