/*
Given a binary tree rooted at root, the depth of each node is the shortest distance to the root.

A node is deepest if it has the largest depth possible among any node in the entire tree.

The subtree of a node is that node, plus the set of all descendants of that node.

Return the node with the largest depth such that it contains all the deepest nodes in its subtree.



Example 1:

Input: [3,5,1,6,2,0,8,null,null,7,4]
Output: [2,7,4]
Explanation:



We return the node with value 2, colored in yellow in the diagram.
The nodes colored in blue are the deepest nodes of the tree.
The input "[3, 5, 1, 6, 2, 0, 8, null, null, 7, 4]" is a serialization of the given tree.
The output "[2, 7, 4]" is a serialization of the subtree rooted at the node with value 2.
Both the input and output have TreeNode type.


Note:

The number of nodes in the tree will be between 1 and 500.
The values of each node are unique.
 */

package main

func maxInteger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxInteger(depth(root.Left), depth(root.Right)) + 1
}

func getPaths(root *TreeNode, dep int, curDepth *int, path *[]*TreeNode, paths *[][]*TreeNode) {
	if root == nil {
		return
	}
	*path = append(*path, root)
	if dep == *curDepth {
		curPath := make([]*TreeNode, 0)
		curPath = append(curPath, *path...)
		*paths = append(*paths, curPath)
		*path = (*path)[:len(*path)-1]
		return
	}
	*curDepth++
	getPaths(root.Left, dep, curDepth, path, paths)
	getPaths(root.Right, dep, curDepth, path, paths)
	*path = (*path)[:len(*path)-1]
	*curDepth--
	return
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	dep := depth(root)
	if dep == 0 || dep == 1 {
		return root
	}
	path := make([]*TreeNode, 0)
	paths := make([][]*TreeNode, 0)
	curDepth := 1
	getPaths(root, dep, &curDepth, &path, &paths)
	var tmp *TreeNode
	if len(paths) == 1 {
		return paths[0][len(paths[0])-1]
	}
	for j := 0; j < len(paths[0]); j++ {
		tmp = paths[0][j]
		for i := 1; i < len(paths); i++ {
			if tmp != paths[i][j] {
				return paths[i][j-1]
			}
		}
	}
	return root
}