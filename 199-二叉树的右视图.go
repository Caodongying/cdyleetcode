package cdyleetcode

// *************************************************
// 利用层序遍历
// *************************************************
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    queue := []*TreeNode{root}
    result := []int{}

    for len(queue) > 0 {
        levelLength := len(queue)
        // 把队列里最后面的值加入result
        result = append(result, queue[len(queue)-1].Val)
        // 把下一层子节点加入队列尾端
        for levelLength > 0 {
            front := queue[0]
            queue = queue[1:]

            if front.Left != nil {
                queue = append(queue, front.Left)
            }

            if front.Right != nil {
                queue = append(queue, front.Right)
            }

            levelLength--
        }
    }

    return result
}

// 每一层的最右边的节点

// 用队列进行层序遍历
// 记录每一层的最后一个节点

// *************************************************
//
// *************************************************

// *************************************************
//
// *************************************************

// *************************************************
//
// *************************************************
