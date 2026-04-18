package cdyleetcode

// *************************************************
// ❌❌ dp公式错了
// *************************************************
func canPartition(nums []int) bool {
    target := 0
    for _, v := range nums {
        target = target + v
    }

    if target % 2 != 0 { // 奇数
        return false
    }

    target = target / 2

    dp := make([]int, len(nums))
    if nums[0] <= target {
        dp[0] = nums[0]
    }

    for i:=1; i<len(nums); i++ {
        if dp[i-1]+nums[i] <= target {
            dp[i] = max(dp[i-1], dp[i-1]+nums[i])
            if dp[i] == target {
                return true
            }
        }
    }

    return false
}

// 1 5 5 11 一共22 需要找两个和为11的子集，true
// 1 5 5 7 一共18 需要找两个和为9的子集，false
// 和为奇数，直接false

// 也就是说问题转化为，nums中是否存在一个和为sum/2的子数组。这就是背包问题了。我们假设背包总容量为sum/2，dp[i]表示到第i个下标为止时，背包中最多可以放入的重量。如果dp[i]==sum/2，就说明找到了。

// dp[i] = max(dp[i-1], dp[i-1]+nums[i]) （dp[i-1]+nums[i]必须小于等于sum/2)

// ---- 下面的是我自己想的，遇到的问题在于，需要找nums中和为target的任意子数组。麻烦了，而且不会

// 从最大值开始，最大值必然要在一个子集里，如果最大值大于half，就是false
// 如果最大值等于half，就是true
// 如果最大值小于half，就要在前面的数组里，找和为half-max且长度小于前面数组长度的子集

// *************************************************
// ❌❌ 数组越界，写[j-num]要讨论大小
// *************************************************
func canPartition(nums []int) bool {
    target := 0
    for _, v := range nums {
        target = target + v
    }

    if target % 2 != 0 { // 奇数
        return false
    }

    target = target / 2

    dp := make([][]bool, len(nums)+1) // dp[i][j]表示从前i个数里，能否凑出和j
    for i, _ := range dp {
        dp[i] = make([]bool, target+1)
    }

    dp[0][0] = true

    for i:=1; i<=len(nums); i++ {
        for j:=0; j<=target; j++ {
            num := nums[i]
            dp[i][j] = dp[i-1][j] || dp[i-1][j-num]
        }
    }

    return dp[len(nums)][target]
}

// 1 5 5 11 一共22 需要找两个和为11的子集，true
// 1 5 5 7 一共18 需要找两个和为9的子集，false
// 和为奇数，直接false

// 也就是说问题转化为，nums中是否存在一个和为sum/2的子数组。这就是背包问题了。

// dp[i][j]表示从前i个数里，能否凑出和j，最终结果就是dp[len(nums)][sum/2]
// dp[i+1][j] = dp[i][j] || dp[i][j-num] 其中num是nums[i+1] 也就是不选当前数字，就保存上一轮结果；如果选当前数字，上一轮就应该凑出j-num

// *************************************************
// 修改了dp表达式，直接看的题解。背包问题要用二维dp
// *************************************************
func canPartition(nums []int) bool {
    target := 0
    for _, v := range nums {
        target = target + v
    }

    if target % 2 != 0 { // 奇数
        return false
    }

    target = target / 2

    dp := make([][]bool, len(nums)+1) // dp[i][j]表示从前i个数里，能否凑出和j
    for i, _ := range dp {
        dp[i] = make([]bool, target+1)
    }

    dp[0][0] = true

    for i:=1; i<=len(nums); i++ {
        for j:=0; j<=target; j++ {
            num := nums[i-1]
            if j < num {
                // 要凑的目标和，比当前元素都小，说明当前元素肯定用不了
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = dp[i-1][j] || dp[i-1][j-num]
            }
        }
    }

    return dp[len(nums)][target]
}

// 1 5 5 11 一共22 需要找两个和为11的子集，true
// 1 5 5 7 一共18 需要找两个和为9的子集，false
// 和为奇数，直接false

// 也就是说问题转化为，nums中是否存在一个和为sum/2的子数组。这就是背包问题了。

// dp[i][j]表示从前i个数里，能否凑出和j，最终结果就是dp[len(nums)][sum/2]
// dp[i+1][j] = dp[i][j] || dp[i][j-num] 其中num是nums[i+1] 也就是不选当前数字，就保存上一轮结果；如果选当前数字，上一轮就应该凑出j-num