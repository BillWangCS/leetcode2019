/*
Merge two sorted linked lists and return it as a new list.
The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
 */

package main

import "fmt"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var res *ListNode
	if l1.Val < l2.Val {
		res = &ListNode{Val: l1.Val}
		l1 = l1.Next
	} else {
		res = &ListNode{Val: l2.Val}
		l2 = l2.Next
	}
	fmt.Println(res.Val)
	x := res
	for l1 != nil && l2 != nil {
		fmt.Println(l1.Val, l2.Val)
		if l1.Val < l2.Val {
			x.Next = &ListNode{Val: l1.Val}
			x = x.Next
			l1 = l1.Next
		} else {
			x.Next = &ListNode{Val: l2.Val}
			x = x.Next
			l2 = l2.Next
		}
	}
	for l1 != nil {
		x.Next = &ListNode{Val: l1.Val}
		x = x.Next
		l1 = l1.Next
	}
	for l2 != nil {
		x.Next = &ListNode{Val: l2.Val}
		x = x.Next
		l2 = l2.Next
	}
	return res
}
