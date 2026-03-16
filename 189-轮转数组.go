package cdyleetcode

// *************************************************
// ❌❌题目隐含要求是改变参数切片本身，用切片复制没用
// *************************************************
func rotate(nums []int, k int)  {
    k = k % len(nums)

    left := nums[len(nums)-k:len(nums)]

    right := nums[:len(nums)-k]

    nums = append(left, right...)
}

// *************************************************
// 使用copy可以直接改变原切片
// *************************************************
func rotate(nums []int, k int) {
    k = k % len(nums)

    left := nums[len(nums)-k:len(nums)]

    right := nums[:len(nums)-k]

    // nums = append(left, right...)
    result := append(left, right...)
    copy(nums, result)
}

// *************************************************
// 原地翻转
// *************************************************
func rotate(nums []int, k int) {
    k = k % len(nums)

    reverse := func(nums []int) {
        // for i, j := 0, len(nums)-1; i != j; i++, j-- {
        for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
            nums[i], nums[j] = nums[j], nums[i]
        }
    }

    // 翻转整个数组
    // 翻转前k个
    // 翻转后面的

    reverse(nums)
    reverse(nums[:k])
    reverse(nums[k:])
}


// *************************************************
// 灵神用了slices包
// *************************************************
func rotate(nums []int, k int) {
    k = k % len(nums)

    slices.Reverse(nums)
    slices.Reverse(nums[:k])
    slices.Reverse(nums[k:])

}
