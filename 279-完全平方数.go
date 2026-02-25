package cdyleetcode

import (
    "math"
)

func numSquares(n int) int {
    dp := make([]int, n+1) // dp[n]表示和为n的完全平方数的最小数量
    dp[0] = 0

    for i:=1; i<len(dp); i++ {
        dp[i] = math.MaxInt32
    }

    // 计算从1到n的所有值的dp
    for i:=1; i<=n; i++ {
        // 看选用哪一个平方数，因为dp[i]一定是某个平方数x加dp[n-x]，只是要选最小的
        for j:=1; j*j<=i; j++ {
            dp[i] = min(dp[i], 1+dp[i-j*j])
        }
    }

    return dp[n]
}

// func min(x int, y int) int {
//     if x <= y {
//         return x
//     }

//     return y
// }
// ---不对
// 构建切片squareSums[0, 1, 4, 9, 16, 25, 36, 49, 64, 81...] (100个数)
// dp[i]表示和为i-1的完全平方数的最少数量
// 如果n在构建切片squareSums中，返回1
// 不然找到squareSums中比n小的最大数x，返回1+dp[n-x]