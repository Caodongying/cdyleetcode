package cdyleetcode

// *************************************************
// 灵神解法，直接看右上角，进行排除法
// *************************************************
func searchMatrix(matrix [][]int, target int) bool {
    top := 0
    right := len(matrix[0])-1

    for top < len(matrix) && right >= 0 {
        topRightValue := matrix[top][right]

        if target == topRightValue{
            return true
        } else if target > topRightValue {
            top++
        } else {
            right--
        }
    }

    return false
}

// 直接看了灵神的解法。灵神是只看右上角的值，然后进行排除
// 我想了想，用左下角的值也行。但是不能用左上角或者右下角
// 因为左上角和右下角，是（最小，最小）或者（最大，最大）的组合。
// 而排序法药借助（最小，最大）或者（最大，最小）的组合。

// target如果大于右上角值，说明必定不在top行里，排除top行，top++
// target如果小于右上角值，说明必定不在right列里，排除right列，right--

// *************************************************
// 一点点优化
// *************************************************
func searchMatrix(matrix [][]int, target int) bool {
    n := len(matrix)
    top := 0
    right := len(matrix[0])-1

    for top < n && right >= 0 {
        topRightValue := matrix[top][right]

        if target == topRightValue{
            return true
        }
        if target > topRightValue {
            top++
        } else {
            right--
        }
    }

    return false
}
