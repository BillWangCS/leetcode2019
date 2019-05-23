/*
Given a binary tree, each node has value 0 or 1.
Each root-to-leaf path represents a binary number starting with the most significant bit.
For example, if the path is 0 -> 1 -> 1 -> 0 -> 1, then this could represent 01101 in binary, which is 13.

For all leaves in the tree, consider the numbers represented by the path from the root to that leaf.

Return the sum of these numbers.



Example 1:



Input: [1,0,1,0,1,0,1]
Output: 22
Explanation: (100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22


Note:

The number of nodes in the tree is between 1 and 1000.
node.val is 0 or 1.
The answer will not exceed 2^31 - 1.
 */

package main

func binaryNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res = res * 2 + v
	}
	return res
}

func dfs_1022(root *TreeNode, sum *int, path *[]int) {
	if root.Left == nil && root.Right == nil {
		*path = append(*path, root.Val)
		*sum += binaryNumber(*path)
		*path = (*path)[:len(*path)-1]
		return
	}
	*path = append(*path, root.Val)
	if root.Left != nil {
		dfs_1022(root.Left, sum, path)
	}
	if root.Right != nil {
		dfs_1022(root.Right, sum, path)
	}
	*path = (*path)[:len(*path)-1]
}

func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	path := make([]int, 0)
	dfs_1022(root, &sum, &path)
	return sum
}


