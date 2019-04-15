/*
Return any binary tree that matches the given preorder and postorder traversals.

Values in the traversals pre and post are distinct positive integers.



Example 1:

Input: pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
Output: [1,2,3,4,5,6,7]


Note:

1 <= pre.length == post.length <= 30
pre[] and post[] are both permutations of 1, 2, ..., pre.length.
It is guaranteed an answer exists. If there exists multiple answers, you can return any of them.
 */

package main

func treeConstructor(root *TreeNode, pre, post []int, direct int) {
	if len(pre) == 0 {
		return
	}
	prePos := 0
	postPos := 0
	if len(pre) > 1 {
		for i := 0; i < len(pre); i++ {
			if pre[i] == post[len(post)-2] {
				prePos = i
				break
			}
		}
		for i := 0; i < len(post); i++ {
			if post[i] == pre[1] {
				postPos = i
				break
			}
		}
	}
	if direct == 0 {
		root.Left = &TreeNode{Val: pre[0]}
		if len(pre) == 1 {
			return
		}
		if postPos + 1 == len(post)-1 {
			postPos = -1
		}
		treeConstructor(root.Left, pre[1:prePos], post[:postPos + 1], 0)
		treeConstructor(root.Left, pre[prePos:], post[postPos+1:len(post) - 1], 1)
	} else {
		root.Right = &TreeNode{Val: pre[0]}
		if len(pre) == 1 {
			return
		}
		if postPos + 1 == len(post)-1 {
			postPos = -1
		}
		treeConstructor(root.Right, pre[1:prePos], post[:postPos + 1], 0)
		treeConstructor(root.Right, pre[prePos:], post[postPos+1:len(post) - 1], 1)
	}
}

func constructFromPrePost(pre []int, post []int) *TreeNode {
	root := TreeNode{Val: pre[0]}
	if len(pre) == 1 {
		return &root
	}
	prePos := 0
	postPos := 0
	if len(pre) > 1 {
		for i := 0; i < len(pre); i++ {
			if pre[i] == post[len(post)-2] {
				prePos = i
				break
			}
		}
		for i := 0; i < len(post); i++ {
			if post[i] == pre[1] {
				postPos = i
				break
			}
		}
		if postPos + 1 == len(post)-1 {
			postPos = -1
		}
		treeConstructor(&root, pre[1:prePos], post[:postPos+1], 0)
		treeConstructor(&root, pre[prePos:], post[postPos+1:len(post)-1], 1)
	}
	return &root
}

/*
2 0
[2 4 5] [4 5 2]
 */