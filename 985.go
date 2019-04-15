/*
We have an array A of integers, and an array queries of queries.

For the i-th query val = queries[i][0], index = queries[i][1], we add val to A[index].
Then, the answer to the i-th query is the sum of the even values of A.

(Here, the given index = queries[i][1] is a 0-based index, and each query permanently modifies the array A.)

Return the answer to all queries.  Your answer array should have answer[i] as the answer to the i-th query.



Example 1:

Input: A = [1,2,3,4], queries = [[1,0],[-3,1],[-4,0],[2,3]]
Output: [8,6,2,4]
Explanation:
At the beginning, the array is [1,2,3,4].
After adding 1 to A[0], the array is [2,2,3,4], and the sum of even values is 2 + 2 + 4 = 8.
After adding -3 to A[1], the array is [2,-1,3,4], and the sum of even values is 2 + 4 = 6.
After adding -4 to A[0], the array is [-2,-1,3,4], and the sum of even values is -2 + 4 = 2.
After adding 2 to A[3], the array is [-2,-1,3,6], and the sum of even values is -2 + 6 = 4.


Note:

1 <= A.length <= 10000
-10000 <= A[i] <= 10000
1 <= queries.length <= 10000
-10000 <= queries[i][0] <= 10000
0 <= queries[i][1] < A.length
 */

package main

import "fmt"

func evenSum(A []int) int {
	sum := 0
	for _, v := range A {
		if v % 2 == 0 {
			sum += v
		}
	}
	return sum
}

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	res := make([]int, 0)
	for i := 0; i < len(queries); i++ {
		fmt.Println(i, queries[i][1])
		A[queries[i][1]] += queries[i][0]
		res = append(res, evenSum(A))
	}
	return res
}

func main() {
	fmt.Println(sumEvenAfterQueries([]int{1,2,3,4}, [][]int{[]int{1,0}, []int{-3,1}, []int{-4,0}, []int{2, 3}}))
}