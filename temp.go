package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var findPath func(root, target *TreeNode, result *[]*TreeNode) bool
    findPath = func(root, target *TreeNode, result *[]*TreeNode) bool {
        if root == nil {
            return false
        }
        *result = append(*result, root)

        if root.Val == target.Val {
            fmt.Print("找到了，此时的result是")
            for _, node := range *result {
                fmt.Print(node.Val, " ")
            }
            fmt.Println()
            return true
        }

        if root.Left == nil && root.Right == nil {
            *result = (*result)[:len(*result)-1]
            return false
        }

        if findPath(root.Left, target, result) {
            return true
        }
        rightExist := findPath(root.Right, target, result)
        if !rightExist {
            *result = (*result)[:len(*result)-1]
        }
        return rightExist
    }

    pathP := []*TreeNode{}
    pathQ := []*TreeNode{}

    findPath(root, p, &pathP)
    findPath(root, q, &pathQ)

    length := min(len(pathP), len(pathQ))

    var i int
    for i=0; i < length && pathP[i].Val == pathQ[i].Val; i++ {}

    // return i-1
    return pathP[i-1] // 返回值要正确
}

// 理想状况：深度优先遍历dfs
// 确定找到p时的路径，存在切片里
// 确定找到q时的路径，存在切片里
// 这两个路径必定以root.Val开头
// 找到这两个切片从下标0开始的，最后一个相同的值
// 这个值对应的节点就是最终要的结果


func main() {
    // [3,5,1,6,2,0,8,null,null,7,4]
    root := &TreeNode{Val: 3}
    root.Left = &TreeNode{Val: 5}
    root.Right = &TreeNode{Val: 1}
    root.Left.Left = &TreeNode{Val: 6}
    root.Left.Right = &TreeNode{Val: 2}
    root.Right.Left = &TreeNode{Val: 0}
    root.Right.Right = &TreeNode{Val: 8}
    root.Left.Right.Left = &TreeNode{Val: 7}
    root.Left.Right.Right = &TreeNode{Val: 4}

    p := &TreeNode{Val: 5}
    q := &TreeNode{Val: 1}

    result := lowestCommonAncestor(root, p, q)

    fmt.Println(result.Val)
}
