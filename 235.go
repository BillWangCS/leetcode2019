/*
Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.

According to the definition of LCA on Wikipedia:
“The lowest common ancestor is defined between two nodes p and q as the lowest node in T
that has both p and q as descendants (where we allow a node to be a descendant of itself).”

Given binary search tree:  root = [6,2,8,0,4,7,9,null,null,3,5]




Example 1:

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.
Example 2:

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.


Note:

All of the nodes' values will be unique.
p and q are different and both values will exist in the BST.
 */

package main

func getPath(root, p *TreeNode, path *[]*TreeNode) {
	*path = append(*path, root)
	if root.Val == p.Val {
		return
	}
	if root.Val > p.Val {
		getPath(root.Left, p, path)
	} else {
		getPath(root.Right, p, path)
	}
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	res := root
	path1 := make([]*TreeNode, 0)
	path2 := make([]*TreeNode, 0)
	getPath(root, p, &path1)
	getPath(root, q, &path2)
	for i := 0; i < minInt(len(path1), len(path2)); i++ {
		if path1[i] == path2[i] {
			res = path1[i]
		} else {
			break
		}
	}
	return res
}