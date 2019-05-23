/*
Design your implementation of the linked list. You can choose to use the singly linked list or the doubly linked list.
A node in a singly linked list should have two attributes: val and next.
val is the value of the current node, and next is a pointer/reference to the next node.
If you want to use the doubly linked list, you will need one more attribute prev to indicate the previous node in the linked list.
Assume all nodes in the linked list are 0-indexed.

Implement these functions in your linked list class:

get(index) : Get the value of the index-th node in the linked list. If the index is invalid, return -1.
addAtHead(val) : Add a node of value val before the first element of the linked list.
After the insertion, the new node will be the first node of the linked list.
addAtTail(val) : Append a node of value val to the last element of the linked list.
addAtIndex(index, val) : Add a node of value val before the index-th node in the linked list.
If index equals to the length of linked list, the node will be appended to the end of linked list.
If index is greater than the length, the node will not be inserted.
deleteAtIndex(index) : Delete the index-th node in the linked list, if the index is valid.
Example:

MyLinkedList linkedList = new MyLinkedList();
linkedList.addAtHead(1);
linkedList.addAtTail(3);
linkedList.addAtIndex(1, 2);  // linked list becomes 1->2->3
linkedList.get(1);            // returns 2
linkedList.deleteAtIndex(1);  // now the linked list is 1->3
linkedList.get(1);            // returns 3
Note:

All values will be in the range of [1, 1000].
The number of operations will be in the range of [1, 1000].
Please do not use the built-in LinkedList library.
 */

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	Head *ListNode
	Len int
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	h := this.Head
	cnt := 0
	for h != nil {
		if index == cnt {
			return h.Val
		}
		cnt++
		h = h.Next
	}
	return -1
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	h := ListNode{val, this.Head}
	this.Head = &h
	this.Len++
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	t := ListNode{val, nil}
	h := this.Head
	if this.Len == 0 {
		this.Head = &t
		return
	}
	for h.Next != nil {
		h = h.Next
	}
	h.Next = &t
	this.Len++
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	prev := &ListNode{0, this.Head}
	h := this.Head
	var last *ListNode
	cnt := 0
	for h != nil {
		if index == cnt {
			last = h
			break
		}
		prev = prev.Next
		h = h.Next
		cnt++
	}
	if h == nil && index == cnt {
		this.AddAtTail(val)
		return
	}
	if index > cnt {
		return
	}
	n := ListNode{val, last}
	prev.Next = &n
	this.Len++
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this == nil || index < 0 {
		return
	}
	prev := &ListNode{0, this.Head}
	h := this.Head
	cnt := 0
	for h != nil {
		if index == cnt {
			prev.Next = h.Next
			break
		}
		prev = prev.Next
		h = h.Next
		cnt++
	}
	if index <= 0 {
		this.Head = prev.Next
	}
}

func main() {
	obj := Constructor()
//	obj.AddAtHead(1)
	//obj.AddAtTail(3)
	obj.AddAtIndex(-1,0)
	for i := 0; i < obj.Len; i++ {
		fmt.Println(obj.Get(i))
	}
}
/*
["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]
[[],[1],[3],[1,2],[1],[1],[1]]*/

/*
这个case我觉得不make sense不过为了AC我也简单迎合了AddAtIndex的index小于0的情况
["MyLinkedList","addAtIndex","get","deleteAtIndex"]
[[],[-1,1],[0],[-1]]
 */