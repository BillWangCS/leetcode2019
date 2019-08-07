/*In an infinite binary tree where every node has two children, the nodes are labelled in row order.

In the odd numbered rows (ie., the first, third, fifth,...), the labelling is left to right, while in the even numbered rows (second, fourth, sixth,...), the labelling is right to left.



Given the label of a node in this tree, return the labels in the path from the root of the tree to the node with that label.



Example 1:

Input: label = 14
Output: [1,3,4,14]
Example 2:

Input: label = 26
Output: [1,2,6,10,26]


Constraints:

1 <= label <= 10^6*/

package main

import "fmt"

func power(n uint32) int {
	return 1 << n
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

func pathInZigZagTree(label int) []int {
	res := make([]int, 0)
	layer := 0
	for i := 0; i < 32; i++ {
		if label / power(uint32(i)) == 0 {
			layer = i
			break
		}
	}
	for {
		if layer == 0 {
			break
		}
		res = append(res, label)
		if layer % 2 == 0 {
			diff := power(uint32(layer)) - 1 - label
			layer = layer - 1
			label = power(uint32(layer-1)) + diff/2
		} else {
			diff := label - power(uint32(layer-1))
			layer = layer - 1
			label = power(uint32(layer)) - 1 - diff/2
		}
	}
	reverseArr(&res)
	return res
}

func main() {
	fmt.Println(pathInZigZagTree(14))
}