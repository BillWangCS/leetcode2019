/*
We have a list of points on the plane.  Find the K closest points to the origin (0, 0).

(Here, the distance between two points on a plane is the Euclidean distance.)

You may return the answer in any order.  The answer is guaranteed to be unique (except for the order that it is in.)



Example 1:

Input: points = [[1,3],[-2,2]], K = 1
Output: [[-2,2]]
Explanation:
The distance between (1, 3) and the origin is sqrt(10).
The distance between (-2, 2) and the origin is sqrt(8).
Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
We only want the closest K = 1 points from the origin, so the answer is just [[-2,2]].
Example 2:

Input: points = [[3,3],[5,-1],[-2,4]], K = 2
Output: [[3,3],[-2,4]]
(The answer [[-2,4],[3,3]] would also be accepted.)


Note:

1 <= K <= points.length <= 10000
-10000 < points[i][0] < 10000
-10000 < points[i][1] < 10000
 */


package main

import "fmt"

type pointInfo struct {
	point []int
	distance int
}

func quickK(A *[]pointInfo, start, end int, K int) {
	if start >= end {
		return
	}
	left, right := start, end
	pivot := (*A)[(end - start) / 2 + start]
	for left <= right {
		for left <= right && (*A)[left].distance < pivot.distance {
			left++
		}
		for left <= right && (*A)[right].distance > pivot.distance {
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

func kClosest(points [][]int, K int) [][]int {
	res := make([][]int, 0)
	pointsInfo := make([]pointInfo, 0)
	for _, p := range points {
		pointsInfo = append(pointsInfo, pointInfo{point: p, distance: p[0]*p[0] + p[1]*p[1]})
	}
	quickK(&pointsInfo, 0, len(pointsInfo) - 1, K)
	for i := 0; i < K; i++ {
		res = append(res, pointsInfo[i].point)
	}
	return res
}

func main() {
	fmt.Println(kClosest([][]int{[]int{3, 3}, []int{5, -1}, []int{-2, 4}}, 2))
}