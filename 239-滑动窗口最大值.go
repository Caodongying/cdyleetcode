package cdyleetcode

// *************************************************
// 代码随想录的解法 - 单调队列
// ❌错误解法，因为要先单独处理第一个滑动窗口，这样之后滑动窗口每移动一次，只要检查新加入的元素
// 现在这个做法是对每个滑动窗口的所有元素都进行检测，一个元素被处理很多次
// *************************************************
type Deque struct {
    deque []int
}

// 获取deque的队首值
func (dq *Deque) Get() (bool, int) {
    if len(dq.deque) == 0 {
        return false, 0
    } else {
        return true, dq.deque[0]
    }
}

// 向deque队尾添加元素
func (dq *Deque) Push(value int){
    dq.deque = append(dq.deque, value)
}

// 移除deque队首元素
func (dq *Deque) Pop(){
    if len(dq.deque) > 0 {
        dq.deque = dq.deque[1:]
    }
}

func maxSlidingWindow(nums []int, k int) []int {
    var deque Deque
    result := make([]int, len(nums)-k+1)
    for i:=0; i <= len(nums)-k; i++ {
        exist, front := deque.Get()
        // 先进行pop检测
        if exist && i-1>= 0 && front == nums[i-1] {
            // 上一个滑动窗口的最大值就是其最左边的值，而且已经移出去了
            deque.Pop()
        }

        // 再push滑动窗口内的数值
        for j:=i; j<i+k; j++ {
            deque.Push(nums[j])
            // 检查当前加入的值是否大于等于deque内在其前面的所有值
            // 因为这个特性，deque是单调递减数列，所以只要与队首值比较
            exist, front = deque.Get()
            if exist && nums[j] >= front {
                for k:=0; k<j; k++ {
                    deque.Pop()
                }
            }
        }

        // 保存当前滑动数组的最大值
        _, result[i] = deque.Get()

    }

    return result
}

// 维护一个双端队列deque(double-end queue)
// 队首pop，队尾push，也就是右进左出
// 这个deque的队首为当前滑动窗口的最大值
// 学问就在pop()和push()规则上
// 最high-level的过程：在滑动窗口移动过程中，把每个元素按照规则push进deque，每一个push后，如果当前元素大于deque里前面的所有元素，就把那些元素都pop出去（因为这些值对于当前滑动窗口而言，必然不是最大值了）。也就是说，push是必然的，pop是optional的。此外，如果被移出滑动窗口的数字，等于deque的队首值，就说明上一个滑动窗口的最大值在最左端，移出去了，相应的在deque里也要移除
// 也就是说，在滑动窗口移动的过程中，要进行两种检测
// 在针对当前滑动窗口的deque维护过程结束后，返回deque首部，就是当前滑动窗口的最大值

// 思索的例子： 5 4 3 2 1 （k=3) 和 4 5 3 2 1 (k=3)

// 此方法的时间复杂度：O(n-k+1)，也就是O(n)



// *************************************************
// 反复修改后的版本，语法、小错误
// *************************************************
type Deque struct {
    deque []int
}

// 获取deque的队首值
func (dq *Deque) Get() (bool, int) {
    if len(dq.deque) == 0 {
        return false, 0
    } else {
        return true, dq.deque[0]
    }
}

// 向deque队尾添加元素
func (dq *Deque) Push(value int){
    dq.deque = append(dq.deque, value)
}

// 移除deque队首元素
func (dq *Deque) Pop(){
    if len(dq.deque) > 0 {
        dq.deque = dq.deque[1:]
    }
}

