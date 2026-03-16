package cdyleetcode

// *************************************************
// 提前算出前缀乘积数组和后缀乘积数组。
// *************************************************
func productExceptSelf(nums []int) []int {
    length := len(nums)
    result := make([]int, length)
    prefix := make([]int, length)
    suffix := make([]int, length)
    // 计算所有前缀乘积
    prefixValue := 1
    for i:=0; i<length; i++ {
        prefix[i] = prefixValue * nums[i]
        prefixValue = prefix[i]
    }
    // 计算所有后缀乘积
    suffixValue := 1
    for i:=length-1; i>=0; i-- {
        suffix[i] = suffixValue * nums[i]
        suffixValue = suffix[i]
    }

    // 计算结果
    result[0] = suffix[1]
    result[length-1] = prefix[length-2]
    for i:=1; i<length-1; i++ {
        result[i] = prefix[i-1] * suffix[i+1]
    }

    return result
}



// 时间复杂度是O(n)
// 那就提前算出正数i个数的乘积
// 以及倒数i个数的乘积

// *************************************************
//
// *************************************************

// *************************************************
//
// *************************************************

// *************************************************
//
// *************************************************
