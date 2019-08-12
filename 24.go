/*
Given a linked list, swap every two adjacent nodes and return its head.

You may not modify the values in the list's nodes, only nodes itself may be changed.



Example:

Given 1->2->3->4, you should return the list as 2->1->4->3.
 */

package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	odd := make([]*ListNode, 0)
	even := make([]*ListNode, 0)
	i := 0
	for head != nil {
		if i % 2 == 0 {
			even = append(even, head)
		} else {
			odd = append(odd, head)
		}
		i++
		head = head.Next
	}
	for i := 0; i < len(odd); i++ {
		fmt.Println(odd[i].Val)
	}
	if len(even) == 0 {
		return nil
	}
	if len(odd) == 0 {
		return even[0]
	}
	if len(odd) != len(even) {
		for i := 0; i < len(odd); i++ {
			fmt.Println(odd[i].Val, even[i].Val)
			odd[i].Next = even[i]
			if i+1 < len(odd) {
				even[i].Next = odd[i+1]
			} else {
				even[i].Next = even[i+1]
			}
		}
	} else {
		for i := 0; i < len(odd); i++ {
			fmt.Println(odd[i].Val, even[i].Val)
			odd[i].Next = even[i]
			if i+1 < len(odd) {
				even[i].Next = odd[i+1]
			} else {
				even[i].Next = nil
			}
		}
	}
	return odd[0]
}
