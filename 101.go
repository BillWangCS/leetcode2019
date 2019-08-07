/*
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3


But the following [1,2,2,null,3,null,3] is not:

    1
   / \
  2   2
   \   \
   3    3


Note:
Bonus points if you could solve it both recursively and iteratively.
 */

package main

func swapChild(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	swapChild(root.Left)
	swapChild(root.Right)
}

func isSameTree(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 != nil && root2 == nil {
		return false
	}
	if root1 == nil && root2 != nil {
		return false
	}
	return root1.Val == root2.Val && isSameTree(root1.Left, root2.Left) && isSameTree(root1.Right, root2.Right)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	swapChild(root.Right)
	return isSameTree(root.Left, root.Right)
}
