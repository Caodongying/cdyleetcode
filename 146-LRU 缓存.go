package cdyleetcode

// *************************************************
// ❌❌ MoveToHead函数的“如果当前node在cache中间，更新为头部”，这一块有问题，put逻辑有问题，要单独判断已经存在的情况
// *************************************************
type LRUCache struct {
    cache map[int]*Node
    capacity int
    occupied int

    head *Node
    tail *Node
}

type Node struct {
    val int
    next *Node
    prev *Node
}

func Constructor(capacity int) LRUCache {
    lruCache := LRUCache{
        capacity: capacity,
        occupied: 0,
        cache: make(map[int]*Node, capacity),// 别忘记
    }

    return lruCache
}

func (this *LRUCache) MoveToHead(node *Node) {
    // 如果当前node已经在cache头部，返回
    if this.head == node {
        return
    }
    // 如果当前node是cache尾部，更新为头部、尾部更新为prev
    if this.tail == node {
        // 更新尾部
        tailPrev := this.tail.prev
        tailPrev.next = nil
        this.tail = tailPrev

        // 更新头部 - 一开始写错了
        oldHead := this.head
        oldHead.prev = node
        this.head = node
        node.next = oldHead
        return
    }
    // 如果当前node在cache中间，更新为头部
    nodePrev := node.prev
    nodeNext := node.next
    nodePrev.next = nodeNext
    nodeNext.prev = nodePrev

    oldHead := this.head
    oldHead.prev = node
    this.head = node
    node.next = oldHead
}

func (this *LRUCache) Get(key int) int {
    if node, exists := this.cache[key]; exists {
        // key存在，将此节点移到cache头部
        this.MoveToHead(node)
        return node.val
    }

    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    // 先判断是否有足够容量
    if this.occupied == this.capacity {
        // 移除队尾元素
        tailPrev := this.tail.prev
        tailPrev.next = nil
        this.tail = tailPrev
        this.occupied--
    }



    node := &Node{
        val: value,
        next: this.head,
        prev: nil,
    }


    this.cache[key] = node
    this.occupied++

    if this.occupied == 1 {
        this.head, this.tail = node, node
    }

    // 把这个元素移到队首
    node.next = this.head
    this.head.prev = node
    this.head = node
}

// 保存缓存内容的数据结构：数组或者链表（需要保存prev)

// 每一个get或者put，要对这个key进行“升级”：
// 把这个key的元素移到队首（队首是最新使用的，队尾是最久以前使用的）


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// *************************************************
// 自己写的，边界条件是失败一个案例就补一个，不太好。而且没有用哨兵节点，修改头尾也有大量重复逻辑
// *************************************************
type LRUCache struct {
    cache map[int]*Node
    capacity int
    occupied int

    head *Node
    tail *Node
}

type Node struct {
    val int
    key int
    next *Node
    prev *Node
}

func Constructor(capacity int) LRUCache {
    lruCache := LRUCache{
        capacity: capacity,
        occupied: 0,
        cache: make(map[int]*Node, capacity),// 别忘记
    }

    return lruCache
}

func (this *LRUCache) MoveToHead(node *Node) {
    // 如果当前node已经在cache头部，返回
    if this.head == node {
        return
    }
    // 如果当前node是cache尾部，更新为头部、尾部更新为prev
    if this.tail == node {
        // 更新尾部
        tailPrev := this.tail.prev
        tailPrev.next = nil
        this.tail = tailPrev

        // 更新头部 - 一开始写错了
        oldHead := this.head
        oldHead.prev = node
        this.head = node
        node.next = oldHead
        return
    }
    // 如果当前node在cache中间，更新为头部
    nodePrev := node.prev
    nodeNext := node.next
    nodePrev.next = nodeNext
    nodeNext.prev = nodePrev

    oldHead := this.head
    oldHead.prev = node
    this.head = node
    node.next = oldHead
}

func (this *LRUCache) Get(key int) int {
    if node, exists := this.cache[key]; exists {
        // key存在，将此节点移到cache头部
        this.MoveToHead(node)
        return node.val
    }

    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    // 要先判断是否存在，再判断容量
    if node, exists := this.cache[key]; exists {
        node.val = value
        this.MoveToHead(node)
        return
    }

    // 先判断是否有足够容量
    if this.occupied == this.capacity {
        // 移除队尾元素
        tailPrev := this.tail.prev
        if tailPrev != nil {
            tailPrev.next = nil
        }
        // 从map删除
        delete(this.cache, this.tail.key)
        this.tail = tailPrev
        this.occupied--
    }


    node := &Node{
        val: value,
        key: key,
        next: this.head,
        prev: nil,
    }


    this.cache[key] = node
    this.occupied++

    if this.occupied == 1 {
        this.head, this.tail = node, node
        return
    }

    // 把这个元素移到队首
    node.next = this.head
    this.head.prev = node
    this.head = node
}

// 保存缓存内容的数据结构：数组或者链表（需要保存prev)

// 每一个get或者put，要对这个key进行“升级”：
// 把这个key的元素移到队首（队首是最新使用的，队尾是最久以前使用的）


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// *************************************************
// 用了虚拟头尾节点，参考官方题解和deepseek修改
// *************************************************
type LRUCache struct {
    cache map[int]*Node
    capacity int
    occupied int

    dummyHead *Node
    dummyTail *Node
}

type Node struct {
    val int
    key int
    next *Node
    prev *Node
}

func Constructor(capacity int) LRUCache {
    lruCache := LRUCache{
        capacity: capacity,
        occupied: 0,
        cache: make(map[int]*Node, capacity),// 别忘记
        dummyHead: &Node{},
        dummyTail: &Node{},
    }

    // 初始化双向链表
    lruCache.dummyHead.next = lruCache.dummyTail
    lruCache.dummyTail.prev = lruCache.dummyHead

    return lruCache
}

func (this *LRUCache) moveToHead(node *Node) {
    // 把这个结点先移除
    this.removeNode(node)
    // 再把结点加到队首
    this.addToHead(node)
}

func (this *LRUCache) Get(key int) int {
    if node, exists := this.cache[key]; exists {
        // key存在，将此节点移到cache头部
        this.moveToHead(node)
        return node.val
    }

    return -1
}

func (this *LRUCache) removeNode(node *Node) {
    // 让node的前驱和后继节点直接相连
    node.prev.next = node.next
    node.next.prev = node.prev
}

func (this *LRUCache) removeTail() *Node{
    // 如果链表为空（只有虚拟节点），返回 nil
    if this.dummyTail.prev == this.dummyHead {
        return nil
    }
    tail := this.dummyTail.prev
    this.removeNode(tail)
    return tail
}

func (this *LRUCache) addToHead(node *Node) {
    node.next = this.dummyHead.next
    this.dummyHead.next.prev = node
    this.dummyHead.next = node
    node.prev = this.dummyHead
}

func (this *LRUCache) Put(key int, value int)  {
    // 要先判断是否存在，再判断容量
    if node, exists := this.cache[key]; exists {
        node.val = value
        this.moveToHead(node)
        return
    }

    // 先判断是否有足够容量
    if this.occupied == this.capacity {
        // 移除队尾元素
        removed := this.removeTail()
        // 从map删除
        delete(this.cache, removed.key)
        this.occupied--
    }


    node := &Node{
        val: value,
        key: key,
    }

    // 添加到哈希表
    this.cache[key] = node
    this.occupied++

    // 把这个元素移到队首
    this.addToHead(node)
}