func maxSlidingWindow(nums []int, k int) []int {
    var deque Deque
    result := make([]int, len(nums)-k+1)
    // 先处理第一个滑动窗口
    // push滑动窗口内的数值
    for i:=0; i<k; i++ {
        deque.Push(nums[i])
        // 检查当前加入的值是否大于等于deque内在其前面的所有值
        // 因为这个特性，deque是单调递减数列，所以只要与队首值比较
        exist, front := deque.Get()
        if exist && nums[i] >= front { // 这一部分不就是直接清除deque了吗，或者应该先检测这个
            for { // Golang没有while
                exist, v := deque.Get()
                if !exist || v > nums[i]{
                    break
                }
                deque.Pop()
            }

            // 不可以是下面的代码，因为一直在删除队首，deque.deque[j] out of range
            // for j:=0; deque.deque[j]<=nums[i]; j++ {
            //     deque.Pop()
            // }
        }
    }

    _, result[0] = deque.Get()

    // 接下来处理后面的
    for i:=1; i <= len(nums)-k; i++ {
        exist, front := deque.Get()
        // 先进行pop检测
        if exist && i-1>= 0 && front == nums[i-1] {
            // 上一个滑动窗口的最大值就是其最左边的值，而且已经移出去了
            deque.Pop()
        }

        // 再push新滑动窗口里的最后一个新数值
        deque.Push(nums[i+k-1])
        // 检查当前加入的值是否大于等于deque内在其前面的所有值
        // 因为这个特性，deque是单调递减数列，所以只要与队首值比较
        exist, front = deque.Get()
        if exist && nums[i+k-1] >= front {
            for { // Golang没有while
                exist, v := deque.Get()
                if !exist || v > nums[i+k-1]{
                    break
                }
                deque.Pop()
            }

        }

        // 保存当前滑动数组的最大值
        _, result[i] = deque.Get()

    }

    return result
}


// *************************************************
// 只比较了新元素和队首，导致队列不是单调的
// *************************************************
type Deque struct {
    deque []int
}

// 获取deque的队首值
func (dq *Deque) Get() (bool, int) {
    if len(dq.deque) == 0 {
        return false, 0
    } else {
        return true, dq.deque[0]
    }
}

// 向deque队尾添加元素
func (dq *Deque) Push(value int){
    dq.deque = append(dq.deque, value)
}

// 移除deque队首元素
func (dq *Deque) Pop(){
    if len(dq.deque) > 0 {
        dq.deque = dq.deque[1:]
    }
}

func maxSlidingWindow(nums []int, k int) []int {
    var deque Deque
    result := make([]int, len(nums)-k+1)
    // 先处理第一个滑动窗口
    // push滑动窗口内的数值
    for i:=0; i<k; i++ {
        // 检查当前加入的值是否大于等于deque内在其前面的所有值
        // 因为这个特性，deque是单调递减数列，所以只要与队首值比较
        exist, front := deque.Get()
        if exist && nums[i] >= front { // 先检测再Push，方便进行deque前面元素的删除
            for { // Golang没有while
                exist, _ := deque.Get()
                if !exist {
                    break
                }
                deque.Pop()
            }

            // 不可以是下面的代码，因为一直在删除队首，deque.deque[j] out of range
            // for j:=0; deque.deque[j]<=nums[i]; j++ {
            //     deque.Pop()
            // }
        }

        deque.Push(nums[i])
    }

    _, result[0] = deque.Get()

    // 接下来处理后面的
    for i:=1; i <= len(nums)-k; i++ {
        exist, front := deque.Get()
        // 先进行pop检测
        if exist && i-1>= 0 && front == nums[i-1] {
            // 上一个滑动窗口的最大值就是其最左边的值，而且已经移出去了
            deque.Pop()
        }

        // 检查当前加入的值是否大于等于deque内在其前面的所有值
        // 因为这个特性，deque是单调递减数列，所以只要与队首值比较
        exist, front = deque.Get()
        if exist && nums[i+k-1] >= front {
            for { // Golang没有while
                exist, _ := deque.Get()
                if !exist {
                    break
                }
                deque.Pop()
            }

        }

        // 再push新滑动窗口里的最后一个新数值
        deque.Push(nums[i+k-1])

        // 保存当前滑动数组的最大值
        _, result[i] = deque.Get()

    }

    return result
}

