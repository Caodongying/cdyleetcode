package cdyleetcode

// *************************************************
// ❌❌ 自己写的递归法，注意Golang里不可以有同名函数，没有重载这会儿事。
// 一开始写错了，因为忽略了最上层的边，好不容易改出来。
// *************************************************
func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }

    return max(diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right), height(root.Left)+height(root.Right))
}


// 计算树的高度，用边数表示
func height(root *TreeNode) int {
    if root == nil || (root.Left == nil && root.Right == nil){
        return 0
    } else{
        return max(height(root.Left), height(root.Right)) + 1
    }
}

// 三种情况取最大值：
// 1. 左子树的直径
// 2. 右子树的直径
// 3. 经过当前节点的路径（左子树高度 + 右子树高度）

// *************************************************
// 修改版，考虑到一开始忘记算入最上层的边了
// *************************************************

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }

    leftTreeHeight := height(root.Left)
    rightTreeHeight := height(root.Right)

    if root.Left != nil {
        leftTreeHeight++
    }

    if root.Right != nil {
        rightTreeHeight++
    }

    return max(diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right), leftTreeHeight+rightTreeHeight)

}


// 计算树的高度，用边数表示
func height(root *TreeNode) int {
    if root == nil || (root.Left == nil && root.Right == nil){
        return 0
    } else{
        return max(height(root.Left), height(root.Right)) + 1
    }
}


// *************************************************
// 优化版 - 使用节点数而不是边长，并且在计算递归时直接更新最大直径，避免重复计算高度
// *************************************************


// 优化点：你的代码片段是独立的两部分，没有组合成一个完整的函数。主要问题在于 diameterOfBinaryTree 函数
// 没有利用递归过程中计算出的高度，而是又去调用了一个独立的 height 函数，这导致了重复计算和逻辑分离。

func diameterOfBinaryTree(root *TreeNode) int {
    maxDiameter := 0

    var depth func(*TreeNode) int // 节点数

    depth = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        leftDepth := depth(node.Left)
        rightDepth := depth(node.Right)

        // 更新直径
        // 经过当前节点的路径长度（边数）=左子树深度+右子树深度
        if leftDepth + rightDepth > maxDiameter {
            maxDiameter = leftDepth + rightDepth
        }

        // 返回当前树的深度
        if leftDepth > rightDepth {
            return leftDepth + 1
        }

        return rightDepth + 1
    }

    depth(root)

    return maxDiameter

}
