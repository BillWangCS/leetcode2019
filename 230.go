/*
Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.

Note:
You may assume k is always valid, 1 ≤ k ≤ BST's total elements.

Example 1:

Input: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
Output: 1
Example 2:

Input: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
Output: 3
Follow up:
What if the BST is modified (insert/delete operations) often and you need to find the kth smallest frequently?
How would you optimize the kthSmallest routine?
 */

package main

func dfs_230(root *TreeNode, cnt *int, res *int, k int) {
	if root == nil {
		return
	}
	dfs_230(root.Left, cnt, res, k)
	(*cnt)++
	if *cnt == k {
		*res = root.Val
	}
	dfs_230(root.Right, cnt, res, k)
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	cnt := 0
	res := 0
	dfs_230(root, &cnt, &res, k)
	return res
}