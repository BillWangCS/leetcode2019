/*
There are 2N people a company is planning to interview. The cost of flying the i-th person to city A is costs[i][0],
and the cost of flying the i-th person to city B is costs[i][1].

Return the minimum cost to fly every person to a city such that exactly N people arrive in each city.



Example 1:

Input: [[10,20],[30,200],[400,50],[30,20]]
Output: 110
Explanation:
The first person goes to city A for a cost of 10.
The second person goes to city A for a cost of 30.
The third person goes to city B for a cost of 50.
The fourth person goes to city B for a cost of 20.

The total minimum cost is 10 + 30 + 50 + 20 = 110 to have half the people interviewing in each city.


Note:

1 <= costs.length <= 100
It is guaranteed that costs.length is even.
1 <= costs[i][0], costs[i][1] <= 1000
 */

package main

import (
	"fmt"
	"sort"
)

func twoCitySchedCost(costs [][]int) int {
	res := 0
	toA := make([]int, 0)
	toB := make([]int, 0)
	for i := 0; i < len(costs); i++ {
		if costs[i][0] > costs[i][1] {
			toB = append(toB, costs[i][0]-costs[i][1])
			res += costs[i][1]
		} else {
			toA = append(toA, costs[i][1]-costs[i][0])
			res += costs[i][0]
		}
	}
	fmt.Println(res)
	fmt.Println(toA, toB)
	la := len(toA)
	lb := len(toB)
	if la > lb {
		sort.Ints(toA)
		cur := 0
		for i := la - 1; i >= len(costs)/2; i-- {
			res += toA[cur]
			cur++
		}
	} else if lb > la {
		sort.Ints(toB)
		cur := 0
		for i := lb - 1; i >= len(costs)/2; i-- {
			res += toB[cur]
			cur++
		}
	}
	return res
}

func main() {
	fmt.Println(twoCitySchedCost([][]int{[]int{259, 770}, []int{926, 667}, []int{184, 139}, []int{448, 54}, []int{840, 118}, []int{577, 496}}))
}

/*
[[259,770],[448,54],[926,667],[184,139],[840,118],[577,469]]
1859
 */
