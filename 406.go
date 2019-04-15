/*
Suppose you have a random list of people standing in a queue.
Each person is described by a pair of integers (h, k),
where h is the height of the person and k is the number of people in
front of this person who have a height greater than or equal to h. Write an algorithm to reconstruct the queue.

Note:
The number of people is less than 1,100.


Example

Input:
[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]

Output:
[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
 */

package main

import (
	"fmt"
	"sort"
)

func insertArr(res *[][]int, heights []int, n int) {
	for i := 0; i < len(heights); i++ {
		cnt := 0
		fmt.Println(*res, heights[i], n, cnt)
		for j := 0; j < len(*res); j++ {
			if (*res)[j][0] >= heights[i] {
				cnt++
			}
			if cnt == n {
				tmp := []int{heights[i], n}
				res1 := (*res)[:j+1]
				res2 := make([][]int, 0)
				res2 = append(res2, (*res)[j+1:]...)
				res1 = append(res1, tmp)
				res1 = append(res1, res2...)
				*res = res1
				break
			}
		}
	}
}

func reverseArr(A *[]int) {
	i := 0
	j := len(*A) - 1
	for i <= j {
		(*A)[i], (*A)[j] = (*A)[j], (*A)[i]
		i++
		j--
	}
}

func reconstructQueue(people [][]int) [][]int {
	res := make([][]int, 0)
	heightDict := make(map[int][]int)
	for _, person := range people {
		heightDict[person[1]] = append(heightDict[person[1]], person[0])
	}
	for i, v := range heightDict {
		sort.Ints(v)
		reverseArr(&v)
		heightDict[i] = v
	}
	for i := len(heightDict[0]) - 1; i >= 0; i-- {
		tmp := []int{heightDict[0][i], 0}
		res = append(res, tmp)
	}
	for i := 1; i < len(people); i++ {
		if len(heightDict[i]) == 0 {
			continue
		}
		insertArr(&res, heightDict[i], i)
	}
	return res
}

func main() {
	fmt.Println(reconstructQueue([][]int{[]int{7,0}, []int{4,4}, []int{7,1}, []int{5,0}, []int{6,1}, []int{5,2}}))
}