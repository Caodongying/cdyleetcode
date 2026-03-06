package cdyleetcode

// *************************************************
// 递归法，一开始写有一些语法错误
// *************************************************
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    if root.Left == nil && root.Right == nil {
        return []int{root.Val} // 可以是[root.Val]吗 - 不可以
    }

    result := make([]int, 0)

    left := inorderTraversal(root.Left)
    right := inorderTraversal(root.Right)

    result = append(left, root.Val)
    result = append(result, right...)
    // 这是错的 result = append(result, right)
    // 因为append的

    return result
}

// *************************************************
// 看了题解后，优化了的递归法，用闭包
// 我一开始的解法不好的原因：
// 		性能较差 - 每次递归调用都会创建新切片，append操作可能触发多次内存分配和拷贝
//		空间复杂度高 - 会产生大量临时切片
// *************************************************

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    result := []int{}

    var traversal func(node *TreeNode)
    traversal = func(node *TreeNode) {
        if node == nil {
            return
        }

       traversal(node.Left)
       result = append(result, node.Val)
       traversal(node.Right)
    }

    traversal(root)

    return result
}

// *************************************************
// ❌❌迭代法实现（用栈） 这里我写错了——
// 逻辑混乱：试图用递归函数模拟栈操作，但中序遍历的栈实现应该是迭代的，而不是递归+栈混合
// 左子树的压栈逻辑不完整（只处理了左子树的左节点，没有处理完整的左链）
// 出栈后没有继续处理当前节点的右子树
// *************************************************
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    result := []int{}

    // 定义栈
    stack := []*TreeNode{}

    var stackProcess func(root *TreeNode)
    stackProcess = func(root *TreeNode) {
        // 左子树压栈处理
        for {
            left := root.Left
            if left == nil {
                break
            }
            stack = append(stack, left)
            root = root.Left
        }

        // 左子树出栈处理
        for {
            // 栈顶弹出元素
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            result = append(result, top.Val)
            // 压入右节点
            if top.Right != nil {
                stack = append(stack, top.Right)
            }
        }
    }

    stackProcess(root.Left)
    result = append(result, originalRoot.Val)
    stackProcess(root.Right)

    return result

}

// 定义一个栈（后进先出）- 用切片实现
// 先对左子树进行压栈、出栈处理
// 然后对右子树压栈、出栈
// 中序遍历就是左-中-右

// *************************************************
// Deepseek给改的，其实迭代自己没写出来
// *************************************************
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    result := []int{}

    // 定义栈
    stack := []*TreeNode{}
    current := root

    for current != nil || len(stack) > 0 {
        // 左节点全部压栈
        for current != nil {
            stack = append(stack, current)
            current = current.Left
        }

        // 对左节点进行出栈处理
        current = stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        result = append(result, current.Val)

        current = current.Right
    }

    return result

}

// 定义一个栈（后进先出）- 用切片实现
// 先对左子树进行压栈、出栈处理
// 然后对右子树压栈、出栈
// 中序遍历就是左-中-右