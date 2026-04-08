package cdyleetcode

// *************************************************
// 自己写的，保存最小值的index。弹栈的时候如果弹出的是最小值，就手动搜索一下新的最小值。
// *************************************************
type MinStack struct {
    stack []int
    minIdx int
}


func Constructor() MinStack {
    minStack := MinStack{}
    minStack.stack = make([]int, 0)
    return minStack
}


func (this *MinStack) Push(val int)  {
    this.stack = append(this.stack, val)
    // 看是否需要更新最小值索引
    if val < this.stack[this.minIdx] {
        this.minIdx = len(this.stack)-1
    }
}


func (this *MinStack) Pop() {
    this.stack = this.stack[:len(this.stack)-1]
    if this.minIdx == len(this.stack) {
        // 弹出的元素是原本的最小值，需要重新构建最小值
        min := 0
        for i:=1; i<len(this.stack); i++ {
            if this.stack[i] < this.stack[min] {
                min = i
            }
        }
        this.minIdx = min
    }
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
    return this.stack[this.minIdx]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

// *************************************************
// ❌❌ 看了题解后用了diff，但是忘记在栈空的时候，处理完毕然后return了
// *************************************************
type MinStack struct {
    stack []int
    min int
}


func Constructor() MinStack {
    minStack := MinStack{}
    minStack.stack = make([]int, 0)
    return minStack
}


func (this *MinStack) Push(val int)  {
    // 要单独处理栈空的情况
    // if len(this.stack) == 0 {
    //     this.min = val
    //     this.stack = append(this.stack, 0)
    // }

    diff := val - this.min
    this.stack = append(this.stack, diff)
    // 看是否需要更新最小值索引
    if val < this.min {
        this.min = val // 此时diff是负数
    }
}


func (this *MinStack) Pop() {
    topDiff := this.stack[len(this.stack)-1]
    this.stack = this.stack[:len(this.stack)-1]

    // 判断是否要更新min
    if topDiff < 0 {
        this.min = this.min - topDiff
    }
}


func (this *MinStack) Top() int {
   topDiff := this.stack[len(this.stack)-1]
   if topDiff >= 0 {
        return this.min + topDiff
   } else {
        return this.min
   }
}


func (this *MinStack) GetMin() int {
    return this.min
}

// stack不存储值本身，而是存储一个特殊的值diff
// 用diff来动态保存最小值信息
// 如果diff大于0，说明当前值大于栈的最小值。
// 如果diff小于0，说明当前值小于栈的最小值。
// push的时候，先更新diff（curr-min)，然后更新min（if需要更新）
// 这样，pop的时候，从diff还原原数值的时候，就看：如果diff大于0，那min+diff就是原数值，min不用改变；如果diff等于0，说明额外还有一个一样的值是最小值；如果diff<0，说明min已经被更新过了（因为diff<0才更新min），min就是当前值。然后恢复新的最小值，就是用min-diff（因为diff=curr-旧的min，而现在curr就是min，所以旧的min是curr-diff，也就是min-diff）

// *************************************************
// 修正了栈空时的return问题
// *************************************************
type MinStack struct {
    stack []int
    min int
}


func Constructor() MinStack {
    minStack := MinStack{}
    minStack.stack = make([]int, 0)
    return minStack
}


func (this *MinStack) Push(val int)  {
    // 要单独处理栈空的情况
    if len(this.stack) == 0 {
        this.min = val
        this.stack = append(this.stack, 0)
        return
    }

    diff := val - this.min
    this.stack = append(this.stack, diff)
    // 看是否需要更新最小值索引
    if val < this.min {
        this.min = val // 此时diff是负数
    }
}


func (this *MinStack) Pop() {
    topDiff := this.stack[len(this.stack)-1]
    this.stack = this.stack[:len(this.stack)-1]

    // 判断是否要更新min
    if topDiff < 0 {
        this.min = this.min - topDiff
    }
}


func (this *MinStack) Top() int {
   topDiff := this.stack[len(this.stack)-1]
   if topDiff >= 0 {
        return this.min + topDiff
   } else {
        return this.min
   }
}


func (this *MinStack) GetMin() int {
    return this.min
}

// stack不存储值本身，而是存储一个特殊的值diff
// 用diff来动态保存最小值信息
// 如果diff大于0，说明当前值大于栈的最小值。
// 如果diff小于0，说明当前值小于栈的最小值。
// push的时候，先更新diff（curr-min)，然后更新min（if需要更新）
// 这样，pop的时候，从diff还原原数值的时候，就看：如果diff大于0，那min+diff就是原数值，min不用改变；如果diff等于0，说明额外还有一个一样的值是最小值；如果diff<0，说明min已经被更新过了（因为diff<0才更新min），min就是当前值。然后恢复新的最小值，就是用min-diff（因为diff=curr-旧的min，而现在curr就是min，所以旧的min是curr-diff，也就是min-diff）

