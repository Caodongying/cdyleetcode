package cdyleetcode

// *************************************************
// ❌❌ 自己写的错误解法，本题不可以用贪心算法！不然[1,3,4,5]，amount=7就不对了
// *************************************************
import (
    "sort"
)

func coinChange(coins []int, amount int) int {
    result := -1

    // 对coins排序
    sort.Ints(coins)

    var helper func(coins []int, amount int) int
    helper = func(coins []int, amount int) int{
        curr := math.MaxInt64
        if amount == 0 || len(coins) == 0 {
            return 0
        }

        // 找当前可用的最大硬币
        var maxCoinIndex int
        for i := len(coins)-1; i >= 0; i-- {
            if coins[i] <= amount {
                maxCoinIndex = i
                break
            }
        }

        if maxCoinIndex == -1 {
            return 0
        }

        maxCoinCount := amount / coins[maxCoinIndex]

        for i:=maxCoinCount; i>=0; i-- {
            curr = min(curr, i+helper(coins[:maxCoinIndex], amount-i*coins[maxCoinIndex]))
        }

        return curr
    }

    result = helper(coins, amount)

    return result
}

// *************************************************
// ❌❌ 错误地写了remain:=amount-coin，应该是i-coin
// *************************************************
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    dp[0] = 0

    // 先全部初始化为-1
    for i := 1; i<len(dp); i++ {
        dp[i] = -1
    }

    // 从金额1开始计算，结果写入dp[i]
    for i := 1; i<=amount; i++ {
        minCount := math.MaxInt32
        for _, coin := range coins {
            if coin <= i {
                remain := amount - coin
                if dp[remain] != -1 {
                    minCount = min(minCount, 1+dp[remain])
                    dp[i] = minCount
                }
            }
        }
    }

    return dp[amount]
}

// 类比爬楼梯问题
// 比如对于[1,2,5]，amount = 11
// 其实就是三个方案：
// 1. 拿1个1元硬币，剩余10元，最终结果就是1+拿10元需要的最小硬币个数
// 2. 拿1个2元硬币，剩余9元，最终结果就是1+拿9元需要的最小硬币个数
// 3. 拿1个5元硬币，剩余6元，最终结果就是1+拿6元需要的最小硬币个数
// 当然硬币面值要小于当前的amount

// 动态规划，可以自底向上计算。就是从前往后遍历。dp表示在当前coins数组下，拿amount所需要的最少硬币个数

// *************************************************
// 修复了i-coin的问题
// *************************************************
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    dp[0] = 0

    // 先全部初始化为-1
    for i := 1; i<len(dp); i++ {
        dp[i] = -1
    }

    // 从金额1开始计算，结果写入dp[i]
    for i := 1; i<=amount; i++ {
        minCount := math.MaxInt32
        for _, coin := range coins {
            if coin <= i {
                remain := i - coin // 不是amount-coin
                if dp[remain] != -1 {
                    minCount = min(minCount, 1+dp[remain])
                    dp[i] = minCount
                }
            }
        }
    }

    return dp[amount]
}

// 类比爬楼梯问题
// 比如对于[1,2,5]，amount = 11
// 其实就是三个方案：
// 1. 拿1个1元硬币，剩余10元，最终结果就是1+拿10元需要的最小硬币个数
// 2. 拿1个2元硬币，剩余9元，最终结果就是1+拿9元需要的最小硬币个数
// 3. 拿1个5元硬币，剩余6元，最终结果就是1+拿6元需要的最小硬币个数
// 当然硬币面值要小于当前的amount

// 动态规划，可以自底向上计算。就是从前往后遍历。dp表示在当前coins数组下，拿amount所需要的最少硬币个数

// *************************************************
// 简化一下for
// *************************************************
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    dp[0] = 0

    // 从金额1开始计算，结果写入dp[i]
    for i := 1; i<=amount; i++ {
        dp[i] = -1
        minCount := math.MaxInt32
        for _, coin := range coins {
            if coin <= i && dp[i-coin] != -1 {
                minCount = min(minCount, 1+dp[i - coin])
                dp[i] = minCount
            }
        }
    }

    return dp[amount]
}

// 类比爬楼梯问题，这也是完全背包问题
// 比如对于[1,2,5]，amount = 11
// 其实就是三个方案：
// 1. 拿1个1元硬币，剩余10元，最终结果就是1+拿10元需要的最小硬币个数
// 2. 拿1个2元硬币，剩余9元，最终结果就是1+拿9元需要的最小硬币个数
// 3. 拿1个5元硬币，剩余6元，最终结果就是1+拿6元需要的最小硬币个数
// 当然硬币面值要小于当前的amount

// 动态规划，可以自底向上计算。就是从前往后遍历。dp表示在当前coins数组下，拿amount所需要的最少硬币个数