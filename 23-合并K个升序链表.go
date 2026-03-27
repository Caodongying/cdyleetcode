package cdyleetcode

// *************************************************
// ❌❌ 错误的递归
// *************************************************
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }

    if len(lists) == 1 {
        return lists[0]
    }


    if len(lists) == 2 {
        // 合并两个链表
        result := ListNode{}
        curr := result
        p, q := lists[0], lists[1]
        for p != nil && q != nil {
            if p.Val <= q.Val {
                curr.Next = p
                curr = curr.Next
                p = p.Next
            } else {
                curr.Next = q
                curr = curr.Next
                q = q.Next
            }
        }

        if p != nil {
            curr.Next = p
        } else if q != nil {
            curr.Next = q
        }

        return result.Next
    }

    return mergeKLists(lists[0], mergeKLists(lists[1]))
}

// 递归法：合并第一个链表和后面的链表（假设后面的链表已经合并）
// 递归退出条件：链表数组里只有两个链表了

// *************************************************
// 递归法思路很顺畅，写起来误把输入当成二维数组所以磕绊了一些。返回值是mergeTwoLists(lists[0], mergeKLists(lists[1:]))
// *************************************************
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }

    if len(lists) == 1 {
        return lists[0]
    }

    var mergeTwoLists func(list1, list2 *ListNode) *ListNode
    mergeTwoLists = func(list1, list2 *ListNode) *ListNode{
        if list1 == nil && list2 == nil {
            return nil
        }

        if list1 == nil {
            return list2
        }

        if list2 == nil {
            return list1
        }

        // 合并两个链表
        result := &ListNode{}
        curr := result
        p, q := list1, list2
        for p != nil && q != nil {
            if p.Val <= q.Val {
                curr.Next = p
                curr = curr.Next
                p = p.Next
            } else {
                curr.Next = q
                curr = curr.Next
                q = q.Next
            }
        }

        if p != nil {
            curr.Next = p
        } else if q != nil {
            curr.Next = q
        }

        return result.Next
    }

    return mergeTwoLists(lists[0], mergeKLists(lists[1:]))
}

// 递归法：合并第一个链表和后面的链表（假设后面的链表已经合并）
// 递归退出条件：链表数组里只有两个链表了
// 本题要格外注意输入类型，不是数组的数组，而是单数组，只不过类型是”链表头结点“

// *************************************************
// 简化边界条件，优化效果一般
// *************************************************
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }

    if len(lists) == 1 {
        return lists[0]
    }

    var mergeTwoLists func(list1, list2 *ListNode) *ListNode
    mergeTwoLists = func(list1, list2 *ListNode) *ListNode{
        // 合并两个链表
        dummy := &ListNode{}
        curr := dummy
        for list1 != nil && list2 != nil {
            if list1.Val <= list2.Val {
                curr.Next = list1
                curr = curr.Next
                list1 = list1.Next
            } else {
                curr.Next = list2
                curr = curr.Next
                list2 = list2.Next
            }
        }

        if list1 != nil {
            curr.Next = list1
        } else if list2 != nil {
            curr.Next = list2
        }

        return dummy.Next
    }

    return mergeTwoLists(lists[0], mergeKLists(lists[1:]))
}

// 递归法：合并第一个链表和后面的链表（假设后面的链表已经合并）
// 递归退出条件：链表数组里只有两个链表了
// 本题要格外注意输入类型，不是数组的数组，而是单数组，只不过类型是”链表头结点“

// *************************************************
// 递归+分治法：先对合并左半边，再合并右半边，最后合并左右，效果好！
// *************************************************
func mergeTwoLists(list1, list2 *ListNode) *ListNode{
    // 合并两个链表
    dummy := &ListNode{}
    curr := dummy
    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
            curr.Next = list1
            curr = curr.Next
            list1 = list1.Next
        } else {
            curr.Next = list2
            curr = curr.Next
            list2 = list2.Next
        }
    }

    if list1 != nil {
        curr.Next = list1
    } else if list2 != nil {
        curr.Next = list2
    }

    return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
    n := len(lists)
    if n == 0 {
        return nil
    }

    if n == 1 {
        return lists[0]
    }

    left := mergeKLists(lists[:n/2])
    right := mergeKLists(lists[n/2:])
    return mergeTwoLists(left, right)
}

// 递归+分治法：先对合并左半边，再合并右半边，最后合并左右

