package cdyleetcode

// *************************************************
// 看了笨猪爆破组的题解后自己写的
// *************************************************
func spiralOrder(matrix [][]int) []int {
    top := 0
    bottom := len(matrix)-1
    left := 0
    right := len(matrix[0])-1

    result := make([]int, 0, len(matrix)*len(matrix[0]))

    for left < right && top < bottom {
        // →
        for i:= left; i<=right; i++ {
            result = append(result, matrix[top][i])
        }
        top++
        // ↓
        for i:= top; i<=bottom; i++ {
            result = append(result, matrix[i][right])
        }
        right--
        // ←
        for i:= right; i>= left; i-- {
            result = append(result, matrix[bottom][i])
        }
        bottom--
        // ↑
        for i := bottom; i>=top; i-- {
            result = append(result, matrix[i][left])
        }
        left++
    }

    // 循环退出时，可能的两情况
    // 1. left == right 只剩一列
    if left == right {
        for i:=top; i<= bottom; i++ {
            result = append(result, matrix[i][left])
        }
    } else if top == bottom {
    // 2. top == bottom 只剩一行
        for i:= left; i<= right; i++ {
            result = append(result, matrix[top][i])
        }
    }

    return result
}

// 设置上下左右四个边界
// 按照四个方向进行遍历，更新边界
// 最后处理剩余行/列
