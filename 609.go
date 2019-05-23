/*
Given a list of directory info including directory path,
and all the files with contents in this directory,
you need to find out all the groups of duplicate files in the file system in terms of their paths.

A group of duplicate files consists of at least two files that have exactly the same content.

A single directory info string in the input list has the following format:

"root/d1/d2/.../dm f1.txt(f1_content) f2.txt(f2_content) ... fn.txt(fn_content)"

It means there are n files (f1.txt, f2.txt ... fn.txt with content f1_content, f2_content ... fn_content, respectively)
in directory root/d1/d2/.../dm. Note that n >= 1 and m >= 0. If m = 0, it means the directory is just the root directory.

The output is a list of group of duplicate file paths.
For each group, it contains all the file paths of the files that have the same content.
A file path is a string that has the following format:

"directory_path/file_name.txt"

Example 1:

Input:
["root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"]
Output:
[["root/a/2.txt","root/c/d/4.txt","root/4.txt"],["root/a/1.txt","root/c/3.txt"]]


Note:

No order is required for the final output.
You may assume the directory name, file name and file content only has letters and digits, and the length of file content is in the range of [1,50].
The number of files given is in the range of [1,20000].
You may assume no files or directories share the same name in the same directory.
You may assume each given directory info represents a unique directory. Directory path and file info are separated by a single blank space.


Follow-up beyond contest:
Imagine you are given a real file system, how will you search files? DFS or BFS?
If the file content is very large (GB level), how will you modify your solution?
If you can only read the file by 1kb each time, how will you modify your solution?
What is the time complexity of your modified solution? What is the most time-consuming part and memory consuming part of it? How to optimize?
How to make sure the duplicated files you find are not false positive?
 */

package main

import (
	"fmt"
	"strings"
)

func findDuplicate(paths []string) [][]string {
	dict := make(map[string][]string)
	res := make([][]string, 0)
	if len(paths) == 0 {
		return res
	}
	for _, v := range paths {
		path := strings.Split(v, " ")
		for i := 1; i < len(path); i++ {
			nameContent := strings.Split(path[i], "(")
			value := path[0] + "/" + nameContent[0]
			if _, ok := dict[nameContent[1]]; !ok {
				tmp := make([]string, 0)
				tmp = append(tmp, value)
				dict[nameContent[1]] = tmp
			} else {
				dict[nameContent[1]] = append(dict[nameContent[1]], value)
			}
		}
	}
	for _, v := range dict {
		if len(v) > 1 {
			res = append(res, v)
		}
	}
	return res
}

/*
func findDuplicate(paths []string) [][]string {
	dict := make(map[string]string)
	res := make([][]string, 0)
	if len(paths) == 0 {
		return res
	}
	for _, v := range paths {
		path := strings.Split(v, " ")
		for i := 1; i < len(path); i++ {
			nameContent := strings.Split(path[i], "(")
			key := path[0] + "/" + nameContent[0]
			dict[key] = nameContent[1]
		}
	}
	visited := make(map[string]bool)
	for k1, v1 := range dict {
		group := make([]string, 0)
		if !visited[k1] {
			visited[k1] = true
			group = append(group, k1)
		}
		for k2, v2 := range dict {
			if !visited[k2] {
				if v1 == v2 {
					visited[k2] = true
					group = append(group, k2)
				}
			}
		}
		if len(group) > 0 {
			res = append(res, group)
		}
	}
	return res
}
*/

func main() {
	fmt.Println(findDuplicate([]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"}))
}