package cdyleetcode

// *************************************************
// ❌❌自己看了点前缀和题解后写的，错了很多
// *************************************************
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {

    result := 0

    prefix := make(map[int]int)
    prefix[0] = 1 // 表示还没出发

    var dfs func(root *TreeNode)

    currentSum := 0
    dfs = func(root *TreeNode){
        if root == nil {
            return
        }
        currentSum += root.Val
        prefix[currentSum]++
        if prefix[currentSum-targetSum] > 0 {
            result++
        }

        if root.Left != nil {
            dfs(root.Left)
            prefix[currentSum]-- //  really
        }

        dfs(root.Right)
    }

    dfs(root)

    return result
}

// 暴力解法：深度优先遍历所有路径(不写了)
// 应该使用前缀和，利用一个map。假设这个map叫prefix
// prefix[i]为键，表示从根节点到当前节点的路径和;key值就是这个路径和的出现次数
// 向下遍历的时候，把和加入map；向上回溯的时候，要从map里把当前前缀和对应的key-1。因为不然的话，路径需要向上返回，比如满足条件的前缀和是要经过左节点的，但当前节点是右节点

// 每一步，都看一下prefix[i]-targetSum的值在map中是否大于0，因为如果大于0，就说明有路径

// *************************************************
// ❌❌看了正确版本，尝试更改dfs的参数，以及回溯时机。错误原因是在计算前缀路径时，把自己算进去了
// *************************************************

func pathSum(root *TreeNode, targetSum int) int {

    result := 0

    prefix := make(map[int]int)
    prefix[0] = 1 // 表示还没出发

    var dfs func(root *TreeNode, currentSum int)

    dfs = func(root *TreeNode, currentSum int){
        if root == nil {
            return
        }

        // 更新当前前缀和
        currentSum += root.Val
        prefix[currentSum]++

        result += prefix[currentSum-targetSum] // 包含了为0的情况


        dfs(root.Left, currentSum)

        if root.Left != nil {
            prefix[currentSum+root.Left.Val]--
        }

        dfs(root.Right, currentSum)

        if root.Right != nil {
            prefix[currentSum+root.Right.Val]--
        }
    }

    dfs(root, 0)

    return result
}

// *************************************************
// 居然对了
// *************************************************

func pathSum(root *TreeNode, targetSum int) int {

    result := 0

    prefix := make(map[int]int)
    prefix[0] = 1 // 表示还没出发

    var dfs func(root *TreeNode, currentSum int)

    dfs = func(root *TreeNode, currentSum int){
        if root == nil {
            return
        }

        // 更新当前前缀和
        currentSum += root.Val

        if count, exists := prefix[currentSum-targetSum]; exists{
            result += count
        }

        prefix[currentSum]++

        dfs(root.Left, currentSum)

        if root.Left != nil {
            prefix[currentSum+root.Left.Val]--
        }

        dfs(root.Right, currentSum)

        if root.Right != nil {
            prefix[currentSum+root.Right.Val]--
        }
    }

    dfs(root, 0)

    return result
}

// *************************************************
// 修改了回溯退场时机，居然也对了
// *************************************************
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {

    result := 0

    prefix := make(map[int]int)
    prefix[0] = 1 // 表示还没出发

    var dfs func(root *TreeNode, currentSum int)

    dfs = func(root *TreeNode, currentSum int){
        if root == nil {
            return
        }

        // 更新当前前缀和
        currentSum += root.Val

        if count, exists := prefix[currentSum-targetSum]; exists{
            result += count
        }

        prefix[currentSum]++

        dfs(root.Left, currentSum)

        dfs(root.Right, currentSum)

        prefix[currentSum]--
    }

    dfs(root, 0)

    return result
}

// *************************************************
// 最终完善
// *************************************************

func pathSum(root *TreeNode, targetSum int) int {

    result := 0

    prefix := make(map[int]int)
    prefix[0] = 1 // 表示还没出发

    var dfs func(root *TreeNode, currentSum int)

    dfs = func(root *TreeNode, currentSum int){
        if root == nil {
            return
        }

        // 更新当前前缀和
        currentSum += root.Val

        // 1. 先查
        result += prefix[currentSum-targetSum] // 包含了为0的情况

        // 2. 入场
        prefix[currentSum]++

        dfs(root.Left, currentSum)
        dfs(root.Right, currentSum)

        prefix[currentSum]-- // 恢复现场
    }

    dfs(root, 0)

    return result
}
