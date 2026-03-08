package cdyleetcode

// *************************************************
// ❌❌错误，没有考虑到右子树中也许有节点比根节点小，无法检测到
// *************************************************
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }

    if root.Left != nil && root.Left.Val >= root.Val {
        return false
    }

    if root.Right != nil && root.Right.Val <= root.Val {
        return false
    }

    return isValidBST(root.Left) && isValidBST(root.Right)
}

// 用递归法


// *************************************************
// ❌❌依然错误。
// 错误，因为只检查了子节点和当前节点的关系，但没有检查当前节点和祖先节点传递下来的边界的关系。考虑这个无效BST：

//     5
//    / \
//   3   7
//      / \
//     4   8

// 根节点5：option=-1，不检查边界

// 节点7：option=0，检查 7 > 5 ✅

// 节点4：option=1（因为它是7的左子节点），检查 4 < 7 ✅

// 但这里节点4没有检查是否 4 > 5，而这是必须的！
// *************************************************

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {

    isValidWithLimit := func(root *TreeNode, min *int, max *int) bool {
        if root == nil {
            return true
        }

		// 检查当前值是否在上下限范围以内
		if min != nil && root.Val <= *min {
			return false
		}

		if max != nil && root.Val >= *max {
			return false
		}

		// 检查左节点
        if root.Left != nil && root.Left.Val >= root.Val {
            return false
        }

		// 检查右节点
        if root.Right != nil && root.Right.Val <= root.Val {
            return false
        }

        return isValidWithLimit(root.Left, min, &root.Val) && isValidWithLimit(root.Right, &root.Val, max)
    }

    return isValidWithLimit(root, &Limit{option: -1})

}


// *************************************************
// 成功的递归法，用了双边界
// *************************************************
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    var isValidWithLimit func(root *TreeNode, min *int, max *int) bool
    isValidWithLimit = func(root *TreeNode, min *int, max *int) bool {
        if root == nil {
            return true
        }

		// 检查当前值是否在上下限范围以内
		if min != nil && root.Val <= *min {
			return false
		}

		if max != nil && root.Val >= *max {
			return false
		}

		// 检查左节点
        if root.Left != nil && root.Left.Val >= root.Val {
            return false
        }

		// 检查右节点
        if root.Right != nil && root.Right.Val <= root.Val {
            return false
        }

        return isValidWithLimit(root.Left, min, &root.Val) && isValidWithLimit(root.Right, &root.Val, max)
    }

    return isValidWithLimit(root, nil, nil)

}

// 用递归法


// *************************************************
//
// *************************************************
