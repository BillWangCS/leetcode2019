/*
Given a binary tree and a sum, determine if the tree has a root-to-leaf
path such that adding up all the values along the path equals the given sum.

Note: A leaf is a node with no children.

Example:

Given the below binary tree and sum = 22,

      5
     / \
    4   8
   /   / \
  11  13  4
 /  \      \
7    2      1
return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
 */

package main

func dfs_22(root *TreeNode, target int, current int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		if current + root.Val == target {
			return true
		} else {
			return false
		}
	}
	return dfs_22(root.Left, target, current+root.Val) || dfs_22(root.Right, target, current+root.Val)
}

func hasPathSum(root *TreeNode, sum int) bool {
	current := 0
	return dfs_22(root, sum, current)
}