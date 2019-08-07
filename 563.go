/*
Given a binary tree, return the tilt of the whole tree.

The tilt of a tree node is defined as the absolute difference between
the sum of all left subtree node values and the sum of all right subtree node values. Null node has tilt 0.

The tilt of the whole tree is defined as the sum of all nodes' tilt.

Example:
Input:
         1
       /   \
      2     3
Output: 1
Explanation:
Tilt of node 2 : 0
Tilt of node 3 : 0
Tilt of node 1 : |2-3| = 1
Tilt of binary tree : 0 + 0 + 1 = 1
Note:

The sum of node values in any subtree won't exceed the range of 32-bit integer.
All the tilt values won't exceed the range of 32-bit integer.
 */

package main

func sumOfTree(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	*sum += root.Val
	sumOfTree(root.Left, sum)
	sumOfTree(root.Right, sum)
}

func getTilt(root *TreeNode, tilt *int) {
	if root == nil {
		return
	}
	sumLeft, sumRight := 0, 0
	sumOfTree(root.Left, &sumLeft)
	sumOfTree(root.Right, &sumRight)
	*tilt += abs(sumLeft - sumRight)
	getTilt(root.Left, tilt)
	getTilt(root.Right, tilt)
}

func findTilt(root *TreeNode) int {
	res := 0
	getTilt(root, &res)
	return res
}