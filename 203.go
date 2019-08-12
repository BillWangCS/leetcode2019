/*
Remove all elements from a linked list of integers that have value val.

Example:

Input:  1->2->6->3->4->5->6, val = 6
Output: 1->2->3->4->5
 */

package main

func removeElements(head *ListNode, val int) *ListNode {
	prev := &ListNode{val-1, head}
	cur := head
	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
			cur = cur.Next
			if head.Val == val {
				head = head.Next
			}
		} else {
			prev = prev.Next
			cur = cur.Next
		}
	}
	return head
}
