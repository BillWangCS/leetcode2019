/*
Given the root of a tree, you are asked to find the most frequent subtree sum.
The subtree sum of a node is defined as the sum of all the node values formed by
the subtree rooted at that node (including the node itself). So what is the most frequent subtree sum value? If there is a tie,
return all the values with the highest frequency in any order.

Examples 1
Input:

  5
 /  \
2   -3
return [2, -3, 4], since all the values happen only once, return all of them in any order.
Examples 2
Input:

  5
 /  \
2   -5
return [2], since 2 happens twice, however -5 only occur once.
Note: You may assume the sum of values in any subtree is in the range of 32-bit signed integer.
 */

package main

func sumOfTree(root * TreeNode) int {
	if root == nil {
		return 0
	}
	return sumOfTree(root.Left) + sumOfTree(root.Right) + root.Val
}

func freqCount(root *TreeNode, dict *map[int]int, freq *int) {
	if root == nil {
		return
	}
	(*dict)[sumOfTree(root)]++
	if (*dict)[sumOfTree(root)] > *freq {
		*freq = (*dict)[sumOfTree(root)]
	}
	freqCount(root.Left, dict, freq)
	freqCount(root.Right, dict, freq)
}

func findFrequentTreeSum(root *TreeNode) []int {
	dict := make(map[int]int)
	res := make([]int, 0)
	maxFreq := -1
	freqCount(root, &dict, &maxFreq)
	for k, v := range dict {
		if v == maxFreq {
			res = append(res, k)
		}
	}
	return res
}