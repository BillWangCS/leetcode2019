/*
On a staircase, the i-th step has some non-negative cost cost[i] assigned (0 indexed).

Once you pay the cost, you can either climb one or two steps.
You need to find minimum cost to reach the top of the floor,
and you can either start from the step with index 0, or the step with index 1.

Example 1:
Input: cost = [10, 15, 20]
Output: 15
Explanation: Cheapest is start on cost[1], pay that cost and go to the top.
Example 2:
Input: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
Output: 6
Explanation: Cheapest is start on cost[0], and only step on 1s, skipping cost[3].
Note:
cost will have a length in the range [2, 1000].
Every cost[i] will be an integer in the range [0, 999].
 */

package main

import "fmt"


func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minCostClimbingStairs(cost []int) int {
	f1, f2 := 0, 0
	for i := len(cost)-1; i >= 0; i-- {
		f0 := cost[i] + minInt(f1, f2)
		f2 = f1
		f1 = f0
	}
	return minInt(f1, f2)
}
/*
func minHelper(cost []int, pos int) int {
	if pos == 0 {
		return cost[0]
	}
	if pos == 1 {
		return cost[1]
	}
	return minInt(minHelper(cost, pos-1) + cost[pos], minHelper(cost, pos-2) + cost[pos])
}

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 1 || len(cost) == 0 {
		return 0
	}
	if len(cost) == 2 {
		return minInt(cost[0], cost[1])
	}
	a := minHelper(cost, len(cost)-2)
	b := minHelper(cost, len(cost)-3)
	return minInt(a, b + cost[len(cost)-1])
}
*/

func main() {
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}