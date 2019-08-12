/*
Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:

Given n will always be valid.

Follow up:

Could you do this in one pass?
 */

package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	listArr := make([]*ListNode, 0)
	h := head
	for h != nil {
		listArr = append(listArr, h)
		h = h.Next
	}
	l := len(listArr)
	if n > l || n <= 0 {
		return head
	}
	if n == l {
		if !(l == 1 && n == 1) {
			return listArr[1]
		} else {
			return nil
		}
	}
	listArr[l-n-1].Next = listArr[l-n].Next
	return head
}
