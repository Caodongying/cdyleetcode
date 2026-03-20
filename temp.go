package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type Node struct {
    Val int
    Next *Node
    Random *Node
}


type ListNode struct {
    Val int
    Next *ListNode
}


/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func printListNode(head *ListNode) {
    curr := head
    for curr != nil {
        fmt.Printf("%d ", curr.Val)
        curr = curr.Next
    }
    fmt.Println()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 2 4 3 7 -1 0 3 2
// [2 4] [3 7] [-1 0] [3 2]

func getListLength(head *ListNode) int {
    length := 0

    for head != nil {
        length++
        head = head.Next
    }

    return length
}

// 分割链表，前step个节点分割出去，返回剩余链表的头结点
func getSplitList(head *ListNode, step int) *ListNode {
    curr, prev := head, head
    for i:=1; i <= step && curr != nil; i++ {
        prev = curr
        curr = curr.Next
    }

    // 如果链表长度<=size
    if curr == nil {
        return nil
    }

    prev.Next = nil

    fmt.Println("原链表")
    printListNode(head)
    fmt.Printf("切割%d长度后的链表为: ", step)
    printListNode(curr)

    return curr
}

func mergeTwoList(head1 *ListNode, head2 *ListNode) *ListNode {
    resultHead := &ListNode{}
    curr := resultHead

    for head1 != nil && head2 != nil {
        if head1.Val < head2.Val {
            curr.Next = head1
            head1 = head1.Next
        } else {
            curr.Next = head2
            head2 = head2.Next
        }
        curr = curr.Next
    }

    if head1 != nil {
        curr.Next = head1
    } else if head2 != nil {
        curr.Next = head2
    }

    return resultHead.Next
}

func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    var step int
    listLength := getListLength(head)
    curr := head
    var dummyHead *ListNode

    for step = 1 ; step <= listLength; step = step * 2 {
        dummyHead = &ListNode{}
        newCurr := dummyHead
        for curr != nil {
            head1 := curr // 此时head1后面不是nil，而是整个链表
            head2 := getSplitList(head1, step) // 从head1后面开始切割，并且把head1和后面断开
            curr = getSplitList(head2, step) // 下次分割从这开始
            // 合并head1和head2
            merged := mergeTwoList(head1, head2)
            newCurr.Next = merged
            for newCurr.Next != nil {
                newCurr = newCurr.Next
            }
        }

        curr = dummyHead.Next

    }

    fmt.Println("最终结果是:")
    printListNode(dummyHead.Next)
    return dummyHead.Next
}

// 自底向上，step从1开始双倍增加
// 多次遍历，直到step大于等于原链表的长度
// 在每一轮遍历里：按照step，先顺序分割出两组（每组大小为step），然后merge。之后继续分割，直到链表末尾。这个过程也是在循环里实现
// 每一轮遍历的最后，更新step

func main() {
    // [-1,5,3,4,0]
    head := &ListNode{Val: -1}
    head.Next = &ListNode{Val: 5}
    head.Next.Next = &ListNode{Val: 3}
    head.Next.Next.Next = &ListNode{Val: 4}
    head.Next.Next.Next.Next = &ListNode{Val: 0}

    sortList(head)

}
