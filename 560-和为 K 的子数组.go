package cdyleetcode

// *************************************************
// ❌❌ 错误，因为算上了之前的值，应该是看valueCountMap中，当前位之后有多少个数
// *************************************************
func subarraySum(nums []int, k int) int {
    // 计算所有前缀和
    prefixSum := make([]int, len(nums))
    prefixSum[0] = nums[0]
    for i:=1; i<len(nums); i++ {
        prefixSum[i] = prefixSum[i-1] + nums[i]
    }

   // 将数组构造为字典，key是数组值，value是出现次数
    valueCountMap := make(map[int]int)
    for _, v := range(prefixSum) {
        if _, exists := valueCountMap[v]; !exists {
            valueCountMap[v] = 1
        } else {
            valueCountMap[v] += 1
        }
    }

    // 计算子数组个数
    sum := 0
    var expectedPrefixSum int
    for j:=0; j<len(nums); j++ {
    if j==0 {
        expectedPrefixSum = k
    } else {
        expectedPrefixSum = k + prefixSum[j-1]
    }
    sum += valueCountMap[expectedPrefixSum]
    }

    return sum
}
// 感觉只有暴力解法

// *************************************************
// 每来到一个新的起始位，就将前一位对应的前缀和数量-1，表示“上一位无效”。
// 这个解法是先把前缀和全部算出来，保存在数组里，然后再用字典进行次数统计，并且带有一个“之前的前缀和失效”的考虑。不太好。
// *************************************************
func subarraySum(nums []int, k int) int {
    // 计算所有前缀和
    prefixSum := make([]int, len(nums))
    prefixSum[0] = nums[0]
    for i:=1; i<len(nums); i++ {
        prefixSum[i] = prefixSum[i-1] + nums[i]
    }

   // 将数组构造为字典，key是数组值，value是出现次数
    valueCountMap := make(map[int]int)
    for _, v := range(prefixSum) {
        if _, exists := valueCountMap[v]; !exists {
            valueCountMap[v] = 1
        } else {
            valueCountMap[v] += 1
        }
    }

    // 计算子数组个数
    sum := 0
    var expectedPrefixSum int
    for j:=0; j<len(nums); j++ {
        if j==0 {
            expectedPrefixSum = k
        } else {
            valueCountMap[prefixSum[j-1]]-- // 重要
            expectedPrefixSum = k + prefixSum[j-1]
        }
        sum += valueCountMap[expectedPrefixSum]
    }

    return sum
}

// *************************************************
// 动态构建map，而不是先计算所有前缀和再转化为字典
// *************************************************

func subarraySum(nums []int, k int) int {
    subArrayCount := 0
    // 将数组构造为字典，key是数组值，value是出现次数
    prefixCountMap := make(map[int]int)
    prefixSum := 0

    prefixCountMap[0] = 1 // 处理空数组

    for _, num := range(nums) {
        prefixSum += num
        if count, exists := prefixCountMap[prefixSum-k]; exists {
            subArrayCount += count
        }
        prefixCountMap[prefixSum]++
    }
    return subArrayCount
    }

// 子数组[i,j]（i到j），满足prefix[j]-prefix[i]=k，也就是要存在prefix[i]=prefix[j]-k

// 如果存在某个 prefixSum[i-1] 等于 prefixSum[j] - k，那么 nums[i..j] 就是一个符合条件的子数组。
