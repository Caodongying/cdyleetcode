package cdyleetcode

root =
[-10,9,20,null,null,15,7]

// *************************************************
// ❌❌自己写的，最后的路径有不符合逻辑的地方
// *************************************************

func maxPathSum(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := maxPathSum(root.Left)
    right := maxPathSum(root.Right)

    var result int

    if left <= 0 {
        if right > 0 {
            result = max(root.Val+right, right)
        } else {
            result = max(left, root.Val, right)
        }
    } else if left > 0 {
        if right > 0 {
            result = max(left+right, root.Val+left+right)
        } else {
            result = max(root.Val, root.Val+left)
        }
    }

    return result

}

// 路径和分为多种情况，然后取最大值
// 0. 纯root
// 1. root.Left的maxPathSum
// 2. root.Left的maxPathSum，经过root
// 3. root.Left的maxPathSum，经过root，然后到root.Right的maxPathSum
// 4-6. 对于root.Right也有以上三种情况

// *************************************************
// ❌❌ 没有包括所有情况
// *************************************************

func maxPathSum(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := maxPathSum(root.Left)
    right := maxPathSum(root.Right)

    var result int

    if left <= 0 {
        if right > 0 {
            // 一定不包括左边，一定包括右边
            result = max(root.Val+right, right)
        } else {
            // 一定不同时包括左右两边，矮子里挑将军
            result = max(left, root.Val, right)
        }
    } else if left > 0 {
        if right > 0 {
            // 如果同时包括左右两边，必须也包括root
            // 单独考虑左边或者右边，是防止root.Val为负数
            result = max(root.Val+left+right, left, right)
        } else {
            // 一定不包括右边
            result = max(root.Val, root.Val+left)
        }
    }

    return result

}

// 路径和分为多种情况，然后取最大值
// 0. 纯root
// 1. root.Left的maxPathSum
// 2. root.Left的maxPathSum，经过root
// 3. root.Left的maxPathSum，经过root，然后到root.Right的maxPathSum
// 4-6. 对于root.Right也有以上三种情况

// *************************************************
// ❌❌ 不可以用0作为root为nil的返回值。而且，当调用 maxPathSum(root.Left) 时，得到的是左子树中任意路径的最大和，但这个路径可能不经过 root.Left 节点本身。
// *************************************************

func maxPathSum(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := maxPathSum(root.Left)
    right := maxPathSum(root.Right)

    var result int

    // if left <= 0 {
    //     if right > 0 {
    //         // 一定不包括左边，一定包括右边
    //         result = max(root.Val+right, right)
    //     } else {
    //         // 一定不同时包括左右两边，矮子里挑将军
    //         result = max(left, root.Val, right)
    //     }
    // } else if left > 0 {
    //     if right > 0 {
    //         // 如果同时包括左右两边，必须也包括root
    //         // 单独考虑左边或者右边，是防止root.Val为负数
    //         result = max(root.Val+left+right, left, right)
    //     } else {
    //         // 一定不包括右边
    //         result = max(root.Val, root.Val+left)
    //     }
    // }

    result = max(left, left+root.Val, left+root.Val+right, root.Val, root.Val+right, right)
    return result

}

// 路径和分为多种情况，然后取最大值
// 0. 纯root
// 1. root.Left的maxPathSum
// 2. root.Left的maxPathSum，经过root
// 3. root.Left的maxPathSum，经过root，然后到root.Right的maxPathSum
// 4-5. 对于root.Right也有以上两种种情况（不包括3，重复了）

// *************************************************
// ❌❌ 依然是大小比较出错，其实是尝试在一次递归中同时解决“局部最优”和“全局最优”，结果导致了逻辑混乱。
// *************************************************

func maxPathSum(root *TreeNode) int {
    var result int

    if root.Left == nil && root.Right == nil {
        return root.Val
    }

    if root.Left != nil { // 不可以省略!=nil
        left := maxPathSum(root.Left)
        if root.Val >= 0 {
            result = max(left+root.Val, root.Val)
        } else {
            result = max(left, root.Val)
        }
    }

    if root.Right != nil {
        right := maxPathSum(root.Right)
        if root.Val >= 0 {
            result = max(result, result+right)
        } else {
            result = max(root.Val, right)
        }
    }


    // if left <= 0 {
    //     if right > 0 {
    //         // 一定不包括左边，一定包括右边
    //         result = max(root.Val+right, right)
    //     } else {
    //         // 一定不同时包括左右两边，矮子里挑将军
    //         result = max(left, root.Val, right)
    //     }
    // } else if left > 0 {
    //     if right > 0 {
    //         // 如果同时包括左右两边，必须也包括root
    //         // 单独考虑左边或者右边，是防止root.Val为负数
    //         result = max(root.Val+left+right, left, right)
    //     } else {
    //         // 一定不包括右边
    //         result = max(root.Val, root.Val+left)
    //     }
    // }

    return result

}

// *************************************************
// ❌❌ 错误理解为“必须要经过根节点"
// *************************************************


func maxPathSum(root *TreeNode) int {

    var dfs func(root *TreeNode) int // 以root为根节点的合法路径的最大路径和
    dfs = func(root *TreeNode) int { // 这个int不能省略
        if root == nil {
            return 0
        }

        left := dfs(root.Left)
        right := dfs(root.Right)

        result := max(root.Val, root.Val+max(left, right))
        return result
    }

    return dfs(root)

}


// 华南溜达虎的讲解：
// 枚举以每一个节点为根节点的合法路径的最大路径和（注意是为根节点，而不是经过该节点的最大路径和）

// *************************************************
// 看了华南溜达虎的讲解
// *************************************************

import "math"
func maxPathSum(root *TreeNode) int {
    result := math.MinInt // 不是minInt

    var dfs func(root *TreeNode) int// 以root为根节点的合法路径的最大路径和
    dfs = func(root *TreeNode) int{
        if root == nil {
            return 0
        }

        left := max(0, dfs(root.Left)) // 因为路径和可能是负数
        right := max(0, dfs(root.Right))

        result = max(result, left+root.Val+right)

        currentMaxSum := max(root.Val, root.Val+max(left, right)) // 这是以当前root为根节点的最大路径和
        return currentMaxSum
    }

    dfs(root)
    return result
}


// 华南溜达虎的讲解：
// 枚举以每一个节点为根节点的合法路径的最大路径和（注意是为根节点，而不是经过该节点的最大路径和）