/*
A full binary tree is a binary tree where each node has exactly 0 or 2 children.

Return a list of all possible full binary trees with N nodes.
Each element of the answer is the root node of one possible tree.

Each node of each tree in the answer must have node.val = 0.

You may return the final list of trees in any order.



Example 1:

Input: 7
Output: [[0,0,0,null,null,0,0,null,null,0,0],
[0,0,0,null,null,0,0,0,0],
[0,0,0,0,0,0,0],
[0,0,0,0,0,null,null,null,null,0,0],[0,0,0,0,0,null,null,0,0]]
Explanation:



Note:

1 <= N <= 20
 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	layer := make([]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		layer = nil
		for l > 0 {
			t := queue[0]
			queue = queue[1:]
			if t.Left != nil {
				queue = append(queue, t.Left)
			}
			if t.Right != nil {
				queue = append(queue, t.Right)
			}
			layer = append(layer, t.Val)
			l--
		}
		res = append(res, layer)
	}
	return res
}

var x []*TreeNode

func isSameTree(r1 *TreeNode, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 != nil && r2 != nil {
		if r1.Val == r2.Val {
			return isSameTree(r1.Left, r2.Left) && isSameTree(r1.Right, r2.Right)
		}
	}
	return false
}

func exist(root *TreeNode, results []*TreeNode) bool {
	fmt.Println("bbbbb")
	for _, v := range results {
		if isSameTree(root, v) {
			return true
		}
	}
	return false
}

func insertFullTree(t *TreeNode, n int, root *TreeNode, results []*TreeNode) []*TreeNode {
	if n <= 0 {
		return results
	}
	if t.Left == nil {
		t.Left = &TreeNode{Val: 0}
		t.Right = &TreeNode{Val: 0}
		n = n - 2
		if n <= 0 {
			fmt.Println(levelOrder(root), n)
			if exist(root, results) {
				fmt.Println("2222")
				t.Left = nil
				t.Right = nil
				n += 2
			} else {
				fmt.Println("aaaa", n)
				r := &TreeNode{Val: 0, Left: root.Left, Right: root.Right}
				fmt.Println(r, root)
				results = append(results, r)
				t.Left = nil
				t.Right = nil
				n += 2
			}
		}
	}
	if t.Left != nil {
		results = insertFullTree(t.Left, n, root, results)
		results = insertFullTree(t.Right, n, root, results)
	}
	return results
}

func allPossibleFBT(N int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if N%2 == 0 {
		return res
	}
	root := &TreeNode{Val: 0}
	if N == 1 {
		return []*TreeNode{&TreeNode{Val: 0}}
	}
	return insertFullTree(root, N-1, root, res)
}

func main() {
	fmt.Println(allPossibleFBT(7))
}
