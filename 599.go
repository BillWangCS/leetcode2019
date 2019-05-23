/*
Suppose Andy and Doris want to choose a restaurant for dinner,
and they both have a list of favorite restaurants represented by strings.

You need to help them find out their common interest with the least list index sum.
If there is a choice tie between answers, output all of them with no order requirement.
You could assume there always exists an answer.

Example 1:
Input:
["Shogun", "Tapioca Express", "Burger King", "KFC"]
["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"]
Output: ["Shogun"]
Explanation: The only restaurant they both like is "Shogun".
Example 2:
Input:
["Shogun", "Tapioca Express", "Burger King", "KFC"]
["KFC", "Shogun", "Burger King"]
Output: ["Shogun"]
Explanation: The restaurant they both like and have the least index sum is "Shogun" with index sum 1 (0+1).
Note:
The length of both lists will be in the range of [1, 1000].
The length of strings in both lists will be in the range of [1, 30].
The index is starting from 0 to the list length minus 1.
No duplicates in both lists.
 */

package main

import "fmt"

func findRestaurant(list1 []string, list2 []string) []string {
	dict1 := make(map[string]int)
	for index, v := range list1 {
		dict1[v] = index
	}
	res := make([]string, 0)
	min := 10000
	for i, v := range list2 {
		if _, ok := dict1[v]; ok {
			sum := i + dict1[v]
			if min >= sum - 2000 {
				min = sum - 2000
				dict1[v] = min
			}
		}
	}
	for k, v := range dict1 {
		if v == min {
			res = append(res, k)
		}
	}
	return res
}

func main() {
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}))
}