package cdyleetcode

// *************************************************
// 递归法
// *************************************************

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    }

    if list2 == nil {
        return list1
    }

    var newList *ListNode

    if list1.Val <= list2.Val {
        newList = &ListNode{Val: list1.Val}
        newList.Next = mergeTwoLists(list1.Next, list2)
    } else {
        newList = &ListNode{Val: list2.Val}
        newList.Next = mergeTwoLists(list1, list2.Next)
    }

    return newList
}