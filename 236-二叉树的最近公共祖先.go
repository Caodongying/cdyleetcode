package cdyleetcode

// *************************************************
// ❌❌自己想的递归法，dfs路径寻找写错了，返回路径是[]
// *************************************************
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var findPath func(root, target *TreeNode, result []*TreeNode)
    findPath = func(root, target *TreeNode, result []*TreeNode) {
        if root == nil {
            return
        }
        result = append(result, root)

        if root.Val == target.Val {
            // 找到目标节点了，返回
            return
        }

        // 如果当前节点已经是叶子结点，说明目标节点不在这个路径上，需要从数组中弹出
        if root.Left == nil && root.Right == nil {
            result = result[:len(result)-1]
            return
        }

        findPath(root.Left, target, result) // 参数要写对
        findPath(root.Right, target, result)
    }

    pathP := []*TreeNode{} // 别丢了{}
    pathQ := []*TreeNode{}

    findPath(root, p, pathP)
    findPath(root, q, pathQ)

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

// *************************************************
// ❌❌改正了数组直接作为参数传递的错误（因为这个数组是希望被直接更改的），但因为节点弹出的错误处理，还是不对
// *************************************************

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var findPath func(root, target *TreeNode, result *[]*TreeNode) bool
    findPath = func(root, target *TreeNode, result *[]*TreeNode) bool {
        if root == nil {
            return false
        }
        *result = append(*result, root)

        if root.Val == target.Val {
            return true
        }

        if root.Left == nil && root.Right == nil {
            *result = (*result)[:len(*result)-1]
            return false
        }

        if findPath(root.Left, target, result) {
            return true
        }
        return findPath(root.Right, target, result)
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


// *************************************************
// dfs寻找路径法。在Gemini提示下，发现右边节点在不是叶子结点也不是目标节点的情况下，没有正确弹出
// *************************************************

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var findPath func(root, target *TreeNode, result *[]*TreeNode) bool
    findPath = func(root, target *TreeNode, result *[]*TreeNode) bool {
        if root == nil {
            return false
        }
        *result = append(*result, root)

        if root.Val == target.Val {
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

// *************************************************
// 题解的递归回溯，直接背！
// *************************************************

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    // 基础情况：如果你自己就是 p 或 q，那你就是那个“被发现的人”，直接向上头汇报
    if root == nil || root == p || root == q{
        return root
    }


    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)

    if left != nil && right != nil {
        return root
    }

    if left != nil {
        return left
    }

    return right
}


// 换一种方法：递归回溯
// 先分析几种情况
// 1. p和q都在左子树中，那么结果就是先遍历到的节点（p或者q）。此时右子树结果就是nil
// 2. p和q都在右子树中，同上
// 3. p和q一个在左子树一个在右子树，那么结果就是当前节点

// 不要把它看作“找祖先”，把它看作“搜寻队员”。
// 想象你是一个搜救队长，你在根节点。你派了两个分队分别去左子树和右子树搜寻 $p$ 和 $q$。