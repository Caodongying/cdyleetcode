package cdyleetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode  {
    if head == nil {
        return nil
    }

    p := head

    nodeMap := make(map[*ListNode]struct{})
    for p != nil {
        // 检查map
        if _, ok := nodeMap[p]; ok {
            return p
        }

        nodeMap[p] = struct{}{}
        p = p.Next
    }

    return nil
}