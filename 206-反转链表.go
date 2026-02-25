package cdyleetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // 迭代法
func reverseList1(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    currentHead := head
    headNext := head.Next
    head.Next = nil //非常重要


    for headNext != nil {
        headNextNext := headNext.Next
        headNext.Next = currentHead
        currentHead = headNext

        headNext = headNextNext
    }

    return currentHead
}

// 递归法
func reverseList(head *ListNode) *ListNode {
    cur := head

    if cur == nil {
        return nil
    }

    if cur.Next == nil {
        return cur
    }

    tempListHead := reverseList(cur.Next)
    curNext := cur.Next
    curNext.Next = cur
    cur.Next = nil

    return tempListHead
}