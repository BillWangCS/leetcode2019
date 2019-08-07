/*
You have N gardens, labelled 1 to N.  In each garden, you want to plant one of 4 types of flowers.

paths[i] = [x, y] describes the existence of a bidirectional path from garden x to garden y.

Also, there is no garden that has more than 3 paths coming into or leaving it.

Your task is to choose a flower type for each garden such that, for any two gardens connected by a path,
they have different types of flowers.

Return any such a choice as an array answer, where answer[i] is the type of flower planted in the (i+1)-th garden.
The flower types are denoted 1, 2, 3, or 4.  It is guaranteed an answer exists.



Example 1:

Input: N = 3, paths = [[1,2],[2,3],[3,1]]
Output: [1,2,3]
Example 2:

Input: N = 4, paths = [[1,2],[3,4]]
Output: [1,2,1,2]
Example 3:

Input: N = 4, paths = [[1,2],[2,3],[3,4],[4,1],[1,3],[2,4]]
Output: [1,2,3,4]


Note:

1 <= N <= 10000
0 <= paths.size <= 20000
No garden has 4 or more paths coming into or leaving it.
It is guaranteed an answer exists.
 */

package main

import "fmt"

type node struct {
	flower int
	notAllowed [4]int
}

func gardenNoAdj(N int, paths [][]int) []int {
	if N <= 4 {
		res := make([]int, N)
		for i := 1; i <= N; i++ {
			res[i-1] = i
		}
		return res
	}
	if len(paths) == 0 {
		res := make([]int, N)
		for i := 0; i < N; i++ {
			res[i] = 1
		}
		return res
	}
	nodes := make([]node, N+1)
	for i := 0 ; i < len(paths); i++ {
		pos := paths[i][0]
		for j := 0; j < 4; j++ {
			if nodes[pos].notAllowed[j] == 0 {
				nodes[pos].flower = j+1
				break
			}
		}
		nodes[paths[i][1]].notAllowed[nodes[paths[i][0]].flower-1] = 1
	}
	for i := 0 ; i < len(paths); i++ {
		if nodes[paths[i][1]].flower != 0 {
			nodes[paths[i][0]].notAllowed[nodes[paths[i][1]].flower-1] = 1
		}
	}
	res := make([]int, N+1)
	for i := 1; i <= N; i++ {
		for j := 0; j < 4; j++ {
			if nodes[i].notAllowed[j] == 0 {
				if j+1 != nodes[i].flower {
					for k := 0; k < len(paths); k++ {
						if paths[k][0] == i {
							nodes[paths[k][1]].notAllowed[j] = 1
						}
						if paths[k][1] == i {
							nodes[paths[k][0]].notAllowed[j] = 1
						}
					}
				}
				nodes[i].flower = j+1
				break
			}
		}
		res[i] = nodes[i].flower
	}
	return res[1:]
}

func main() {
	fmt.Println(gardenNoAdj(5, [][]int{[]int{4,5}, []int{4,2}, []int{4,3}, []int{2,5}, []int{1,2}, []int{1,5}}))
}

/*
2,2,3,1,2,1
[[4,2],[6,2],[6,3],[2,3],[5,3],[6,5],[5,4],[4,1]]
5
[[4,1],[4,2],[4,3],[2,5],[1,2],[1,5]]
[2,2,2,1,1]
5
[[4,5],[4,3],[2,3],[3,5],[2,4]]
[1,1,2,3,1]
 */