/*
A binary watch has 4 LEDs on the top which represent the hours (0-11), and the 6 LEDs on the bottom represent the minutes (0-59).

Each LED represents a zero or one, with the least significant bit on the right.


For example, the above binary watch reads "3:25".

Given a non-negative integer n which represents the number of LEDs that are currently on, return all possible times
the watch could represent.

Example:

Input: n = 1
Return: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]
Note:
The order of output does not matter.
The hour must not contain a leading zero, for example "01:00" is not valid, it should be "1:00".
The minute must be consist of two digits and may contain a leading zero, for example "10:2" is not valid, it should be "10:02".
 */

package main

import (
	"fmt"
	"strconv"
)

func dfs_401(n, l, number, pos int, res *[]int, visited *map[int]bool) {
	fmt.Println(n, pos, l)
	if n == 0 {
		*res = append(*res, number)
		return
	}
	for i := pos; i < l; i++ {
		if !(*visited)[i] {
			(*visited)[i] = true
			dfs_401(n-1, l, number|1<<uint(i), i+1, res, visited)
			(*visited)[i] = false
		}
	}
}

func readBinaryWatch(num int) []string {
	res := make([]string, 0)
	for i := 0; i < 4 && i <= num; i++ {
		if num - i > 6 {
			continue
		}
		res1 := make([]int, 0)
		res2 := make([]int, 0)
		visited := make(map[int]bool)
		dfs_401(i, 4, 0, 0, &res1, &visited)
		dfs_401(num-i, 6, 0, 0, &res2, &visited)
		for x := 0; x < len(res1); x++ {
			if res1[x] > 11 {
				continue
			}
			for y := 0; y < len(res2); y++ {
				if res2[y] > 59 {
					continue
				}
				if len(strconv.Itoa(res2[y])) == 1 {
					res = append(res, strconv.Itoa(res1[x]) + ":0" + strconv.Itoa(res2[y]))
				} else {
					res = append(res, strconv.Itoa(res1[x])+":"+strconv.Itoa(res2[y]))
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(readBinaryWatch(1))
}