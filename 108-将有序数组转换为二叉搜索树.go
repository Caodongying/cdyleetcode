package cdyleetcode

// *************************************************
// 递归法，秒了
// *************************************************
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }

    mid := len(nums)/2

    leftTree := sortedArrayToBST(nums[:mid])
    rightTree := sortedArrayToBST(nums[mid+1:])

    node := TreeNode{
        Val: nums[mid],
        Left: leftTree,
        Right: rightTree,
    }

    return &node
}

// 递归法
