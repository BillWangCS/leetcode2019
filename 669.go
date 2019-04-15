/*
Given a binary search tree and the lowest and highest boundaries as L and R,
trim the tree so that all its elements lies in [L, R] (R >= L).
You might need to change the root of the tree,
so the result should return the new root of the trimmed binary search tree.

Example 1:
Input:
    1
   / \
  0   2

  L = 1
  R = 2

Output:
    1
      \
       2
Example 2:
Input:
    3
   / \
  0   4
   \
    2
   /
  1

  L = 1
  R = 3

Output:
      3
     /
   2
  /
 1
 */

package main

func getVals(root *TreeNode, L, R int, res *[]int) {
	if root == nil {
		return
	}
	if root.Val >= L && root.Val <= R {
		*res = append(*res, root.Val)
	}
	getVals(root.Left, L, R, res)
	getVals(root.Right, L, R, res)
}

func insertBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
	}
	if root.Val > val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
		} else {
			insertBST(root.Left, val)
		}
	}
	if root.Val < val {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
		} else {
			insertBST(root.Right, val)
		}
	}
	return root
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	var res *TreeNode
	vals := make([]int, 0)
	getVals(root, L, R, &vals)
	for _, v := range vals {
		res = insertBST(res, v)
	}
	return res
}