// 维护一个双端队列deque(double-end queue)
// 队首pop，队尾push，也就是右进左出
// 这个deque的队首为当前滑动窗口的最大值
// 学问就在pop()和push()规则上
// 最high-level的过程：在滑动窗口移动过程中，把每个元素按照规则push进deque，每一个push后，如果当前元素大于deque里前面的所有元素，就把那些元素都pop出去（因为这些值对于当前滑动窗口而言，必然不是最大值了）。也就是说，push是必然的，pop是optional的。此外，如果被移出滑动窗口的数字，等于deque的队首值，就说明上一个滑动窗口的最大值在最左端，移出去了，相应的在deque里也要移除
// 也就是说，在滑动窗口移动的过程中，要进行两种检测
// 在针对当前滑动窗口的deque维护过程结束后，返回deque首部，就是当前滑动窗口的最大值

// 思索的例子： 5 4 3 2 1 （k=3) 和 4 5 3 2 1 (k=3)

// 此方法的时间复杂度：O(n-k+1)，也就是O(n)

// *************************************************
// 自己改了一下
// *************************************************
type Deque struct {
    deque []int
}

// 获取deque的队首值
func (dq *Deque) GetFront() (bool, int) {
    if len(dq.deque) == 0 {
        return false, 0
    } else {
        return true, dq.deque[0]
    }
}

// 获取deque的队尾值
func (dq *Deque) GetBack() (bool, int) {
    if len(dq.deque) == 0 {
        return false, 0
    } else {
        return true, dq.deque[len(dq.deque)-1]
    }
}

// 向deque队尾添加元素
func (dq *Deque) Push(value int){
    dq.deque = append(dq.deque, value)
}

// 移除deque队首元素
func (dq *Deque) Pop(){
    if len(dq.deque) > 0 {
        dq.deque = dq.deque[1:]
    }
}

func maxSlidingWindow(nums []int, k int) []int {
    var deque Deque
    result := make([]int, len(nums)-k+1)

    // 先处理第一个滑动窗口
    // push滑动窗口内的数值
    for i:=0; i<k; i++ {
        // 检查当前加入的值是否大于等于deque内在其前面的部分值
        // 如果这样，就要删除这些值，因为这些值必然之后没用了
        // 因为这个特性，deque是单调递减数列，所以只要与队尾值比较
        for { // Golang没有while
            exist, back := deque.GetBack()
            if !exist || back >= nums[i]{
                break
            }
            deque.deque = deque.deque[:len(deque.deque)-1]
        }

        // 不可以是下面的代码，因为一直在删除队首，deque.deque[j] out of range
        // for j:=0; deque.deque[j]<=nums[i]; j++ {
        //     deque.Pop()
        // }


        deque.Push(nums[i])
    }

    _, result[0] = deque.GetFront()

    // 接下来处理后面的
    for i:=1; i <= len(nums)-k; i++ {
        exist, front := deque.GetFront()
        // 先进行pop检测
        if exist && i-1>= 0 && front == nums[i-1] {
            // 上一个滑动窗口的最大值就是其最左边的值，而且已经移出去了
            deque.Pop()
        }

        // 检查当前加入的值是否大于等于deque内在其前面的部分值

        for { // Golang没有while
            exist, back := deque.GetBack()
            if !exist || back >= nums[i+k-1]{
                break
            }
            deque.deque = deque.deque[:len(deque.deque)-1]
        }

        // 再push新滑动窗口里的最后一个新数值
        deque.Push(nums[i+k-1])

        // 保存当前滑动数组的最大值
        _, result[i] = deque.GetFront()

    }

    return result
}

// 维护一个双端队列deque(double-end queue)
// 队首pop，队尾push，也就是右进左出
// 这个deque的队首为当前滑动窗口的最大值
// 学问就在pop()和push()规则上
// 最high-level的过程：在滑动窗口移动过程中，把每个元素按照规则push进deque，每一个push后，如果当前元素大于deque里前面的部分元素，就把那些元素都pop出去（因为这些值对于当前滑动窗口而言，必然不是最大值了）。也就是说，push是必然的，pop是optional的。此外，如果被移出滑动窗口的数字，等于deque的队首值，就说明上一个滑动窗口的最大值在最左端，移出去了，相应的在deque里也要移除
// 也就是说，在滑动窗口移动的过程中，要进行两种检测
// 在针对当前滑动窗口的deque维护过程结束后，返回deque首部，就是当前滑动窗口的最大值

