/*
Consider all the leaves of a binary tree.  From left to right order,
the values of those leaves form a leaf value sequence.



For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).

Two binary trees are considered leaf-similar if their leaf value sequence is the same.

Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.



Note:

Both of the given trees will have between 1 and 100 nodes.
 */

package main

func getLeafSequence(root *TreeNode, seq *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*seq = append(*seq, root.Val)
		return
	}
	getLeafSequence(root.Left, seq)
	getLeafSequence(root.Right, seq)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	seq1 := make([]int, 0)
	seq2 := make([]int, 0)
	getLeafSequence(root1, &seq1)
	getLeafSequence(root2, &seq2)
	if len(seq1) != len(seq2) {
		return false
	}
	for i := 0; i < len(seq1); i++ {
		if seq1[i] != seq2[i] {
			return false
		}
	}
	return true
}

[119,113,null,11,30,43,76,15,60,67,74]
[11,69,60,115,66,15,60,67,74,null,76]