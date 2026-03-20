package cdyleetcode

// *************************************************
// 看的灵神题解，用的归并排序和快慢指针，时间复杂度O(nlogn)，空间复杂度O(logn）。但是一开始递归的终止条件漏了head.Next==nil
// *************************************************
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil { // 一定不要漏了head.Next==nil
        return head
    }

    // 寻找链表的中点
    prev, slow, fast := head, head, head
    for fast != nil && fast.Next != nil{
        prev = slow // prev是slow的前一个
        slow = slow.Next
        fast = fast.Next.Next
    }
    // 拆分链表，一分为二
    // 此时prev是前一半链表的最后一个节点，slow是后一半链表的第一个节点
    prev.Next = nil

    // 对两个链表进行排序
    head1 := sortList(head)
    head2 := sortList(slow)

    // 合并链表
    result := mergeListNode(head1, head2)
    return result
}

func mergeListNode(head1 *ListNode, head2 *ListNode) *ListNode {
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

// 暴力解法：冒泡排序
// 需要O(n^2)的时间复杂度
// 从第一个节点开始，如果之后有节点比第一个结点小，就交换两节点
// 然后从第二个节点，向后做一样的比较和交换
// 这种方法不太好，因为链表交换结点比较麻烦

// 看了华南溜达虎的题解，知道应该用归并排序
// 先通过一次遍历，确定mid点
// 接着拆分链表，一分为二
// 左边排序，右边排序
// 排序：当链表长度为1的时候，比较，然后大的链到小的后面
// 之后归并两个链表



// *************************************************
// ❌❌ 快排法，用基准值和三个哨兵节点，但是merge函数出错了
// *************************************************
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    base := head.Val

    smallDummy := &ListNode{}
    equalDummy := &ListNode{}
    largeDummy := &ListNode{}

    small := smallDummy
    equal := equalDummy
    large := largeDummy

    for p:=head; p!=nil; p=p.Next{
        if p.Val < base {
            small.Next = p
            small = small.Next
        } else if p.Val == base {
            equal.Next = p
            equal = equal.Next
        } else {
            large.Next = p
            large = large.Next
        }
    }

    // 三个链表结尾置空，防止成环
    small.Next, equal.Next, large.Next = nil, nil, nil

    sortList(smallDummy.Next)
    sortList(largeDummy.Next)

    result := mergeThreeListNode(smallDummy.Next, equalDummy.Next, largeDummy.Next)

    return result

}

func mergeThreeListNode(small, equal, large *ListNode) *ListNode {
    resultHead := &ListNode{}
    curr := resultHead

    linkNewList := func(newList *ListNode) {
        if newList != nil {
            curr.Next = newList
            curr = curr.Next
        }

        for curr.Next != nil {
            curr = curr.Next
        }

    }

    linkNewList(small)
    linkNewList(equal)
    linkNewList(large)

    return resultHead.Next
}

// 用快排，也就是选取基准值，遍历，小于基准值的放在左边，大于基准值的放在右边
// 对于链表，可以用三个哨兵节点，分别表示大于基准、小于基准、等于基准
// 注意，需要递归，因为比如第一次遍历的时候，虽然分出了三个链表，但是这三个链表各自是无序的
// 还有，遍历完后，要将三个链表的尾部置为nil，防止形成环

// *************************************************
// ❌❌ 修正了没有保存sortList返回值的错误，但是链接的函数出错了
// *************************************************
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    base := head.Val

    smallDummy := &ListNode{}
    equalDummy := &ListNode{}
    largeDummy := &ListNode{}

    small := smallDummy
    equal := equalDummy
    large := largeDummy

    for p:=head; p!=nil; p=p.Next{
        if p.Val < base {
            small.Next = p
            small = small.Next
        } else if p.Val == base {
            equal.Next = p
            equal = equal.Next
        } else {
            large.Next = p
            large = large.Next
        }
    }

    // 三个链表结尾置空，防止成环
    small.Next, equal.Next, large.Next = nil, nil, nil

    small = sortList(smallDummy.Next)
    large = sortList(largeDummy.Next)

    result := mergeThreeListNode(small, equalDummy.Next, large)

    return result

}

