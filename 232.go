/*
Implement the following operations of a queue using stacks.

push(x) -- Push element x to the back of queue.
pop() -- Removes the element from in front of queue.
peek() -- Get the front element.
empty() -- Return whether the queue is empty.
Example:

MyQueue queue = new MyQueue();

queue.push(1);
queue.push(2);
queue.peek();  // returns 1
queue.pop();   // returns 1
queue.empty(); // returns false
Notes:

You must use only standard operations of a stack -- which means only push to top, peek/pop from top, size, and is empty operations are valid.
Depending on your language, stack may not be supported natively. You may simulate a stack by using a list or deque (double-ended queue), as long as you use only standard operations of a stack.
You may assume that all operations are valid (for example, no pop or peek operations will be called on an empty queue).
 */

package main

type MyQueue struct {
	queue []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	queue := make([]int, 0)
	return MyQueue{queue}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.queue = append(this.queue, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.queue) == 0 {
		return -1
	}
	res := this.queue[0]
	this.queue = this.queue[1:]
	return res
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.queue[0]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.queue) <= 0
}