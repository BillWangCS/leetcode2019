/*
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

Example 1:

Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
 */

package main

import "fmt"

func climbing(n int, visited *[]int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if (*visited)[n] != -1 {
		return (*visited)[n]
	}
	(*visited)[n] = climbStairs(n-1) + climbStairs(n-2)
	return (*visited)[n]
}

func climbStairs(n int) int {
	visited := make([]int, n+1)
	for i := 0; i < len(visited); i++ {
		visited[i] = -1
	}
	return climbing(n, &visited)
}

func main() {
	fmt.Println(climbStairs(3))
}