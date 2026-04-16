package main

import "fmt"
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

// 最后一步，可能走两个台阶，可能走一个台阶
// f(n) = f(n-2) + f(n-1)
// 然后初始化

// 优化：避免重复计算


func main() {
   fmt.Println(climbStairs(4))
}
