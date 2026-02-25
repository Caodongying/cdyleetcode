package cdyleetcode

func rob(nums []int) int {
    // var dp [len(nums)]int 长度必须在编译时就确定

    dp := make([]int, len(nums))

    if len(nums) == 1 {
        return nums[0]
    }

    dp[0] = nums[0]
    dp[1] = max(nums[0], nums[1])

    for i:=2; i<len(nums); i++ {
        dp[i] = max(nums[i]+dp[i-2], dp[i-1])
    }

    return dp[len(nums)-1]
}

// 当前房屋偷不偷取决于前一个和前2个房屋是够被偷
// 也就是当前状态和前面状态存在依赖关系
// 写出这个依赖关系
// 用dp[i]表示从nums[0]到nums[i-1]最多能偷多少