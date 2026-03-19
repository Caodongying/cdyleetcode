package cdyleetcode

// *************************************************
// 用map存储新旧节点之间的映射关系。自己写的，边写边优化
// *************************************************
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }

    // *Node不就是int吗
    nodemap := make(map[*Node]*Node) // key和value都是Node的地址。命名不可以是map

    currentCopy := &Node{} // 哨兵节点，虚拟头结点

    for p, q:=head, currentCopy; ; p, q = p.Next, q.Next {
        // 检查当前节点的值是否已经拷贝过
        if _, exists := nodemap[p]; !exists {
            // 没有拷贝过，创建新节点
            nodemap[p] = &Node{
                Val: p.Val,
            }
        }

        q.Next = nodemap[p]
        q = q.Next

        // 创建q的random
        if p.Random == nil {
            q.Random = nil
        } else {
            if _, exists := nodemap[p.Random]; !exists{
                nodemap[p.Random] = &Node{
                    Val: p.Random.Val,
                }
            }
            q.Random = nodemap[p.Random]
        }

        // 创建q的next
        if p.Next == nil {
            q.Next = nil
            break
        }

        if _, exists := nodemap[p.Next]; !exists{
            nodemap[p.Next] = &Node{
                Val: p.Next.Val,
            }
        }
        q.Next = nodemap[p.Next]

    }

    return currentCopy.Next
}

// 关键点是存储Random
// 建立原链表中的节点node1和新链表中的复制节点node2的对应关系
// 可以用一个nodemap存储这个映射关系

// *************************************************
// 用Gemini优化，把节点复制放在函数里，简化逻辑
// *************************************************

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }

    // *Node不就是int吗
    nodemap := make(map[*Node]*Node) // key和value都是Node的地址。命名不可以是map

    getCopy := func(node *Node) *Node {
        if node == nil {
            return nil
        }

        if _, exists := nodemap[node]; !exists {
            // 没有拷贝过，创建新节点
            nodemap[node] = &Node{
                Val: node.Val,
            }
        }

        return nodemap[node]
    }

    p := head
    for p != nil {
        clone := getCopy(p)
        clone.Next = getCopy(p.Next)
        clone.Random = getCopy(p.Random)

        p = p.Next
    }

    return nodemap[head]
}

// Gemini推荐的优化：只要遇到一个节点（无论是通过 Next 还是 Random 遇到的），如果它不在 map 里，就创建它。


/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

// *************************************************
// 原地解法，在每个节点后面复制一个Val一样的节点，然后链接Random，最后恢复Next
// *************************************************

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }

    // 原地复制节点，只复制值
    curr := head

    for curr != nil {
        copy := &Node{
            Val: curr.Val,
        }
        copy.Next = curr.Next
        curr.Next = copy
        curr = copy.Next
    }

    // 复制Random指针
    curr = head
    for curr != nil {
        if curr.Random != nil {
            curr.Next.Random = curr.Random.Next
        }
        curr = curr.Next.Next
    }

    // 进行拆分
    newHead := head.Next
    curr = head
    for curr != nil {
        clone := curr.Next
        curr.Next = clone.Next // 恢复链表。注意curr.Random没有受到影响
        if curr.Next != nil {
            clone.Next = curr.Next.Next
        }
        curr = curr.Next
    }

    return newHead
}

// 原地解法，不使用map（看了灵神解法，结合Gemini讲解)
// 在每个节点后面，直接复制当前节点。比如A->B，变成A->A'->B->B'
// 这样A的复制节点就是A'(A.Next)，A.Next的复制节点就是A'.Next(A.Next.Next)，A.Random的复制节点就是A'.Random(A.Random.Next)
// 然后复制Random
// 接着根据奇偶进行拆分