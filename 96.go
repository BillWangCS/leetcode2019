/*
Given n, how many structurally unique BST's (binary search trees) that store values 1 ... n?

Example:

Input: 3
Output: 5
Explanation:
Given n = 3, there are a total of 5 unique BST's:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
 */

package main

import "fmt"

func numTrees(n int) int {
	dp := make([]int, 0)
	dp = append(dp, 1)
	dp = append(dp, 1)
	for i := 2; i <= n; i++ {
		s := 0
		for j := 1; j <= i; j++ {
			s += dp[i-j]*(dp[j-1])
		}
		dp = append(dp, s)
	}
	return dp[n]
}

func main() {
	fmt.Println(numTrees(6))
}
