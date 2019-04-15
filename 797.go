/*
Given a directed, acyclic graph of N nodes.  Find all possible paths from node 0 to node N-1,
and return them in any order.

The graph is given as follows:  the nodes are 0, 1, ..., graph.length - 1.
graph[i] is a list of all nodes j for which the edge (i, j) exists.

Example:
Input: [[1,2], [3], [3], []]
Output: [[0,1,3],[0,2,3]]
Explanation: The graph looks like this:
0--->1
|    |
v    v
2--->3
There are two paths: 0 -> 1 -> 3 and 0 -> 2 -> 3.
Note:

The number of nodes in the graph will be in the range [2, 15].
You can print different paths in any order, but you should keep the order of nodes inside one path.
 */

package main

import "fmt"

func getPath(graph [][]int, path []int, res *[][]int, pos int) {
	if len(graph[pos]) == 0 {
		x := make([]int, 0)
		x = append(x, path...)
		*res = append(*res, x)
		return
	}
	for _, v := range graph[pos] {
		path = append(path, v)
		getPath(graph, path, res, v)
		path = path[:len(path)-1]
	}
}

func allPathsSourceTarget(graph [][]int) [][]int {
	res := make([][]int, 0)
	if len(graph) == 0 {
		return res
	}
	path := make([]int, 0)
	path = append(path, 0)
	for _, v := range graph[0] {
		path = append(path, v)
		getPath(graph, path, &res, v)
		path = path[:len(path)-1]
		fmt.Println(path)
	}
	return res
}

func main() {
	fmt.Println(allPathsSourceTarget([][]int{[]int{4, 3, 1}, []int{3, 2, 4}, []int{3}, []int{4}, []int{}}))
	//[[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]
}