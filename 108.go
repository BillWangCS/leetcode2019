/*
Given an array where elements are sorted in ascending order, convert it to a height balanced BST.

For this problem,
a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example:

Given the sorted array: [-10,-3,0,5,9],

One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:

      0
     / \
   -3   9
   /   /
 -10  5
 */

package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func dfs_108(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	pos := len(nums) / 2
	root := TreeNode{Val: nums[pos]}
	if pos + 1 < len(nums) {
		root.Right = dfs_108(nums[pos+1:])
	}
	root.Left = dfs_108(nums[:pos])
	return &root
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return dfs_108(nums)
}

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

func main() {
	fmt.Println(inorderTraversal(sortedArrayToBST([]int{-10,-3,0,5,9})))
}