/*
Suppose you have N integers from 1 to N.
We define a beautiful arrangement as an array that is constructed by these N numbers
successfully if one of the following is true for the ith position (1 <= i <= N) in this array:

The number at the ith position is divisible by i.
i is divisible by the number at the ith position.


Now given N, how many beautiful arrangements can you construct?

Example 1:

Input: 2
Output: 2
Explanation:

The first beautiful arrangement is [1, 2]:

Number at the 1st position (i=1) is 1, and 1 is divisible by i (i=1).

Number at the 2nd position (i=2) is 2, and 2 is divisible by i (i=2).

The second beautiful arrangement is [2, 1]:

Number at the 1st position (i=1) is 2, and 2 is divisible by i (i=1).

Number at the 2nd position (i=2) is 1, and i (i=2) is divisible by 1.


Note:

N is a positive integer and will not exceed 15.
 */

package main

import "fmt"

func dfs_526(N int, res *int, path *[]int, visited *map[int]bool) {
	if len(*path) == N {
		*res++
		return
	}
	for i := 1; i <= N; i++ {
		if !(*visited)[i] && (i % (len(*path)+1) == 0 || (len(*path)+1) % i == 0) {
			(*visited)[i] = true
			*path = append(*path, i)
			dfs_526(N, res, path, visited)
			(*visited)[i] = false
			*path = (*path)[:len(*path)-1]
		}
	}
}

func countArrangement(N int) int {
	res := 0
	visited := make(map[int]bool)
	path := make([]int, 0)
	dfs_526(N, &res, &path, &visited)
	return res
}

func main() {
	fmt.Println(countArrangement(3))
}