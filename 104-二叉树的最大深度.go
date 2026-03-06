package cdyleetcode

// *************************************************
// 递归法
// *************************************************
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    if root.Left == nil && root.Right == nil {
        return 1
    }

    return max(1+maxDepth(root.Left), 1+maxDepth(root.Right))

}

// *************************************************
// 递归法（深度优先），简化一个结束条件
// *************************************************
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    return max(1+maxDepth(root.Left), 1+maxDepth(root.Right))

}
