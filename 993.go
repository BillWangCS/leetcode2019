/*
In a binary tree, the root node is at depth 0, and children of each depth k node are at depth k+1.

Two nodes of a binary tree are cousins if they have the same depth, but have different parents.

We are given the root of a binary tree with unique values, and the values x and y of two different nodes in the tree.

Return true if and only if the nodes corresponding to the values x and y are cousins.
 */

package main

type NodeInfo struct {
	Depth  int
	Parent int
}

func dfsHelper(root *TreeNode, x int, y int, resArr *[]NodeInfo, depth int, parent int) {
	if root == nil {
		return
	}
	if root.Val == x || root.Val == y {
		*resArr = append(*resArr, NodeInfo{depth, parent})
	}
	dfsHelper(root.Left, x, y, resArr, depth+1, root.Val)
	dfsHelper(root.Right, x, y, resArr, depth+1, root.Val)
}

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil || root.Val == x || root.Val == y {
		return false
	}
	resultArr := make([]NodeInfo, 0)
	dfsHelper(root.Left, x, y, &resultArr, 0, root.Val)
	dfsHelper(root.Right, x, y, &resultArr, 0, root.Val)
	if len(resultArr) != 2 {
		return false
	}
	if resultArr[0].Parent != resultArr[1].Parent && resultArr[0].Depth == resultArr[1].Depth {
		return true
	}
	return false
}
