/*
Given a non-empty integer array of size n, find the minimum number of moves required to make all array elements equal,
where a move is incrementing n - 1 elements by 1.

Example:

Input:
[1,2,3]

Output:
3

Explanation:
Only three moves are needed (remember each move increments two elements):

[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
 */

 /* Better solution
 func findMin(nums []int) int{
    min:=nums[0]

    for i:=1;i<len(nums);i++{
        if nums[i] < min {
            min = nums[i]
        }
    }
    return min
}
func minMoves(nums []int) int {

    diff := 0
    min := findMin(nums)

    for i:=0;i<len(nums);i++{
        old := nums[i]- min
        diff =diff+old
    }

    return diff
}
  */

package main

import (
	"fmt"
	"sort"
)

func findMax(nums []int) (int, int)  {
	pos := 0
	diff := 0
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			diff = nums[i] - max
			max = nums[i]
			pos = i
		}
	}
	return pos, diff
}

func increaseElements(nums *[]int, pos, diff int) {
	for i := 0; i < len(*nums); i++ {
		if i != pos {
			(*nums)[i] += diff
		}
	}
}

func allElementsEqual(nums []int) bool {
	for i := 0; i + 1 < len(nums); i++ {
		if nums[i] != nums[i+1] {
			return false
		}
	}
	return true
}

func minMoves(nums []int) int {
	res := 0
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	for !allElementsEqual(nums) {
		pos, diff := findMax(nums)
		res += diff
		increaseElements(&nums, pos, diff)
	}
	return res
}

func main() {
	fmt.Println(minMoves([]int{83,86,77,15,93,35,86,92,49,21}))
}
