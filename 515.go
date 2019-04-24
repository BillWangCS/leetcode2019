/*
You need to find the largest value in each row of a binary tree.

Example:
Input:

          1
         / \
        3   2
       / \   \
      5   3   9

Output: [1, 3, 9]
 */

package main

import "math"

func largestValues(root *TreeNode) []int {
	queue := make([]*TreeNode, 0)
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		max := math.MinInt32
		for l > 0 {
			t := queue[0]
			queue = queue[1:]
			if t.Val > max {
				max = t.Val
			}
			if t.Left != nil {
				queue = append(queue, t.Left)
			}
			if t.Right != nil {
				queue = append(queue, t.Right)
			}
			l--
		}
		res = append(res, max)
	}
	return res
}