package cdyleetcode

// *************************************************
// ❌❌ 粗略看了题解后乱写的，没有断开分组
// *************************************************
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    // 计算链表总长度
    n := 0
    for curr := head; curr != nil; curr = curr.Next {
        n++
    }

    dummy := &ListNode{Next: head,}
    prev := dummy // 初始化为哨兵节点
    curr := dummy.Next // 初始化为head
    prevHead := dummy.Next

// 假设k为2
// 初始dummy -> 1 -> 2 -> 3 -> 4 -> 5 i=0, 执行后next是2，prev是1，curr是2
// 初始dummy <- 1 -> 2 -> 3 ->4 -> 5 i=1, 执行后next是3，prev是2，curr是3
// 初始dummy <- 1 <- 2 -> 3 ->4 -> 5 此时退出内层循环，下一步需要dummy的Next指向2，上一轮翻转前的实际头结点指向3，


    for ; n >= k; n=n-k {
        // 翻转当前k长度的链表
        for i:=0; i<k; i++ {
            next := curr.Next
            curr.Next = prev
            prev = curr
            curr = next
        }

        prevHead.Next = curr // 把上一轮翻转的链表和下一轮的头结点连起来
        prevHead = curr // 更新现在这一轮的头结点
    }

    return dummy.Next
}



// 暴力解法
// 先算出链表总长度，作为判断剩下的链表长度是否满足翻转条件的依据
// 用for循环，步长为k
//   每次循环：保存原本的链表尾指向的元素，把子链表尾指向nil，然后翻转链表

// 基础的链表翻转：https://www.bilibili.com/video/BV1SCtBzyESd/?spm_id_from=333.337.search-card.all.click&vd_source=e3d59376cbbf7e58fec4949fdb3a00d4

// *************************************************
// 看了题解写写改改
// *************************************************
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    // 计算链表总长度
    n := 0
    for curr := head; curr != nil; curr = curr.Next {
        n++
    }

    // 创建虚拟头节点
    dummy := &ListNode{Next: head}
    prev := dummy  // prev 指向待翻转子链表的前一个节点
    curr := head   // curr 指向当前组的头节点

    for ;n >= k; n-=k {
        // 截取当前组
        start, end := curr, curr

        for i:=0; i<k-1; i++ {
            end = end.Next
        }

        // 断开当前组
        curr = end.Next // 保存下一个组的头结点，注意别写成curr := ...
        end.Next = nil

        // 翻转当前组，然后把翻转后的链表和前面的链表连起来
        prev.Next = reverseList(start) // 翻转后，链表尾就是start，链表头是end

        // 把翻转后的链表和后面的链表连起来
        start.Next = curr


        prev = start // 不是prev = curr

    }

    return dummy.Next
}

// 翻转链表的函数
// 用prev, curr, next三个指针，不用dummyNode
// 迭代法翻转链表
// 其实就是从开始到结尾，依次把当前指针的Next指向前一个节点
// 这就涉及对原来的后一个节点的保存
func reverseList(head *ListNode) *ListNode {
    curr := head
    var prev *ListNode

    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }

    return prev
}

// *************************************************
// 小小改动一下
// *************************************************
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    // 计算链表总长度
    n := 0
    for curr := head; curr != nil; curr = curr.Next {
        n++
    }

    // 创建虚拟头节点
    dummy := &ListNode{Next: head}
    prev := dummy  // prev 指向待翻转子链表的前一个节点
    curr := head   // curr 指向当前组的头节点

    for ;n >= k; n-=k {
        // 截取当前组
        start, end := curr, curr

        for i:=0; i<k-1; i++ {
            end = end.Next
        }

        // 断开当前组
        curr = end.Next // 保存下一个组的头结点，注意别写成curr := ...
        end.Next = nil

        // 翻转当前组，然后把翻转后的链表和前面的链表连起来
        reversedHead := reverseList(start)
        prev.Next = reversedHead // 翻转后，链表尾就是start，链表头是end

        // 把翻转后的链表和后面的链表连起来
        start.Next = curr

        prev = start // 不是prev = curr

    }

    return dummy.Next
}

// 翻转链表的函数
// 用prev, curr, next三个指针，不用dummyNode
// 迭代法翻转链表
// 其实就是从开始到结尾，依次把当前指针的Next指向前一个节点
// 这就涉及对原来的后一个节点的保存
func reverseList(head *ListNode) *ListNode {
    curr := head
    var prev *ListNode

    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }

    return prev
}