// 思索的例子： 5 4 3 2 1 （k=3) 和 4 5 3 2 1 (k=3)

// 此方法的时间复杂度：O(n-k+1)，也就是O(n)

// *************************************************
// 进一步优化
// *************************************************

type Deque struct {
    deque []int
}

func NewDeque() *Deque {
    return &Deque{
        deque: make([]int, 0),
    }
}

// 获取deque的队首值
func (dq *Deque) GetFront() (bool, int) {
    if len(dq.deque) == 0 {
        return false, 0
    } else {
        return true, dq.deque[0]
    }
}

// 获取deque的队尾值
func (dq *Deque) GetBack() (bool, int) {
    if len(dq.deque) == 0 {
        return false, 0
    } else {
        return true, dq.deque[len(dq.deque)-1]
    }
}

// 向deque队尾添加元素
func (dq *Deque) Push(value int){
    // 检查当前加入的值是否大于等于deque内在其前面的部分值
    // 如果这样，就要删除这些值，因为这些值必然之后没用了
    // 因为这个特性，deque是单调递减数列，所以只要与队尾值比较
    for { // Golang没有while
        exist, back := dq.GetBack()
        if !exist || back >= value{
            break
        }
        dq.deque = dq.deque[:len(dq.deque)-1]
    }

    // 不可以是下面的代码，因为一直在删除队首，deque.deque[j] out of range
    // for j:=0; deque.deque[j]<=nvalue; j++ {
    //     deque.Pop()
    // }

    dq.deque = append(dq.deque, value)
}

// 移除deque队首元素
func (dq *Deque) Pop(){
    if len(dq.deque) > 0 {
        dq.deque = dq.deque[1:]
    }
}

func maxSlidingWindow(nums []int, k int) []int {
    deque := NewDeque()
    result := make([]int, 0)

    // 先处理第一个滑动窗口
    // push滑动窗口内的数值
    for i:=0; i<k; i++ {
        deque.Push(nums[i])
    }

    _, front := deque.GetFront()
    result = append(result, front)

    // 接下来处理后面的
    for i:=1; i <= len(nums)-k; i++ {
        exist, front := deque.GetFront()
        // 先进行pop检测
        if exist && i-1>= 0 && front == nums[i-1] {
            // 上一个滑动窗口的最大值就是其最左边的值，而且已经移出去了
            deque.Pop()
        }

        deque.Push(nums[i+k-1])

        // 保存当前滑动数组的最大值
        _, front = deque.GetFront()
        result = append(result, front)


    }

    return result
}

// 维护一个双端队列deque(double-end queue)
// 队首pop，队尾push，也就是右进左出
// 这个deque的队首为当前滑动窗口的最大值
// 学问就在pop()和push()规则上
// 最high-level的过程：在滑动窗口移动过程中，把每个元素按照规则push进deque，每一个push后，如果当前元素大于deque里前面的部分元素，就把那些元素都pop出去（因为这些值对于当前滑动窗口而言，必然不是最大值了）。也就是说，push是必然的，pop是optional的。此外，如果被移出滑动窗口的数字，等于deque的队首值，就说明上一个滑动窗口的最大值在最左端，移出去了，相应的在deque里也要移除
// 也就是说，在滑动窗口移动的过程中，要进行两种检测
// 在针对当前滑动窗口的deque维护过程结束后，返回deque首部，就是当前滑动窗口的最大值

// 思索的例子： 5 4 3 2 1 （k=3) 和 4 5 3 2 1 (k=3)

// 此方法的时间复杂度：O(n-k+1)，也就是O(n)