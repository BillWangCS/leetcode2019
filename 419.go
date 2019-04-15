/*
Given an 2D board, count how many battleships are in it.
The battleships are represented with 'X's, empty slots are represented with '.'s. You may assume the following rules:
You receive a valid board, made of only battleships or empty slots.
Battleships can only be placed horizontally or vertically. In other words,
they can only be made of the shape 1xN (1 row, N columns) or Nx1 (N rows, 1 column), where N can be of any size.
At least one horizontal or vertical cell separates between two battleships - there are no adjacent battleships.
Example:
X..X
...X
...X
In the above board there are 2 battleships.
Invalid Example:
...X
XXXX
...X
This is an invalid board that you will not receive - as battleships will always have a cell separating between them.
Follow up:
Could you do it in one-pass, using only O(1) extra memory and without modifying the value of the board?
 */

package main

import "fmt"

func countBattleships(board [][]byte) int {
	res := 0
	r := len(board)
	c := len(board[0])
	if r == 0 {
		return 0
	}
	visit := make([][]bool, 0)
	visitR := make([]bool, 0)
	for j := 0; j < c; j++ {
		visitR = append(visitR, false)
	}
	for i := 0; i < r; i++ {
		x := make([]bool, 0)
		x = append(x, visitR...)
		visit = append(visit, x)
	}
	for i := 0; i < r; i++ {
		for j := 0; j< c; j++ {
			if visit[i][j] == true {
				continue
			}
			if board[i][j] == 'X' {
				for x := i; x < r; x++ {
					if board[x][j] == 'X' {
						visit[x][j] = true
					} else {
						break
					}
				}
				for y := j; y < c; y++ {
					if board[i][y] == 'X' {
						visit[i][y] = true
					} else {
						break
					}
				}
				res++
			}
		}
	}
	return res
}

func main() {
	fmt.Println(countBattleships([][]byte{[]byte{'.', 'X', '.', '.', 'X'},
										  []byte{'.', 'X', '.', '.', 'X'},
										  []byte{'.', '.', '.', '.', 'X'},
	                                      []byte{'X', '.', 'X', 'X', '.'},
										  []byte{'X', '.', '.', '.', 'X'}}))
}