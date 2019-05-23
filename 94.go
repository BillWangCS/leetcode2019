/*
Given a binary tree, return the inorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
Follow up: Recursive solution is trivial, could you do it iteratively?
 */

package main

func dfs_94(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	dfs_94(root.Left, res)
	*res = append(*res, root.Val)
	dfs_94(root.Right, res)
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	dfs_94(root, &res)
	return res
}