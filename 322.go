/*
You are given coins of different denominations and a total amount of money amount.
Write a function to compute the fewest number of coins that you need to make up that amount.
If that amount of money cannot be made up by any combination of the coins, return -1.

Example 1:

Input: coins = [1, 2, 5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
Example 2:

Input: coins = [2], amount = 3
Output: -1
Note:
You may assume that you have an infinite number of each kind of coin.
 */

package main

import (
	"fmt"
)

//DFS time limited out
/*
func reverseArr(A *[]int) {
	i := 0
	j := len(*A) - 1
	for i <= j {
		(*A)[i], (*A)[j] = (*A)[j], (*A)[i]
		i++
		j--
	}
}

func dfs_322(max *int, amount int, coins []int, step int) {
	if *max < step {
		return
	}
	if amount == 0 {
		if step < *max {
			*max = step
		}
		return
	}
	if amount < 0 {
		return
	}
	for i := 0; i < len(coins); i++ {
		if amount >= coins[i] {
			dfs_322(max, amount-coins[i], coins, step+1)
		}
	}
}

func coinChange(coins []int, amount int) int {
	res := 1<<31
	sort.Ints(coins)
	reverseArr(&coins)
	fmt.Println(coins)
	dfs_322(&res, amount, coins, 0)
	if res == 1<<31 {
		return -1
	}
	return res
}
*/

//DP
func change(coins []int, amount int, f *[]int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return 1<<31
	}
	if (*f)[amount] != -1 {
		return (*f)[amount]
	}
	min := 1<<31
	for i := 0; i < len(coins); i++ {
		tmp := change(coins, amount-coins[i], f) + 1
		if tmp < min {
			min =tmp
		}
	}
	(*f)[amount] = min
	return (*f)[amount]
}

func coinChange(coins []int, amount int) int {
	f := make([]int, 0)
	for i := 0; i <= amount; i++ {
		f = append(f, -1)
	}
	res := change(coins, amount, &f)
	if res == 1<<31 {
		return -1
	}
	return res
}

func main() {
	fmt.Println(coinChange([]int{1,2,5}, 100))
}