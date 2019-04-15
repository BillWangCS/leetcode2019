/*
For a binary tree T, we can define a flip operation as follows:
choose any node, and swap the left and right child subtrees.

A binary tree X is flip equivalent to a binary tree Y
if and only if we can make X equal to Y after some number of flip operations.

Write a function that determines whether two binary trees are flip equivalent.
The trees are given by root nodes root1 and root2.



Example 1:

Input: root1 = [1,2,3,4,5,6,null,null,null,7,8], root2 = [1,3,2,null,6,4,5,null,null,null,null,8,7]
Output: true
Explanation: We flipped at nodes with values 1, 3, and 5.
Flipped Trees Diagram


Note:

Each tree will have at most 100 nodes.
Each value in each tree will be a unique integer in the range [0, 99].
 */


package main

func flipTree(root1, root2 *TreeNode) {
	if root1 == nil || root2 == nil {
		return
	}
	if root1.Left == nil && root1.Right != nil && root2.Left != nil && root2.Right == nil && root1.Right.Val == root2.Left.Val {
		root1.Left = root1.Right
		root1.Right = nil
	} else if root1.Right == nil && root1.Left != nil && root2.Right != nil && root2.Left == nil && root1.Left.Val == root2.Right.Val {
		root1.Right = root1.Left
		root1.Left = nil
	} else if root1.Right != nil && root1.Left != nil && root2.Left != nil && root2.Right != nil && root1.Left.Val == root2.Right.Val && root1.Right.Val == root2.Left.Val {
		tmp := root1.Left
		root1.Left = root1.Right
		root1.Right = tmp
	}
	flipTree(root1.Left, root2.Left)
	flipTree(root1.Right, root2.Right)
}

func isSameTree2(r1 *TreeNode, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 != nil && r2 != nil {
		if r1.Val == r2.Val {
			return isSameTree2(r1.Left, r2.Left) && isSameTree(r1.Right, r2.Right)
		}
	}
	return false
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	flipTree(root1, root2)
	return isSameTree2(root1, root2)
}