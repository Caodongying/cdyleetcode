package cdyleetcode

// *************************************************
// ❌❌ 判断条件写错了，错误地把当前温度和栈顶元素值比较。应该和栈顶元素对应的温度比较
// *************************************************
func dailyTemperatures(temperatures []int) []int {
    stack := make([]int, len(temperatures))
    result := make([]int, len(temperatures))
    for i, v := range temperatures {
        if len(stack) == 0 {
            stack = append(stack, i)
            continue
        }

        for len(stack) > 0 && v > stack[len(stack)-1] {
            // 弹栈
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            // 处理弹出的元素
            result[top] = i - top
        }

        // 压栈当前元素
        stack = append(stack, i)
    }

    // 处理栈内剩余元素
    for i:=0; i<len(stack); i++ {
        result[stack[i]] = 0
    }

    return result

}

// 我只能想到用O(n^2)的暴力循环
// 直接看了题解和B站视频，用单调栈。单调是因为：
// 栈顶元素始终是：还没算出下一个最大元素的数的下标
// 压栈的是index，而不是元素值本身
// [73,74,75,71,69,72,76,73]
// 73压栈(0) → 74大于73，弹栈处理73，入栈74（1）→75大于74，弹栈处理74，入栈75→入栈71→入栈69→72大于69和71，弹栈处理69和71，入栈72→弹栈处理72和75，入栈76→入栈73→最后对栈内剩余元素都赋0

// *************************************************
// 加上了温度值
// *************************************************
func dailyTemperatures(temperatures []int) []int {
    stack := make([]int, len(temperatures))
    result := make([]int, len(temperatures))
    for i, v := range temperatures {
        if len(stack) == 0 {
            stack = append(stack, i)
            continue
        }

        for len(stack) > 0 && v > temperatures[stack[len(stack)-1]] {
            // 弹栈
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            // 处理弹出的元素
            result[top] = i - top
        }

        // 压栈当前元素
        stack = append(stack, i)
    }

    // 处理栈内剩余元素
    for i:=0; i<len(stack); i++ {
        result[stack[i]] = 0
    }

    return result

}

// *************************************************
// 去掉一开始对栈空的判断，因为栈空直接append，而栈不空，最后也要append
// *************************************************
func dailyTemperatures(temperatures []int) []int {
    stack := make([]int, len(temperatures))
    result := make([]int, len(temperatures))
    for i, v := range temperatures {
        for len(stack) > 0 && v > temperatures[stack[len(stack)-1]] {
            // 弹栈
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            // 处理弹出的元素
            result[top] = i - top
        }

        // 压栈当前元素
        stack = append(stack, i)
    }

    // 处理栈内剩余元素
    for i:=0; i<len(stack); i++ {
        result[stack[i]] = 0
    }

    return result

}
