/*
We are given a linked list with head as the first node.  Let's number the nodes in the list: node_1, node_2, node_3, ... etc.

Each node may have a next larger value: for node_i, next_larger(node_i) is the node_j.val such that j > i,
node_j.val > node_i.val, and j is the smallest possible choice.  If such a j does not exist, the next larger value is 0.

Return an array of integers answer, where answer[i] = next_larger(node_{i+1}).

Note that in the example inputs (not outputs) below,
arrays such as [2,1,5] represent the serialization of a linked list with a head node value of 2,
second node value of 1, and third node value of 5.



Example 1:

Input: [2,1,5]
Output: [5,5,0]
Example 2:

Input: [2,7,4,3,5]
Output: [7,0,5,5,0]
Example 3:

Input: [1,7,5,1,9,2,5,1]
Output: [7,9,9,9,0,5,0,0]


Note:

1 <= node.val <= 10^9 for each node in the linked list.
The given list has length in the range [0, 10000].
 */

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Elem struct {
	Val int
	Index int
}

type Stack []Elem

func (obj *Stack) IsEmpty() bool {
	return len(*obj) == 0
}

func (obj *Stack) Peek() Elem {
	if !obj.IsEmpty() {
		return (*obj)[len(*obj)-1]
	}
	return Elem{}
}

func (obj *Stack) Push(elem Elem) {
	*obj = append(*obj, elem)
}

func (obj *Stack) Pop() Elem {
	res := Elem{}
	if !obj.IsEmpty() {
		res = obj.Peek()
		*obj = (*obj)[:len(*obj)-1]
	}
	return res
}


func nextLargerNodes(head *ListNode) []int {
	stk := Stack{}
	cur := head
	cnt := 0
	for cur != nil {
		cnt++
		cur = cur.Next
	}
	res := make([]int, cnt)
	if cnt == 0 {
		return res
	}
	pos := 0
	for head != nil {
		for !stk.IsEmpty() && head.Val > stk.Peek().Val {
			e := stk.Pop()
			res[e.Index] = head.Val
		}
		stk.Push(Elem{head.Val, pos})
		head = head.Next
		pos++
	}
	return res
}