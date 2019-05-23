/*
Given n points in the plane that are all pairwise distinct,
a "boomerang" is a tuple of points (i, j, k) such that the distance between
i and j equals the distance between i and k (the order of the tuple matters).

Find the number of boomerangs.
You may assume that n will be at most 500 and coordinates of points are all in the range [-10000, 10000] (inclusive).

Example:

Input:
[[0,0],[1,0],[2,0]]

Output:
2

Explanation:
The two boomerangs are [[1,0],[0,0],[2,0]] and [[1,0],[2,0],[0,0]]
 */

package main

import "fmt"

/*Better Solution
func numberOfBoomerangs(points [][]int) int {
	result := 0
	for i, pointI := range points {
		hashMap := make(map[int]int, len(points))
		for j, pointJ := range points {
			if i == j {
				continue
			}
			distance := (pointJ[0]-pointI[0])*(pointJ[0]-pointI[0]) + (pointJ[1]-pointI[1])*(pointJ[1]-pointI[1])

			if n, exist := hashMap[distance]; exist {
				result += n * 2  //与现有的任意一个组合并交换顺序
				hashMap[distance]++
			} else {
				hashMap[distance] = 1
			}
		}
	}

	return result
}
 */

func isBoomerangs(points [][]int, pos1, pos2, pos3 int) bool {
	if (points[pos2][0] - points[pos1][0])*(points[pos2][0] - points[pos1][0]) + (points[pos2][1] - points[pos1][1])*(points[pos2][1] - points[pos1][1]) == (points[pos3][0] - points[pos1][0])*(points[pos3][0] - points[pos1][0]) + (points[pos3][1] - points[pos1][1])*(points[pos3][1] - points[pos1][1]) {
		//fmt.Println(points[pos1], points[pos2], points[pos3])
		return true
	}
	return false
}

func inPath(path []int, pos int) bool {
	for _, v := range path {
		if v == pos {
			return true
		}
	}
	return false
}

func countLeft(points [][]int, path *[]int, cnt *int) {
	if len(*path) == 3 {
		if isBoomerangs(points, (*path)[0], (*path)[1], (*path)[2]) {
			*cnt++
		}
		return
	}
	for i := 0; i < len(points); i++ {
		if !inPath(*path, i) {
			*path = append(*path, i)
		} else {
			continue
		}
		countLeft(points, path, cnt)
		*path = (*path)[:len(*path)-1]
	}
}

func numberOfBoomerangs(points [][]int) int {
	res := 0
	path := make([]int, 0)
	countLeft(points, &path, &res)
	return res
}

func main() {
	fmt.Println(numberOfBoomerangs([][]int{[]int{0,0}, []int{1,0}, []int{-1,0}, []int{0,1}, []int{0,-1}}))
}