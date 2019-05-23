/*
A complete binary tree is a binary tree in which every level, except possibly the last, is completely filled,
and all nodes are as far left as possible.

Write a data structure CBTInserter that is initialized with a complete binary tree and supports the following operations:

CBTInserter(TreeNode root) initializes the data structure on a given tree with head node root;
CBTInserter.insert(int v) will insert a TreeNode into the tree with value node.val = v so that the tree remains complete,
and returns the value of the parent of the inserted TreeNode;
CBTInserter.get_root() will return the head node of the tree.


Example 1:

Input: inputs = ["CBTInserter","insert","get_root"], inputs = [[[1]],[2],[]]
Output: [null,1,[1,2]]
Example 2:

Input: inputs = ["CBTInserter","insert","insert","get_root"], inputs = [[[1,2,3,4,5,6]],[7],[8],[]]
Output: [null,3,4,[1,2,3,4,5,6,7,8]]


Note:

The initial given tree is complete and contains between 1 and 1000 nodes.
CBTInserter.insert is called at most 10000 times per test case.
Every value of a given or inserted node is between 0 and 5000.

 */

package main

/*
Better solution Deque:
First, perform a breadth-first search to populate the deque with nodes that have 0 or 1 children, in number order.

Now when inserting a node, the parent is the first element of deque, and we add this new node to our deque.
 */

type CBTInserter struct {
	root *TreeNode
}


func Constructor(root *TreeNode) CBTInserter {
	return CBTInserter{root}
}


func (this *CBTInserter) Insert(v int) int {
	if this.root == nil {
		this.root = &TreeNode{Val: v}
		return -1
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, this.root)
	var parentNode *TreeNode
	for len(queue) > 0 {
		l := len(queue)
		for l > 0 {
			parentNode = queue[0]
			if parentNode.Left == nil {
				parentNode.Left = &TreeNode{Val: v}
				return parentNode.Val
			} else if parentNode.Right == nil {
				parentNode.Right = &TreeNode{Val: v}
				return parentNode.Val
			} else {
				queue = append(queue, parentNode.Left)
				queue = append(queue, parentNode.Right)
			}
			queue = queue[1:]
			l--
		}
	}
}


func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}