/*
Given a binary tree, find the leftmost value in the last row of the tree.

Example 1:
Input:

    2
   / \
  1   3

Output:
1
Example 2:
Input:

        1
       / \
      2   3
     /   / \
    4   5   6
       /
      7

Output:
7
 */

package main

func maxInteger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func treeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxInteger(treeDepth(root.Left), treeDepth(root.Right)) + 1
}

func findBottomLeftValue(root *TreeNode) int {
	depth := treeDepth(root)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	cnt := 0
	for len(queue) > 0 {
		cnt++
		l := len(queue)
		for l > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				if cnt == depth {
					return node.Left.Val
				}
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				if cnt == depth {
					return node.Right.Val
				}
				queue = append(queue, node.Right)
			}
			l--
		}
	}
	return 0
}