package cdyleetcode

// *************************************************
// ❌❌ 循环里不能对循环外定义的变量用:=；for循环条件漏了一个等于
// *************************************************
func largestRectangleArea(heights []int) int {
    // stack保存的是index
    heights = append(heights, 0)
    stack := make([]int, 0, len(heights)) // 不可以是make([]int, len(heights)

    maxArea := 0

    for i, val := range heights {
        if len(stack) == 0 {
            stack = append(stack, i)
            continue
        }

        topIdx := stack[len(stack)-1]
        top := heights[topIdx]

        if val >= top {
            stack = append(stack, i)
        } else {
            for top > val {
                right := i
                var left int
                // 当前柱子是top的右边界。要寻找top的左边界
                for left := topIdx-1; left > 0 && heights[left] > top; left-- {}
                maxArea = max(maxArea, top*(right-left-1))
                // 弹栈
                stack = stack[:len(stack)-1]
                if len(stack) == 0 {
                    break
                }
                topIdx = stack[len(stack)-1]
                top = heights[topIdx]
            }
            // 压入当前index
            stack = append(stack, i)

        }
    }

    return maxArea
}

// 关键的思路：最终的矩形最大面积，必定是某个柱子的高度*（这个柱子右边界能到达的位置-这个柱子的左边界能到达的位置）
// 一次循环，遍历所有柱子。一个柱子的左右边界，都是第一个比自己矮的柱子。
// 用一个栈辅助处理，单调栈，保证栈内元素升序
// 栈为空 → 直接入栈
// 新柱子，如果比栈顶矮，说明是栈顶的右边界。开始循环进行弹栈处理，直到栈顶元素小于等于当前值：当前index就是栈顶元素的右边界，下一步就是找栈顶的左边界：往栈底部找，第一个小于当前栈顶的元素下标，就是左边界。接着弹出栈顶。计算当前矩形面积，与最大值比较。    最后循环退出，把当前的元素值压栈（如果是0就不压栈）。
//        如果比栈顶高或者一样，直接压栈

// *************************************************
// ❌❌ 修整的地方：卡了很久，这里千万注意不能var left int后，在for循环里用left:=...，不然后面的left就都是0！！！！
// 错误的地方：标准的单调栈解法应该是 $O(n)$，因为每个元素只进栈一次、出栈一次。但你的代码在出栈时，又嵌套了一个 for 循环去回溯找 left。其实 top 的左边界根本不需要找，它就躺在栈里！在单调递增栈中，当你准备弹出 stack[top] 时，由于栈是单调递增的，栈中位于 top 下方的那个元素，天生就是左边第一个比它矮的柱子。
// *************************************************
func largestRectangleArea(heights []int) int {
    // stack保存的是index
    heights = append(heights, 0)
    stack := make([]int, 0, len(heights)) // 不可以是make([]int, len(heights)

    maxArea := 0

    for i, val := range heights {
        if len(stack) == 0 {
            stack = append(stack, i)
            continue
        }

        topIdx := stack[len(stack)-1]
        top := heights[topIdx]

        if val >= top {
            stack = append(stack, i)
        } else {
            for top > val {
                right := i
                var left int
                // 当前柱子是top的右边界。要寻找top的左边界
                for left = topIdx-1; left >= 0 && heights[left] >= top; left-- {} // 卡了很久，这里千万注意不能left:=...，不然后面的left就都是0！！！！
                maxArea = max(maxArea, top*(right-left-1))
                // 弹栈
                stack = stack[:len(stack)-1]
                if len(stack) == 0 {
                    break
                }
                topIdx = stack[len(stack)-1]
                top = heights[topIdx]
            }
            // 压入当前index
            stack = append(stack, i)

        }
    }

    return maxArea
}

// 关键的思路：最终的矩形最大面积，必定是某个柱子的高度*（这个柱子右边界能到达的位置-这个柱子的左边界能到达的位置）
// 一次循环，遍历所有柱子。一个柱子的左右边界，都是第一个比自己矮的柱子。
// 用一个栈辅助处理，单调栈，保证栈内元素升序
// 栈为空 → 直接入栈
// 新柱子，如果比栈顶矮，说明是栈顶的右边界。开始循环进行弹栈处理，直到栈顶元素小于等于当前值：当前index就是栈顶元素的右边界，下一步就是找栈顶的左边界：往栈底部找，第一个小于当前栈顶的元素下标，就是左边界。接着弹出栈顶。计算当前矩形面积，与最大值比较。    最后循环退出，把当前的元素值压栈（如果是0就不压栈）。
//        如果比栈顶高或者一样，直接压栈

