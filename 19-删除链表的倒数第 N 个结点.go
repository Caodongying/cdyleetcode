package cdyleetcode

// *************************************************
// 磕磕绊绊的初提交
// *************************************************
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    lengthList := getListLength(head)
    nReverse := lengthList - n + 1
    if nReverse == 1 {
        // 删除的是表头
        head = head.Next
        return head
    }

    ptrPrev := &ListNode{Val:1, Next: head}
    var ptrNext *ListNode // 默认值是nil

    for i:=0; ; i++ {
        if i == nReverse-1 { // 这里很容易错
            break
        }
        ptrPrev = ptrPrev.Next
    }

    if nReverse != lengthList {
        // 删除的不是表尾
        ptrNext = ptrPrev.Next.Next
    }

    ptrPrev.Next = ptrNext

    return head
}

func getListLength(head *ListNode) int {
    length := 1

    if head.Next == nil {
        return length
    }

    for {
        head = head.Next
        if head == nil {
            break
        }
        length += 1
    }
    return length
}
// 计算链表长度length
// 把倒数第n个，转化为正数第nReverse := length-n+1个
// 如果nReverse是1，就说明要删除的就是表头，获得新表头（当然也可能是空链表了）
// 如果nReverse是length，说明要删除的是表尾
// 不然，把nReverse-1的next指向nReverse+1
// 优化：一趟扫描

// *************************************************
// 优化了一下代码，直接使用dummyHead
// *************************************************
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    lengthList := getListLength(head)
    nReverse := lengthList - n + 1
    // 问题变成删除链表中第nReverse个节点

    dummyHead := &ListNode{Val:0, Next: head}
    ptrPrev := dummyHead

    // 找到第nReverse个节点的前一个节点nReverse-1
    for i:=0; i < nReverse -1 ; i++ {
        ptrPrev = ptrPrev.Next
    }

    ptrPrev.Next = ptrPrev.Next.Next

    return dummyHead.Next
}

func getListLength(head *ListNode) int {
    length := 0
    for ; head!=nil; head = head.Next {
        length += 1
    }
    return length
}
// 计算链表长度length
// 把倒数第n个，转化为正数第nReverse := length-n+1个
// 如果nReverse是1，就说明要删除的就是表头，获得新表头（当然也可能是空链表了）
// 如果nReverse是length，说明要删除的是表尾
// 不然，把nReverse-1的next指向nReverse+1
// 优化：一趟扫描