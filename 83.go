/*
Given a sorted linked list, delete all duplicates such that each element appear only once.

Example 1:

Input: 1->1->2
Output: 1->2
Example 2:

Input: 1->1->2->3->3
Output: 1->2->3
 */

package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dict := make(map[int]bool)
	prev := head
	cur := head.Next
	dict[head.Val] = true
	for cur != nil {
		if _, ok := dict[cur.Val]; !ok {
			dict[cur.Val] = true
			prev = prev.Next
		} else {
			prev.Next = cur.Next
		}
		cur = cur.Next
	}
	return head
}