// *************************************************
// 修复寻找左边界的逻辑。左边界可以直接是栈顶的前一个元素下标（注意不是topIdx-1)。就算可能会1,5,5,6，往前处理到第一个5的时候，也可以正确计算出以5为高度的矩形的最大面积
// *************************************************
func largestRectangleArea(heights []int) int {
    // stack保存的是index
    heights = append(heights, 0)
    stack := make([]int, 0, len(heights)) // 不可以是make([]int, len(heights)

    maxArea := 0

    for i, val := range heights {
        if len(stack) == 0 {
            stack = append(stack, i)
            continue
        }

        topIdx := stack[len(stack)-1]
        top := heights[topIdx]

        if val >= top {
            stack = append(stack, i)
        } else {
            for top > val {
                right := i
                left := -1
                if len(stack) > 1 { // 不是>0
                    left = stack[len(stack)-2] // 不是topIdx-1！
                }
                maxArea = max(maxArea, top*(right-left-1))
                // 弹栈
                stack = stack[:len(stack)-1]
                if len(stack) == 0 {
                    break
                }
                topIdx = stack[len(stack)-1]
                top = heights[topIdx]
            }
            // 压入当前index
            stack = append(stack, i)

        }
    }

    return maxArea
}

// 关键的思路：最终的矩形最大面积，必定是某个柱子的高度*（这个柱子右边界能到达的位置-这个柱子的左边界能到达的位置）
// 一次循环，遍历所有柱子。一个柱子的左右边界，都是第一个比自己矮的柱子。
// 用一个栈辅助处理，单调栈，保证栈内元素升序
// 栈为空 → 直接入栈
// 新柱子，如果比栈顶矮，说明是栈顶的右边界。开始循环进行弹栈处理，直到栈顶元素小于等于当前值：当前index就是栈顶元素的右边界，下一步就是找栈顶的左边界：往栈底部找，第一个小于当前栈顶的元素下标，就是左边界。接着弹出栈顶。计算当前矩形面积，与最大值比较。    最后循环退出，把当前的元素值压栈（如果是0就不压栈）。
//        如果比栈顶高或者一样，直接压栈

// *************************************************
// 优化了一下if-else，反正都要append，没必要用太多if-else
// *************************************************
func largestRectangleArea(heights []int) int {
    // 在末尾加一个高度为0的柱子，保证最后能清空栈
    heights = append(heights, 0)
    // stack保存的是index
    stack := make([]int, 0, len(heights)) // 不可以是make([]int, len(heights)

    maxArea := 0

    for i, val := range heights {
        for len(stack) > 0 && heights[stack[len(stack)-1]] > val {
            right := i
            left := -1
            height := heights[stack[len(stack)-1]]
            if len(stack) > 1 { // 不是>0
                left = stack[len(stack)-2] // 不是topIdx-1！
            }
            maxArea = max(maxArea, height*(right-left-1))
            // 弹栈
            stack = stack[:len(stack)-1]
        }
        // 压入当前index
        stack = append(stack, i)
    }

    return maxArea
}

// 关键的思路：最终的矩形最大面积，必定是某个柱子的高度*（这个柱子右边界能到达的位置-这个柱子的左边界能到达的位置）
// 一次循环，遍历所有柱子。一个柱子的左右边界，都是第一个比自己矮的柱子。
// 用一个栈辅助处理，单调栈，保证栈内元素升序
// 栈为空 → 直接入栈
// 新柱子，如果比栈顶矮，说明是栈顶的右边界。开始循环进行弹栈处理，直到栈顶元素小于等于当前值：当前index就是栈顶元素的右边界，下一步就是找栈顶的左边界：往栈底部找，第一个小于当前栈顶的元素下标，就是左边界。接着弹出栈顶。计算当前矩形面积，与最大值比较。    最后循环退出，把当前的元素值压栈（如果是0就不压栈）。
//        如果比栈顶高或者一样，直接压栈