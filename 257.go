/*
Given a binary tree, return all root-to-leaf paths.

Note: A leaf is a node with no children.

Example:

Input:

   1
 /   \
2     3
 \
  5

Output: ["1->2->5", "1->3"]

Explanation: All root-to-leaf paths are: 1->2->5, 1->3
 */

package main

import "strconv"

func dfs_257(root *TreeNode, path *[]*TreeNode, results *[]string) {
	if root == nil {
		return
	}
	*path = append(*path, root)
	if root.Left == nil && root.Right == nil {
		str := ""
		for i := 0; i < len(*path); i++ {
			str += strconv.Itoa((*path)[i].Val) + "->"
		}
		str = str[:len(str)-2]
		*results = append(*results, str)
	}
	dfs_257(root.Left, path, results)
	dfs_257(root.Right, path, results)
	*path = (*path)[:len(*path)-1]
}

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	path := make([]*TreeNode, 0)
	dfs_257(root, &path, &res)
	return res
}
