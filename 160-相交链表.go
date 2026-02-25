package cdyleetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    // 统计A和B各自的长度
    lengthA := countListLength(headA)
    lengthB := countListLength(headB)

    var pLonger *ListNode
    var pShorter *ListNode
    offset := 0
    if lengthA >= lengthB {
        pLonger = headA
        pShorter = headB
        offset = lengthA - lengthB
    } else {
        pLonger = headB
        pShorter = headA
        offset = lengthB - lengthA
    }


    for i:=0; i<offset; i++ {
        pLonger = pLonger.Next
    }

    // 开始遍历
    for {
        if pShorter == pLonger{
            return pShorter
        }
        pShorter = pShorter.Next
        pLonger = pLonger.Next
    }
}

func countListLength(head *ListNode) int {
    length := 0
    p := head
    for {
        if p == nil {
            return length
        }

        length += 1
        p = p.Next
    }
}
// 让长链表先走完和短链表之间的偏差
// 然后用双指针进行判断