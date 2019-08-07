/*
Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.

Note: A leaf is a node with no children.

Example:

Given the below binary tree and sum = 22,

      5
     / \
    4   8
   /   / \
  11  13  4
 /  \    / \
7    2  5   1
Return:

[
   [5,4,11,2],
   [5,8,4,5]
]
 */

package main

func dfs_113(root *TreeNode, target int, current int, path *[]int, res *[][]int) {
	if root == nil {
		return
	}
	current += root.Val
	*path = append(*path, root.Val)
	if root.Left == nil && root.Right == nil {
		if current == target {
			tmp := make([]int, 0)
			tmp = append(tmp, *path...)
			*res = append(*res, tmp)
		}
	}
	dfs_113(root.Left, target, current, path, res)
	dfs_113(root.Right, target, current, path, res)
	*path = (*path)[:len(*path)-1]
}

func pathSum(root *TreeNode, sum int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)
	dfs_113(root, sum, 0, &path, &res)
	return res
}