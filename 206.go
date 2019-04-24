/*
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
Follow up:

A linked list can be reversed either iteratively or recursively. Could you implement both?
 */

package main

import "fmt"

/**
 * Definition for singly-linked list.*/
type ListNode struct {
	Val  int
	Next *ListNode
}

/*Better solution iterate
public ListNode reverseList(ListNode head) {
    ListNode prev = null;
    ListNode curr = head;
    while (curr != null) {
        ListNode nextTemp = curr.next;
        curr.next = prev;
        prev = curr;
        curr = nextTemp;
    }
    return prev;
}
 */

 //Reversal
 func reverseList(head *ListNode) *ListNode {
 	if head == nil || head.Next == nil {
 		return head
	}
 	node := reverseList(head.Next)
 	head.Next.Next = head
 	head.Next = nil
 	return node
 }

 //Iterative
func getTail(head *ListNode) (*ListNode, *ListNode) {
	tail := head
	pre := head
	for tail.Next != nil {
		if pre.Next != nil && pre.Next.Next != nil {
			pre = pre.Next
		}
		tail = tail.Next
	}
	return pre, tail
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	h := head
	preTail1, tail1 := getTail(head)
	tail1.Next = head
	preTail1.Next = nil
	head = tail1
	for {
		preTail, tail := getTail(head)
		tail.Next = head.Next
		head.Next = tail
		head = tail
		preTail.Next = nil
		fmt.Println(tail.Val)
		if tail == h {
			break
		}
	}
	res := tail1
	for tail1.Next != nil {
		tail1 = tail1.Next
	}
	tail1.Next = h
	return res
}
