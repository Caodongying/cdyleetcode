package cdyleetcode

// *************************************************
// 搜索树中序遍历，就是有序数组
// *************************************************
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {

    inorderSortResult := []int{}

    var inorderSort func(root *TreeNode) []int
    inorderSort = func(root *TreeNode) []int {
        if root == nil {
            return nil
        }

        inorderSort(root.Left)
        inorderSortResult = append(inorderSortResult, root.Val)
        inorderSort(root.Right)

        return inorderSortResult
    }

    inorderSort(root)

    return inorderSortResult[k-1]
}

// 下面是误区（一开始以为输入是[]int，受题干误导了）
// root是层序遍历的结果,而且是“满”的，就是空结点都用null表示
// 其实就是对搜索树的层序遍历的结果进行排序

// 反向考虑：如果要得到层序遍历的结果，用队列
// 那就用两个指针，第一个表示“当前节点”，第二个表示“当前节点的子节点”，第二个指针每次向后移两个（跳过null节点）。
// 儿子指针，第一个值，放在当前节点指针的左边。第二个值，放在当前节点指针的右边

// 现在是正确的：其实搜索树中序遍历，就是有序数组


// *************************************************
//
// *************************************************

// *************************************************
//
// *************************************************

// *************************************************
//
// *************************************************
