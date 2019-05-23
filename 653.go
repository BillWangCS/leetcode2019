/*
Given a Binary Search Tree and a target number,
return true if there exist two elements in the BST such that their sum is equal to the given target.

Example 1:

Input:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 9

Output: True


Example 2:

Input:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 28

Output: False
 */

package main

func dfs_653(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	dfs_653(root.Left, arr)
	*arr = append(*arr, root.Val)
	dfs_653(root.Right, arr)
}

func findTarget(root *TreeNode, k int) bool {
	arr := make([]int, 0)
	dfs_653(root, &arr)
	l := len(arr)
	s := 0
	e := l - 1
	for s < e {
		if arr[s] + arr[e] == k {
			return true
		}
		if arr[s] + arr[e] > k {
			e--
		} else {
			s++
		}
	}
	return false
}