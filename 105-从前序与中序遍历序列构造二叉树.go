package cdyleetcode

// *************************************************
// 直接看的题解，递归法
// *************************************************
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    rootVal := preorder[0]
    root := &TreeNode{
        Val: rootVal,
        Left: nil,
        Right: nil,
    }

    var leftLength int
    // var rightLength int
    for i, v := range inorder {
        if v == rootVal {
            leftLength = i
            // rightLength = len(inorder) - i - 1
            break
        }
    }

    root.Left = buildTree(preorder[1:1+leftLength], inorder[:leftLength])
    root.Right = buildTree(preorder[1+leftLength:], inorder[1+leftLength:])

    return root
}

// 递归法，在先序遍历确定根节点，然后在中序遍历中确定左右子树

// *************************************************
// gemini和灵神的优化，边界条件很麻烦
// *************************************************
func buildTree(preorder []int, inorder []int) *TreeNode {
    n := len(inorder)
    inorderMap := make(map[int]int, n)
    for i, v := range inorder {
        inorderMap[v] = i
    }

    var helper func(preStart, preEnd, inStart, inEnd int) *TreeNode
    helper = func(preStart, preEnd, inStart, inEnd int) *TreeNode {
        // 递归终止条件
        if preStart > preEnd {
            return nil
        }

        rootVal := preorder[preStart]
        root := &TreeNode{Val: rootVal}

        // mid 是当前根节点在 inorder 数组中的绝对索引
        mid := inorderMap[rootVal]

        // 左子树的节点数量
        leftSize := mid - inStart

        // --- 核心区：边界推导 ---

        // 左子树：
        // preorder: 根后面紧跟 leftSize 个元素 [preStart+1, preStart+leftSize]
        // inorder: 从当前区域左边界到 mid 前一个 [inStart, mid-1]
        root.Left = helper(preStart+1, preStart+leftSize, inStart, mid-1)

        // 右子树：
        // preorder: 左子树结束后的所有元素 [preStart+leftSize+1, preEnd]
        // inorder: 从 mid 后一个到当前区域右边界 [mid+1, inEnd]
        root.Right = helper(preStart+leftSize+1, preEnd, mid+1, inEnd)

        return root
    }

    return helper(0, n-1, 0, n-1)
}