func mergeThreeListNode(small, equal, large *ListNode) *ListNode {
    resultHead := &ListNode{}
    curr := resultHead

    linkNewList := func(newList *ListNode) {
        if newList != nil {
            curr.Next = newList
            curr = curr.Next
        }

        for curr.Next != nil {
            curr = curr.Next
        }

    }

    linkNewList(small)
    linkNewList(equal)
    linkNewList(large)

    return resultHead.Next
}


// *************************************************
// 修正linkNewList，但是快排就是超时了，最坏情况时间复杂度O(n^2)，就是链表本身有序，则退化为单链表
// *************************************************
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    base := head.Val

    smallDummy := &ListNode{}
    equalDummy := &ListNode{}
    largeDummy := &ListNode{}

    small := smallDummy
    equal := equalDummy
    large := largeDummy

    for p:=head; p!=nil; p=p.Next{
        if p.Val < base {
            small.Next = p
            small = small.Next
        } else if p.Val == base {
            equal.Next = p
            equal = equal.Next
        } else {
            large.Next = p
            large = large.Next
        }
    }

    // 三个链表结尾置空，防止成环
    small.Next, equal.Next, large.Next = nil, nil, nil

    small = sortList(smallDummy.Next)
    large = sortList(largeDummy.Next)

    result := mergeThreeListNode(small, equalDummy.Next, large)

    return result

}

func mergeThreeListNode(small, equal, large *ListNode) *ListNode {
    resultHead := &ListNode{}
    curr := resultHead

    linkNewList := func(newList *ListNode) {
        if newList == nil {
            return
        }

        curr.Next = newList
        curr = curr.Next

        for curr.Next != nil {
            curr = curr.Next
        }

    }

    linkNewList(small)
    linkNewList(equal)
    linkNewList(large)

    return resultHead.Next
}

// 用快排，也就是选取基准值，遍历，小于基准值的放在左边，大于基准值的放在右边
// 对于链表，可以用三个哨兵节点，分别表示大于基准、小于基准、等于基准
// 注意，需要递归，因为比如第一次遍历的时候，虽然分出了三个链表，但是这三个链表各自是无序的
// 还有，遍历完后，要将三个链表的尾部置为nil，防止形成环

// *************************************************
// ❌❌ O(1)空间复杂度的自底向上法，getSplitList切割函数多切了。
// *************************************************
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
    curr := head
    for i:=1; i <= step && curr != nil; i++ {
        curr = curr.Next
    }

    // 如果链表长度<=size
    if curr == nil {
        return nil
    }

    // 正常切割
    secondHead := curr.Next
    // curr.Next = nil

    return secondHead
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
            // head1 := getSplitList(curr, step) // 获取从curr开始的长度为step的链表
            head1 := curr
            head2 := getSplitList(curr, step)
            curr = curr.Next // 下次分割从这开始
            // 合并head1和head2
            merged := mergeTwoList(head1, head2)
            newCurr.Next = merged
            for newCurr.Next != nil {
                newCurr = newCurr.Next
            }
        }

        curr = dummyHead.Next

    }

    return dummyHead.Next
}

// 自底向上，step从1开始双倍增加
// 多次遍历，直到step大于等于原链表的长度
// 在每一轮遍历里：按照step，先顺序分割出两组（每组大小为step），然后merge。之后继续分割，直到链表末尾。这个过程也是在循环里实现
// 每一轮遍历的最后，更新step


// *************************************************
// 修正了getSplitList的逻辑和应用
// *************************************************
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
            // head1 := getSplitList(curr, step) // 获取从curr开始的长度为step的链表
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

    return dummyHead.Next
}

// 自底向上，step从1开始双倍增加
// 多次遍历，直到step大于等于原链表的长度
// 在每一轮遍历里：按照step，先顺序分割出两组（每组大小为step），然后merge。之后继续分割，直到链表末尾。这个过程也是在循环里实现
// 每一轮遍历的最后，更新step