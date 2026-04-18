package cdyleetcode

// *************************************************
// ❌❌ 状态污染了：在计算 curMin 的时候，用到的 curMax 已经不是上一轮的旧值了，
// 而是刚刚被更新过的新 curMax。这就导致了状态污染
// *************************************************
func maxProduct(nums []int) int {
    curMax, curMin, globalMax := nums[0], nums[0], nums[0]

    for i := 1; i < len(nums); i++ {
        curMax = max(nums[i], nums[i]*curMax, nums[i]*curMin)
        curMin = min(nums[i], nums[i]*curMax, nums[i]*curMin)
        globalMax = max(globalMax, curMax)
    }
    return globalMax
}

// 本题涉及对正负数的讨论
// 一次遍历，找以nums[i]结尾的子数组乘积的最大值和最小值
// 然后保持更新全局max

// 遍历到数字num时，我们在心里做了一个抉择——我是把num单独作为一个新子数组的起点（候选值num），还是把它接到之前那个子数组的屁股后面（候选值curMax * num或curMin * num）。如果之前的子数组乘积对现在有利就接上，如果接上反而变小了（比如之前乘积是负的、现在是正的），我们就果断切断联系、从num重新开始。所以子数组的起点是在每一步的max和min比较中动态决定的，并不是写死在索引0上的。

// 错误————————
// dp[i][j]表示从下标i到下标j的乘积
// i优先，先计算i=0时，所有的dp[0][j]
// 再递增i
// 不对啊，这是暴力解法

// *************************************************
// 动态规划，分min和max，直接看的题解
// *************************************************
func maxProduct(nums []int) int {
    curMax, curMin, globalMax := nums[0], nums[0], nums[0]

    for i := 1; i < len(nums); i++ {
        tempMax, tempMin := curMax, curMin
        curMax = max(nums[i], nums[i]*tempMax, nums[i]*tempMin)
        curMin = min(nums[i], nums[i]*tempMax, nums[i]*tempMin)
        globalMax = max(globalMax, curMax)
    }
    return globalMax
}

// 本题涉及对正负数的讨论
// 一次遍历，找以nums[i]结尾的子数组乘积的最大值和最小值
// 然后保持更新全局max

// 遍历到数字num时，我们在心里做了一个抉择——我是把num单独作为一个新子数组的起点（候选值num），还是把它接到之前那个子数组的屁股后面（候选值curMax * num或curMin * num）。如果之前的子数组乘积对现在有利就接上，如果接上反而变小了（比如之前乘积是负的、现在是正的），我们就果断切断联系、从num重新开始。所以子数组的起点是在每一步的max和min比较中动态决定的，并不是写死在索引0上的。
