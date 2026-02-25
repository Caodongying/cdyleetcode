package cdyleetcode

// *************************************************
// 前缀和，不要要维护数组。记录之前的最小前缀和，因为数组里可能存在负数。
// *************************************************

func maxSubArray(nums []int) int {
    minPreviousPrefixSum := nums[0] // 当前元素之前的最小前缀和
    prefixSum := nums[0]
    maxSum := nums[0]
    for i:=1; i<len(nums); i++ {
        // 计算当前前缀和
        prefixSum += nums[i]
        if prefixSum > maxSum {
            maxSum = prefixSum
        }
        if prefixSum - minPreviousPrefixSum > maxSum {
            maxSum = prefixSum - minPreviousPrefixSum
        }

        // 为下一轮循环更新之前的最小前缀和
        if prefixSum < minPreviousPrefixSum {
            minPreviousPrefixSum = prefixSum
        }
    }

    return maxSum
}