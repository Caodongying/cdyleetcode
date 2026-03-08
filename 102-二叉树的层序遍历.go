package cdyleetcode

// *************************************************
// 队列，但是为了维护层级信息，新开了一个Node队列，空间消耗大
// *************************************************

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }

    queue := []*TreeNode{}
    queue = append(queue, root)

    result := [][]int{}

    for len(queue) > 0 {
        levelResult := []int{}
        levelQueue := make([]*TreeNode, len(queue)) // 不可以写成 levelQueue:=[]*TreeNode{}
        copy(levelQueue, queue)
        // 先一次性将队列里的元素写到当前level的数组
        for len(queue) > 0 {
            front := queue[0]
            queue = queue[1:]

            levelResult = append(levelResult, front.Val)

        }
        result = append(result, levelResult)

            // 再通过临时数组，处理当前level的子节点
        for _, node := range(levelQueue) {
            if node.Left!=nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                    queue = append(queue, node.Right)
            }

        }
    }


    return result
}

// 用队列实现
// 右进左出，从队列头部取出，从队列尾部压入
// 每一层结果要单独保存，可以一次性取出队列所有元素，保存在结果数组中，然后再遍历这个结果数组进行队尾元素添加

// *************************************************
// 不用第二个辅助队列，用长度记录层级信息
// *************************************************
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }

    queue := []*TreeNode{}
    queue = append(queue, root)

    result := [][]int{}

    for len(queue) > 0 {
        levelResult := []int{}

        levelLen := len(queue)

        for levelLen > 0 {
            front := queue[0]
            queue = queue[1:]

            levelResult = append(levelResult, front.Val)

            if front.Left!=nil {
                queue = append(queue, front.Left)
            }
            if front.Right != nil {
                queue = append(queue, front.Right)
            }

            levelLen--

        }
        result = append(result, levelResult)

    }


    return result
}

// 用队列实现
// 右进左出，从队列头部取出，从队列尾部压入

// *************************************************
// 稍微有点区别的写法，用数组而不是切片，避免动态内存分配
// *************************************************
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }

    queue := []*TreeNode{}
    queue = append(queue, root)

    result := [][]int{}

    for len(queue) > 0 {

        levelLen := len(queue)
        levelResult := make([]int, levelLen)

        for i := range levelLen {
            front := queue[0]
            queue = queue[1:]

            levelResult[i] = front.Val

            if front.Left!=nil {
                queue = append(queue, front.Left)
            }
            if front.Right != nil {
                queue = append(queue, front.Right)
            }

        }
        result = append(result, levelResult)

    }


    return result
}

// 用队列实现
// 右进左出，从队列头部取出，从队列尾部压入
// 每一层结果要单独保存，可以一次性取出队列所有元素，保存在结果数组中，然后再遍历这个结果数组进行队尾元素添加

// *************************************************
//
// *************************************************
