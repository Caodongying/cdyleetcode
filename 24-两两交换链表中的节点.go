package cdyleetcode

// *************************************************
// 莫名其妙就写起来了，乱七八糟的
// *************************************************
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    // dummyHead := &ListNode{0, head}
    if head == nil {
        return nil
    }
    dummyHead := &ListNode{0, head}
    left := head
    right := head.Next
    prev := dummyHead
    for ; left != nil; {
        // 交换左右指针
        if right == nil {
            // 左指针是最后一个节点了
            break
        } else{
            left.Next = right.Next
            right.Next = left
        }

        // 更新左右指针
        prev.Next = right
        prev = left
        left = left.Next
        if left == nil {
            // 已经到最后一个节点了
            break
        } else {
            right = left.Next
        }
    }

    return dummyHead.Next
}

// 两个指针
// 循环在右指针变成空的时候暂停
// 初始化的时候确定下来新的head

// *************************************************
// 改进一下三指针的代码
// *************************************************
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    // dummyHead := &ListNode{0, head}
    if head == nil {
        return nil
    }
    dummyHead := &ListNode{0, head}
    left, right := head, head
    prev := dummyHead
    for ; left != nil && left.Next != nil; {
        right = left.Next
        // 交换左右指针
        left.Next = right.Next
        right.Next = left

        // 更新左右指针
        prev.Next = right
        prev = left
        left = left.Next
    }

    return dummyHead.Next
}

// 两个指针
// 循环在右指针变成空的时候暂停
// 初始化的时候确定下来新的head

// *************************************************
// 递归法，秒了
// *************************************************
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    if head.Next == nil {
        return head
    }

    left, right := head, head.Next

    // 交换左右指针
    left.Next = swapPairs(right.Next)
    right.Next = left

    return right
}