package cdyleetcode

// *************************************************
// 建立矩阵和数组之间的映射关系，进行一次二分法	
// *************************************************
func searchMatrix(matrix [][]int, target int) bool {
    columns := len(matrix[0])
    left := 0
    right := len(matrix)*columns-1
    for left <= right {
        mid := left + (right-left)/2
        midValue := matrix[mid/columns][mid%columns]
        if target == midValue {
            return true
        } else if target > midValue {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return false
}

// 第一反应：两次二分法，先找行再确定列
// 看了灵神解法：一次二分法，把矩阵看成是一个递增的数组
// 所以问题变成怎么确立从数组元素a[i]到矩阵元素matrix[x][y]的映射
// x = i/len(列), y = i mod len(列)

