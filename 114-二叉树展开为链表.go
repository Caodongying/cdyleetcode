package cdyleetcode

// *************************************************
// 直接看的官方题解，用Morris遍历的变种，实现O(1)的空间复杂度
// *************************************************

// 不断把右子树放到左子树的右下角，然后把这个链接后的右子树直接变成原先的左子树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    current := root

    for current != nil {
        // 找到左子树的最右子节点
        rightMost := current.Left
        if rightMost == nil {
            current = current.Right
            continue
        }
        for rightMost.Right != nil {
            rightMost = rightMost.Right
        }
        // 把原先的右子树链接到最右子节点
        rightMost.Right = current.Right
        // 把左子树变成右子树
        current.Right = current.Left
        current.Left = nil

        current = current.Right
    }

}

// *************************************************
// 小小优化，其实就是改变了一下if的判断顺序
// *************************************************
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    current := root

    for current != nil {
        // 找到左子树的最右子节点
        rightMost := current.Left
        if rightMost != nil {
            for rightMost.Right != nil {
                rightMost = rightMost.Right
            }
            // 把原先的右子树链接到最右子节点
            rightMost.Right = current.Right
            // 把左子树变成右子树
            current.Right = current.Left
            current.Left = nil
        }

        current = current.Right
    }

}

// 不断把右子树放到左子树的右下角，然后把这个链接后的右子树直接变成原先的左子树
