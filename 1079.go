/*
You have a set of tiles, where each tile has one letter tiles[i] printed on it.
Return the number of possible non-empty sequences of letters you can make.



Example 1:

Input: "AAB"
Output: 8
Explanation: The possible sequences are "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA".
Example 2:

Input: "AAABBC"
Output: 188


Note:

1 <= tiles.length <= 7
tiles consists of uppercase English letters.
 */

package main

import "fmt"

func dfs_1079(tiles string, path string, visitedIndex *map[int]bool, visitedRes *map[string]bool) {
	if len(path) == len(tiles) {
		(*visitedRes)[path] = true
		return
	}
	for i := 0; i < len(tiles); i++ {
		if !(*visitedIndex)[i] {
			path += string(tiles[i])
			(*visitedRes)[path] = true
			(*visitedIndex)[i] = true
			dfs_1079(tiles, path, visitedIndex, visitedRes)
			path = path[:len(path)-1]
			(*visitedIndex)[i] = false
		}
	}
}

func numTilePossibilities(tiles string) int {
	visitedIndex := make(map[int]bool)
	visitedRes := make(map[string]bool)
	dfs_1079(tiles, "", &visitedIndex, &visitedRes)
	return len(visitedRes)
}

func main() {
	fmt.Println(numTilePossibilities("AAABBC"))
}
