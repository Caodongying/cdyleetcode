package cdyleetcode

// *************************************************
// 复制链表+翻转链表
// *************************************************
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    copiedHead := copyList(head)
    reversedLinkHead := reverseLink(copiedHead)

    p := head
    pReversed := reversedLinkHead
    result := false
    for {
        if p == nil {
            return result
        }

        if p.Val != pReversed.Val {
            return false
        } else {
            result = true
            p = p.Next
            pReversed = pReversed.Next
        }
    }
}

func reverseLink(head *ListNode) *ListNode {
    var prev *ListNode = nil
    curr := head

    for curr != nil {
        nextTemp := curr.Next // 保存下一个节点
        curr.Next = prev
        prev = curr
        curr = nextTemp
    }

    return prev
}

func copyList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    copiedHead := &ListNode{Val: head.Val, Next: nil}
    p := head
    pCopied := copiedHead

    for {
        if p.Next == nil {
            break
        } else {
            tempNext := &ListNode{Val: p.Next.Val, Next: nil}
            pCopied.Next = tempNext
            pCopied = tempNext
            p = p.Next
        }
    }

    return copiedHead
}

// 先翻转链表(记得要复制原来的head)
// 再比对原链表和翻转链表是否完全一致


// *************************************************
// 快慢指针
// *************************************************
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    mid := middleNode(head)
    head2 := reverseLink(mid)
    for head2 != nil {
        if head.Val != head2.Val {
            return false
        }

        head = head.Next
        head2 = head2.Next
    }

    return true
}

func middleNode(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}

func reverseLink(head *ListNode) *ListNode {
    var prev *ListNode = nil
    curr := head

    for curr != nil {
        nextTemp := curr.Next // 保存下一个节点
        curr.Next = prev
        prev = curr
        curr = nextTemp
    }

    return prev
}

// 先用快慢指针找到中间节点
// 翻转后半部分的链表