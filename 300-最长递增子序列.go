package cdyleetcode

func lengthOfLIS(nums []int) int {
    dp := make([]int, len(nums))
    dp[0] = 1
    for i:=1; i<len(nums); i++ {
        // 找前面的比nums[i]小的且dp值最大的数
        dp[i] = 1
        var j int
        var findSmaller bool = false
        for j=i-1; j>=0; j-- {
            if nums[j] < nums[i] {
                dp[i] = max(dp[i], 1+dp[j])
                findSmaller = true
            }
        }

        if findSmaller == false {
            dp[i] = 1
        }
    }

    // 找到dp的最大值
    maxDp := 1
    for i:=0; i<len(dp); i++ {
        maxDp = max(maxDp, dp[i])
    }

    return maxDp
}


// dp[i]表示以第i个元素为末尾元素的递增子序列的最大长度
// 往前找到最近的比nums[i]小的数，记下标为j，dp[i]就是1+dp[j]
// 如果找不到j，dp[i]就是1
// dp[0] = 1