package cdyleetcode

// *************************************************
// 用map
// *************************************************

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }

    p := head
    nodeMap := make(map[*ListNode]int)
    for p != nil {
        // 检查map
        if nodeMap[p] == 1 {
            return true
        } else {
            nodeMap[p] = 1
        }
        p = p.Next
    }

    return false
}

// *************************************************
// 优化map，用struct{}
// *************************************************
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }

    p := head
    // nodeMap := make(map[*ListNode]int)
    nodeMap := make(map[*ListNode]struct{})
    for p != nil {
        // 检查map
        if _, ok := nodeMap[p]; ok {
            return true
        }

        nodeMap[p] = struct{}{}
        p = p.Next
    }

    return false
}

// *************************************************
// 快慢指针法
// *************************************************
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }

    slow, fast := head, head

    for slow != nil && fast!=nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next

        if slow == fast {
            return true
        }
    }

    return false
}