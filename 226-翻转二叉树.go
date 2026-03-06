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
func invertTree(root *TreeNode) *TreeNode {

    invertNode(root)

    return root
}

func invertNode(node *TreeNode) *TreeNode {
    if node == nil {
        return nil
    }

    left := node.Left
    right := node.Right

    node.Left = right
    node.Right = left

    invertNode(node.Left)
    invertNode(node.Right)

    return node
}

// *************************************************
// 优化版，不用单独另开一个函数
// *************************************************
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    left := root.Left
    right := root.Right

    root.Left = right
    root.Right = left

    invertTree(root.Left)
    invertTree(root.Right)

    return root
}


// *************************************************
//
// *************************************************

// *************************************************
//
// *************************************************
