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

    nodemap := make(map[int]int) // key和value都是Node的地址。命名不可以是map

    currentCopy := &Node{} // 哨兵节点，虚拟头结点

    for p, q:=head, currentCopy; ; p, q = p.Next, q.Next {
        // 检查当前节点的值是否已经拷贝过
        if addr, exists := nodemap[p]; exists {
            // 已经拷贝过，新链表直接链
            q.Next = addr
            q = q.Next
        } else{
            // 没有拷贝过，创建新节点
            newNode := &Node{
                Val: p.Val,
            }
            nodemap[p] = newNode
        }

        // 创建random
        if p.Random == nil {
            newNode.Random = nil
        } else {
            if addr, exists := nodemap[p.Random]; exists{
                newNode.Random = addr
            } else {
                newRandomNode := &Node{
                    Val: p.Random.Val,
                }
                nodemap[p.Random] = newRandomNode
        }

        // 创建next
        if p.Next == nil {
            q.Next = newNode
            q = q.Next
            q.Next = nil
            break
        }
        if addr, exists := nodemap[p.Next]; exists{
            newNode.Next = addr
        } else {
            newNextNode := &Node{
                Val: p.Next.Val,
            }
            nodemap[p.Next] = newNextNode
        }

        // 新链表，链接新节点
        q.Next = newNode
        q = q.Next

    }

    return currentCopy.Next
}
}

// 关键点是存储Random
// 建立原链表中的节点node1和新链表中的复制节点node2的对应关系
// 可以用一个nodemap存储这个映射关系


func main() {
    // [3,5,1,6,2,0,8,null,null,7,4]
    root := &TreeNode{Val: 3}
    root.Left = &TreeNode{Val: 5}
    root.Right = &TreeNode{Val: 1}
    root.Left.Left = &TreeNode{Val: 6}
    root.Left.Right = &TreeNode{Val: 2}
    root.Right.Left = &TreeNode{Val: 0}
    root.Right.Right = &TreeNode{Val: 8}
    root.Left.Right.Left = &TreeNode{Val: 7}
    root.Left.Right.Right = &TreeNode{Val: 4}

    p := &TreeNode{Val: 5}
    q := &TreeNode{Val: 1}

    result := lowestCommonAncestor(root, p, q)

    fmt.Println(result.Val)
}
