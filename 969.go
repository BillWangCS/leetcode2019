/*
Given an array A, we can perform a pancake flip: We choose some positive integer k <= A.length,
then reverse the order of the first k elements of A.
We want to perform zero or more pancake flips (doing them one after another in succession) to sort the array A.

Return the k-values corresponding to a sequence of pancake flips that sort A.
Any valid answer that sorts the array within 10 * A.length flips will be judged as correct.



Example 1:

Input: [3,2,4,1]
Output: [4,2,4,3]
Explanation:
We perform 4 pancake flips, with k values 4, 2, 4, and 3.
Starting state: A = [3, 2, 4, 1]
After 1st flip (k=4): A = [1, 4, 2, 3]
After 2nd flip (k=2): A = [4, 1, 2, 3]
After 3rd flip (k=4): A = [3, 2, 1, 4]
After 4th flip (k=3): A = [1, 2, 3, 4], which is sorted.
Example 2:

Input: [1,2,3]
Output: []
Explanation: The input is already sorted, so there is no need to flip anything.
Note that other answers, such as [3, 3], would also be accepted.


Note:

1 <= A.length <= 100
A[i] is a permutation of [1, 2, ..., A.length]
 */

package main

import "fmt"

/*利用DFS的超时解法感觉自己好牛逼（有限状态自动机)
func isSorted(A []int) bool {
	if len(A) <= 1 {
		return true
	}
	for i := 1; i < len(A); i++ {
		if A[i] < A[i - 1] {
			return false
		}
	}
	return true
}

func arrToString(A []int) string {
	res := ""
	for _, v := range A {
		res = res + " " + strconv.Itoa(v)
	}
	return res
}

func dfs_helper(A []int, step int, parent int, res *[]int, visited *map[string]bool) bool {
	if step > len(A) {
		return false
	}
	if isSorted(A) {
		return true
	}
	for i := 1; i < len(A); i++ {
		if i == parent {
			continue
		}
		*res = append(*res, i+1)
		B := reverseK(A, i+1)
		sb := arrToString(B)
		if (*visited)[sb] {
			continue
		}
		(*visited)[arrToString(B)] = true
		fmt.Println(sb, i+1)
		if dfs_helper(B, step+1, i, res, visited) {
			return true
		}
		*res = (*res)[:len(*res)-1]
	}
	return false
}

func reverseK(A []int, k int) []int {
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		res = append(res, A[k - i - 1])
	}
	for i := k; i < len(A); i++ {
		res = append(res, A[i])
	}
	return res
}

func pancakeSort(A []int) []int {
	res := make([]int, 0)
	if isSorted(A) {
		return res
	}
	step := 0
	visited := make(map[string]bool)
	for i := 1; i < len(A); i++ {
		res = append(res, i+1)
		B := reverseK(A, i+1)
		sb := arrToString(B)
		if visited[sb] {
			continue
		}
		visited[sb] = true
		if dfs_helper(B, step+1, i, &res, &visited) {
			return res
		}
		res = res[:len(res)-1]
	}
	return res
}
*/

func isSorted(A []int) bool {
	if len(A) <= 1 {
		return true
	}
	for i := 1; i < len(A); i++ {
		if A[i] < A[i - 1] {
			return false
		}
	}
	return true
}

func reverseK(A []int, k int) []int {
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		res = append(res, A[k - i - 1])
	}
	for i := k; i < len(A); i++ {
		res = append(res, A[i])
	}
	return res
}

func getPos(A []int, n int) int {
	for i, v := range A {
		if v == n {
			return i
		}
	}
	return -1
}

func pancakeSort(A []int) []int {
	res := make([]int, 0)
	if isSorted(A) {
		return res
	}
	for i := len(A); i > 0; i-- {
		p := getPos(A, i)
		A = reverseK(A, p + 1)
		res = append(res, p + 1)
		A = reverseK(A, i)
		res = append(res, i)
	}
	return res
}

func main() {
	//fmt.Println(reverseK([]int{3,2,4,1}, 4))
	fmt.Println(pancakeSort([]int{10,5,1,6,3,8,2,4,7,9}))
}