/*
You have a list of points in the plane. Return the area of the largest triangle that can be formed by any 3 of the points.

Example:
Input: points = [[0,0],[0,1],[1,0],[0,2],[2,0]]
Output: 2
Explanation:
The five points are show in the figure below. The red triangle is the largest.


Notes:

3 <= points.length <= 50.
No points will be duplicated.
 -50 <= points[i][j] <= 50.
Answers within 10^-6 of the true value will be accepted as correct.
	S=(1/2)*(x1y2+x2y3+x3y1-x1y3-x2y1-x3y2)
 */

package main

import "fmt"

func largestTriangleArea(points [][]int) float64 {
	max := 0
	l := len(points)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			for k := 0; k < l; k++ {
				s := points[i][0]*points[j][1] + points[j][0]*points[k][1] + points[k][0]*points[i][1] - points[i][0]*points[k][1] - points[j][0]*points[i][1] - points[k][0]*points[j][1]
				if s > max {
					fmt.Println(points[i], points[j], points[k])
					max = s
				}
			}
		}
	}
	return float64(max) / 2
}

func main() {
	fmt.Println(largestTriangleArea([][]int{[]int{0,0}, []int{0,1}, []int{1,0}, []int{0,2}, []int{2,0}}))
}