package cdyleetcode

// *************************************************
// ❌❌ 最直观的解法
// *************************************************
func climbStairs(n int) int {
    if n == 1 {
        return 1
    }

    if n == 2 {
        return 2
    }

    return climbStairs(n-2) + climbStairs(n-1)
}

// 最后一步，可能走两个台阶，可能走一个台阶
// f(n) = f(n-2) + f(n-1)
// 然后初始化

// *************************************************
// ❌❌ resultMap在初始化的时候就要设置resultMap[1]和resultMap[2]。或者直接先return
// *************************************************
var resultMap map[int]int

func climbStairs(n int) int {
    if resultMap == nil {
        resultMap = make(map[int]int, n+1)
        resultMap[0] = 0
    }

    if n == 1 {
        resultMap[1] = 1
    } else if n == 2 {
        resultMap[2] = 2
    } else {
        resultMap[n] = resultMap[n-2] + resultMap[n-1]
    }

    return resultMap[n]
}

// 优化：避免重复计算

// *************************************************
// 全局map
// *************************************************
var resultMap map[int]int

func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }

    if resultMap == nil {
        resultMap = make(map[int]int, n+1)
        resultMap[0] = 0
        resultMap[1] = 1
        resultMap[2] = 2
    }

    // 从3到n，逐个计算，保存结果
    for i := 3; i <= n; i++ {
        resultMap[i] = resultMap[i-1] + resultMap[i-2]
    }

    return resultMap[n]
}

// 优化：避免重复计算

// *************************************************
// ❌❌ 因为leetcode内部编译器的缘故，过不了。用的数组
// *************************************************
var resultMap []int

func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }

    if resultMap == nil {
        resultMap = make([]int, n+1)
        resultMap[0] = 0
        resultMap[1] = 1
        resultMap[2] = 2
    }

    // 从3到n，逐个计算，保存结果
    for i := 3; i <= n; i++ {
        resultMap[i] = resultMap[i-1] + resultMap[i-2]
    }

    return resultMap[n]
}

// *************************************************
// 把全局的数组改为局部的，避免leetcode编译器的问题
// *************************************************
func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }


    resultMap := make([]int, n+1)
    resultMap[0] = 0
    resultMap[1] = 1
    resultMap[2] = 2


    // 从3到n，逐个计算，保存结果
    for i := 3; i <= n; i++ {
        resultMap[i] = resultMap[i-1] + resultMap[i-2]
    }

    return resultMap[n]
}


// *************************************************
// 直接用两个变量，不需要数组
// *************************************************
func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }

    prev2, prev1 := 1, 2


    // 从3到n，逐个计算，保存结果
    for i := 3; i <= n; i++ {
        current := prev2 + prev1
        prev2, prev1 = prev1, current
    }

    return prev1
}


// *************************************************
// 简化一下if-else
// *************************************************
func climbStairs(n int) int {
    if n <= 2 {
        return n
    }

    a, b := 1, 2

    // 从3到n，逐个计算，保存结果
    for i := 3; i <= n; i++ {
        a, b = b, a+b
    }

    return b
}

// 最后一步，可能走两个台阶，可能走一个台阶
// f(n) = f(n-2) + f(n-1)
// 然后初始化

// 优化：避免重复计算