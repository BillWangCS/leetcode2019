/*
Given a Binary Search Tree (BST),
convert it to a Greater Tree such that every key of the original BST is changed
to the original key plus sum of all keys greater than the original key in BST.

Example:

Input: The root of a Binary Search Tree like this:
              5
            /   \
           2     13

Output: The root of a Greater Tree like this:
             18
            /   \
          20     13
 */

package main

func dfs_538(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	dfs_538(root.Left, arr)
	*arr = append(*arr, root.Val)
	dfs_538(root.Right, arr)
}

func dfs_538_helper(root *TreeNode, arr *[]int, pos *int) {
	if root == nil {
		return
	}
	dfs_538_helper(root.Left, arr, pos)
	root.Val = (*arr)[*pos]
	*pos ++
	dfs_538_helper(root.Right, arr, pos)
}

func convertBST(root *TreeNode) *TreeNode {
	arr := make([]int, 0)
	dfs_538(root, &arr)
	for i := len(arr) - 2; i >= 0; i-- {
		arr[i] += arr[i+1]
	}
	pos := 0
	dfs_538_helper(root, &arr, &pos)
	return root
}