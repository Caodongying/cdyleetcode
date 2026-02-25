package cdyleetcode

// *************************************************
// 暴力解法
// *************************************************
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    p1 := l1 // l1链表的指针
    p2 := l2 // l2链表的指针

    var l3 *ListNode
    var p3 *ListNode
    var p1Val int
    var p2Val int
    prevAdd := 0

    for {
        if p1 == nil && p2 == nil {
            if prevAdd != 0 {
                tailNode := ListNode{Val: prevAdd, Next: nil}
                p3.Next = &tailNode
            }
            return l3
        }
        if p1 != nil {
            p1Val = p1.Val
        } else {
            p1Val = 0
        }
        if p2 != nil {
            p2Val = p2.Val
        } else {
            p2Val = 0
        }

        curNodeVal := (p1Val + p2Val + prevAdd) % 10
        prevAdd = (p1Val + p2Val + prevAdd) / 10
        newNode := &ListNode{Val: curNodeVal, Next: nil} // 必须是&
        if l3 == nil {
            l3 = newNode
            p3 = l3
        } else {
            p3.Next = newNode
            p3 = p3.Next
        }

        if p1 != nil { // 一开始错了
            p1 = p1.Next
        }

        if p2 != nil {
            p2 = p2.Next
        }
    }

    return l3


}

// 从低到高，一个个算