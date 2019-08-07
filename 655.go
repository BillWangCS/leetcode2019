/*
Print a binary tree in an m*n 2D string array following these rules:

The row number m should be equal to the height of the given binary tree.
The column number n should always be an odd number.
The root node's value (in string format) should be put in the exactly middle of the first row it can be put.
The column and the row where the root node belongs will separate the rest space into two parts (left-bottom part and right-bottom part).
You should print the left subtree in the left-bottom part and print the right subtree in the right-bottom part.
The left-bottom part and the right-bottom part should have the same size. Even if one subtree is none while the other is not,
you don't need to print anything for the none subtree but still need to leave the space as large as that for the other subtree.
However, if two subtrees are none, then you don't need to leave space for both of them.
Each unused space should contain an empty string "".
Print the subtrees following the same rules.
Example 1:
Input:
     1
    /
   2
Output:
[["", "1", ""],
 ["2", "", ""]]
Example 2:
Input:
     1
    / \
   2   3
    \
     4
Output:
[["", "", "", "1", "", "", ""],
 ["", "2", "", "", "", "3", ""],
 ["", "", "4", "", "", "", ""]]
Example 3:
Input:
      1
     / \
    2   5
   /
  3
 /
4
Output:

[["",  "",  "", "",  "", "", "", "1", "",  "",  "",  "",  "", "", ""]
 ["",  "",  "", "2", "", "", "", "",  "",  "",  "",  "5", "", "", ""]
 ["",  "3", "", "",  "", "", "", "",  "",  "",  "",  "",  "", "", ""]
 ["4", "",  "", "",  "", "", "", "",  "",  "",  "",  "",  "", "", ""]]
Note: The height of binary tree is in the range of [1, 10].
 */

package main

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxInteger(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func heightOfTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxInteger(heightOfTree(root.Left), heightOfTree(root.Right)) + 1
}

func printTree(root *TreeNode) [][]string {
	res := make([][]string, 0)
	if root == nil {
		return res
	}
	h := heightOfTree(root)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	layer := 0
	for len(queue) > 0 && layer < h {
		l := len(queue)
		numPre := (1 << (uint(h-layer-1))) - 1
		path := make([]string, 0)
		for l > 0 {
			n := queue[0]
			if n != nil {
				queue = append(queue, n.Left)
				queue = append(queue, n.Right)
			} else {
				queue = append(queue, nil)
				queue = append(queue, nil)
			}
			queue = queue[1:]
			if len(path) == 0 {
				for i := 0; i < numPre; i++ {
					path = append(path, "")
				}
			} else {
				numInterval := (1 << (uint(h-layer))) - 1
				for i := 0; i < numInterval; i++ {
					path = append(path, "")
				}
			}
			if n == nil {
				path = append(path, "")
			} else {
				path = append(path, strconv.Itoa(n.Val))
			}
			l--
		}
		for i := 0; i < numPre; i++ {
			path = append(path, "")
		}
		res = append(res, path)
		layer++
	}
	return res
}