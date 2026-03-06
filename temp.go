package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }

    leftChildHeight := height(root.Left)
    if leftChildHeight != 0 {
        leftChildHeight++
    }

    rightChildHeight := height(root.Right)
    if rightChildHeight != 0 {
        rightChildHeight++
    }

    return max(diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right), leftChildHeight+rightChildHeight)
}



// 计算树的高度，用边数表示
func height(root *TreeNode) int {
    if root == nil || (root.Left == nil && root.Right == nil){
        return 0
    } else{
        return max(height(root.Left), height(root.Right)) + 1
    }
}

func main() {
	// [1,2,3,4,5]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	fmt.Println(diameterOfBinaryTree(root))
}