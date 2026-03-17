package cdyleetcode

// *************************************************
// 看了灵神的题解后写的，先转置后翻转
// *************************************************
func rotate(matrix [][]int)  {
    n := len(matrix)
    // 转置
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }

    // 翻转每一行
    for i:=0; i < n; i++ {
        for j:=0; j<n/2; j++ {
            matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
        }
    }
}

// 顺时针旋转90°，等价于——
// 1. 先沿着主对角线转置 matrix(i,j)替换为matrix(j,i)
// 2. 再交换两端的列，而这也就是翻转每一行，可以用左右指针实现