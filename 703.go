/*
Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order,
not the kth distinct element.

Your KthLargest class will have a constructor which accepts an integer k and an integer array nums,
which contains initial elements from the stream. For each call to the method KthLargest.add,
return the element representing the kth largest element in the stream.

Example:

int k = 3;
int[] arr = [4,5,8,2];
KthLargest kthLargest = new KthLargest(3, arr);
kthLargest.add(3);   // returns 4
kthLargest.add(5);   // returns 5
kthLargest.add(10);  // returns 5
kthLargest.add(9);   // returns 8
kthLargest.add(4);   // returns 8
Note:
You may assume that nums' length ≥ k-1 and k ≥ 1.
 */

package main

import "fmt"

type KthLargest struct {
	arr []int
	k int
}

func quickK(A *[]int, start, end int, K int) {
	if start >= end {
		return
	}
	left, right := start, end
	pivot := (*A)[(end + start)/2]
	for left <= right {
		for left <= right && (*A)[left] > pivot {
			left++
		}
		for left <= right && (*A)[right] < pivot {
			right--
		}
		if left <= right {
			(*A)[left], (*A)[right] = (*A)[right], (*A)[left]
			left++
			right--
		}
	}
	if K >= left {
		quickK(A, left, end, K)
	} else if K <= right {
		quickK(A, start, right, K)
	}
	return
}

func Constructor(k int, nums []int) KthLargest {
	return KthLargest{nums, k}
}


func (this *KthLargest) Add(val int) int {
	this.arr = append(this.arr, val)
	quickK(&this.arr, 0, len(this.arr)-1, this.k-1)
	return this.arr[this.k-1]
}

func main() {
	obj := Constructor(3, []int{4,5,8,2})
	obj.Add(3)
	fmt.Println(obj.arr)
	obj.Add(5)
	fmt.Println(obj.arr)
	obj.Add(10)
	fmt.Println(obj.arr)
	obj.Add(9)
	fmt.Println(obj.arr)
	obj.Add(4)
	fmt.Println(obj.arr)
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
[null,4,5,5,8,8]
 */