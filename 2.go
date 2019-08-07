/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
 */

package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	l := len(num1)
	s := len(num2)
	c := 0
	res := ""
	if l < s {
		s, l = l, s
	}
	for i := 1; i <= l; i++ {
		sum := 0
		if i <= len(num1) {
			n1, _ := strconv.Atoi(string(num1[len(num1)-i]))
			sum += n1
		}
		if i <= len(num2) {
			n2, _ := strconv.Atoi(string(num2[len(num2)-i]))
			sum += n2
		}
		sum += c
		res = strconv.Itoa(sum%10) + res
		c = sum / 10
	}
	if c > 0 {
		res = strconv.Itoa(c) + res
	}
	return res
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := ""
	s2 := ""
	res := ListNode{}
	cur := &res
	for l1 != nil {
		s1 = strconv.Itoa(l1.Val) + s1
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = strconv.Itoa(l2.Val) + s2
		l2 = l2.Next
	}
	s := addStrings(s1, s2)
	fmt.Println(s1, s2, s)
	for i := len(s)-1; i >= 0; i-- {
		n, _ := strconv.Atoi(string(s[i]))
		cur.Next = &ListNode{Val: n, Next: nil}
		cur = cur.Next
	}
	return res.Next
}