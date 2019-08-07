/*
Given the root of a binary tree, each node in the tree has a distinct value.

After deleting all nodes with a value in to_delete, we are left with a forest (a disjoint union of trees).

Return the roots of the trees in the remaining forest.  You may return the result in any order.



Example 1:



Input: root = [1,2,3,4,5,6,7], to_delete = [3,5]
Output: [[1,2,null,4],[6],[7]]


Constraints:

The number of nodes in the given tree is at most 1000.
Each node has a distinct value between 1 and 1000.
to_delete.length <= 1000
to_delete contains distinct values between 1 and 1000.
 */

package main

func delNodesHelper(res *[]*TreeNode, root *TreeNode, dict map[int]bool, parentGone bool, parent *TreeNode, direct bool) {
	if root == nil {
		return
	}
	if dict[root.Val] {
		delNodesHelper(res, root.Left, dict, true, root, true)
		delNodesHelper(res, root.Right, dict, true, root, false)
		if parent != nil {
			if direct {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
		}
	} else {
		if parentGone {
			*res = append(*res, root)
		}
		delNodesHelper(res, root.Left, dict, false, root, true)
		delNodesHelper(res, root.Right, dict, false, root, false)
	}
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	dict := make(map[int]bool)
	res := make([]*TreeNode, 0)
	for i := 0; i < len(to_delete); i++ {
		dict[to_delete[i]] = true
	}
	if len(to_delete) == 0 {
		return res
	}
	delNodesHelper(&res, root, dict, true, nil, true)
	return res
}
