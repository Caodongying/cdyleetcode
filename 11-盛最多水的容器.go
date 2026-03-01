package cdyleetcode

// *************************************************
// 双指针法
// *************************************************

func maxArea(height []int) int {
    left := 0
    right := len(height) - 1

    max := 0

    for left != right {
        var current int
        if height[left] >= height[right] {
            // 右板短
            current = (right-left) * height[right] // ⚠️这里注意，不可以用current:=...代替签名var声明，因为current是在定义变量的区域之外使用的，会有declared but not used的报错！
            right--
        } else {
            // 左板短
            current = (right-left) * height[left]
            left++
        }
        if current > max {
            max = current
        }
    }

    return max
}