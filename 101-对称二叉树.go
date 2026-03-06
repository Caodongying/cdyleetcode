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
func isSymmetric(root *TreeNode) bool {
    return Symmetric(root.Left, root.Right)
}


func Symmetric(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }

    if p == nil || q == nil {
        return false
    }

    return p.Val == q.Val && Symmetric(p.Left, q.Right) && Symmetric(p.Right, q.Left)
}

// *************************************************
// ❌❌迭代法，使用stack==nil作为循环退出条件，错误，因为stack已经初始化过了
// *************************************************
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    // 栈定义
    stack := []*TreeNode{}

    if root.Left == nil && root.Right == nil {
        return true
    }

    if root.Left == nil || root.Right == nil {
        return false
    }

    if root.Left.Val != root.Right.Val {
        return false
    }

    stack  = append(stack, root.Left, root.Right)

    for {
        if stack == nil {
            break
        }
        // 连续弹出两个值
        p := stack[len(stack)-1]
        q := stack[len(stack)-2]
        stack = stack[:len(stack)-2]

        // 比较p.Left和q.Right
        // 一个为空
        if (p.Left == nil && q.Right != nil) || (p.Left != nil && q.Right == nil){
            return false
        }

        // 两个都不为空 （两个都为空就不处理）
        if (p.Left != nil && q.Right != nil) {
            if (p.Left.Val != q.Right.Val) {
                return false
            } else {
                stack = append(stack, p.Left, q.Right)
            }

        }

        // 比较p.Right和q.Left
        if (p.Right== nil && q.Left != nil) || (p.Right!=nil && q.Left == nil){
            return false
        }
        if  (p.Right != nil && q.Left != nil) {
            if (p.Right.Val != q.Left.Val) {
                return false
            } else {
                stack = append(stack, p.Right, q.Left)
            }
        }

    }

    return true
}

// 栈，用for循环
// 每一次，对root结点，同时压栈left和right节点
// 压栈后就进行两个节点的出栈处理
// 如果出栈的两个节点值相同，就对下面两组节点进行值比较和压入：（第一个结点的left、第二个节点的right），（第一个结点的right、第二个节点的left）
// 循环退出：
//      - 出栈的两个节点值不同，或者对应的子节点值不同(false)
// for循环以外，直接返回true

// *************************************************
// 迭代法，修改为len(stack)==0作为循环退出条件
// *************************************************
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    // 栈定义
    stack := []*TreeNode{}

    if root.Left == nil && root.Right == nil {
        return true
    }

    if root.Left == nil || root.Right == nil {
        return false
    }

    if root.Left.Val != root.Right.Val {
        return false
    }

    stack  = append(stack, root.Left, root.Right)

    for {
        if len(stack) == 0 {
            break
        }
        // 连续弹出两个值
        p := stack[len(stack)-1]
        q := stack[len(stack)-2]
        stack = stack[:len(stack)-2]

        // 比较p.Left和q.Right
        // 一个为空
        if (p.Left == nil && q.Right != nil) || (p.Left != nil && q.Right == nil){
            return false
        }

        // 两个都不为空 （两个都为空就不处理）
        if (p.Left != nil && q.Right != nil) {
            if (p.Left.Val != q.Right.Val) {
                return false
            } else {
                stack = append(stack, p.Left, q.Right)
            }

        }

        // 比较p.Right和q.Left
        if (p.Right== nil && q.Left != nil) || (p.Right!=nil && q.Left == nil){
            return false
        }
        if  (p.Right != nil && q.Left != nil) {
            if (p.Right.Val != q.Left.Val) {
                return false
            } else {
                stack = append(stack, p.Right, q.Left)
            }
        }

    }

    return true
}

// 栈，用for循环
// 每一次，对root结点，同时压栈left和right节点
// 压栈后就进行两个节点的出栈处理
// 如果出栈的两个节点值相同，就对下面两组节点进行值比较和压入：（第一个结点的left、第二个节点的right），（第一个结点的right、第二个节点的left）
// 循环退出：
//      - 出栈的两个节点值不同，或者对应的子节点值不同(false)
// for循环以外，直接返回true

// *************************************************
// Deepseek的优化，简化一些不必要的判断
// *************************************************

func isSymmetric(root *TreeNode) bool {
    // 栈定义
    stack := []*TreeNode{}
    stack = append(stack, root.Left, root.Right)

    for len(stack) > 0 {
        // 连续弹出两个值
        p := stack[len(stack)-1]
        q := stack[len(stack)-2]
        stack = stack[:len(stack)-2]

        if p == nil && q == nil {
            continue
        }

        if p== nil || q == nil || p.Val != q.Val {
            return false
        }

        stack = append(stack, p.Left, q.Right)
        stack = append(stack, p.Right, q.Left)
    }

    return true
}