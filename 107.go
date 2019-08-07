/*
Given a binary tree, return the bottom-up level order traversal of its nodes' values.
(ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its bottom-up level order traversal as:
[
  [15,7],
  [9,20],
  [3]
]
 */

package main

func reverseArr2(A *[][]int) {
	i := 0
	j := len(*A) - 1
	for i <= j {
		(*A)[i], (*A)[j] = (*A)[j], (*A)[i]
		i++
		j--
	}
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		layer := make([]int, 0)
		for l > 0 {
			node := queue[0]
			layer = append(layer, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			l--
		}
		res = append(res, layer)
	}
	reverseArr2(&res)
	return res
}