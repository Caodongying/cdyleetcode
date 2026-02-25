package cdyleetcode

// *************************************************
// ❌❌ 错误解法，一边扫描一边置零，无法区分哪些是原本数组里的0，哪些是被置零的
// *************************************************
func setZeroes(matrix [][]int)  {
    // 记录已经被置零的行和列
    alreadySetRow := make(map[int]int)
    alreadySetColumn := make(map[int]int)

    // 遍历矩阵
    for row := 0; row < len(matrix); row++ {
        for column := 0; column < len(matrix[0]); column++ {
            if matrix[row][column] == 0 {
                if _, exists := alreadySetRow[row]; !exists {
                    // setRowZero(matrix, row) // 如何操作matrix本身而不是副本
                    // 将本行全设置为0
                    for i:=0; i<len(matrix[0]); i++ {
                        matrix[row][i] = 0
                    }

                    // 更新map
                    alreadySetRow[row] = 1
                }
                if _, exists := alreadySetColumn[column]; !exists {
                    // setColumnZero(matrix, column)
                    // 将本列全设置为0
                    for j:=0; j<len(matrix); j++ {
                        matrix[j][column] = 0
                    }
                    // 更新map
                    alreadySetColumn[column] = 1
                }
            }
        }
    }
}


// *************************************************
// 先统计需要置零的位，再统一置零。
// *************************************************

type position struct {
    row int
    column int
}

func setZeroes(matrix [][]int)  {
    // 记录0位
    rowCount := len(matrix)
    columnCount := len(matrix[0])
    zeroPosition := make([]position, 0, rowCount*columnCount)

    alreadySetRow := make(map[int]int)
    alreadySetColumn := make(map[int]int)

    // 遍历矩阵
    for row := 0; row < rowCount; row++ {
        for column := 0; column < columnCount; column++ {
            if matrix[row][column] == 0 {
                zeroPosition = append(zeroPosition, position{row, column})
            }
        }
    }

    // 置零
    for _, p := range(zeroPosition) {
        row, column := p.row, p.column
        if _, exists := alreadySetRow[row]; !exists {
            // 将本行全设置为0
            for i:=0; i<len(matrix[0]); i++ {
                matrix[row][i] = 0
            }
            // 更新map
            alreadySetRow[row] = 1
        }
        if _, exists := alreadySetColumn[column]; !exists {
            // 将本列全设置为0
            for j:=0; j<len(matrix); j++ {
                matrix[j][column] = 0
            }
            // 更新map
            alreadySetColumn[column] = 1
        }
    }
}

// *************************************************
// 用两个map来分别保存哪些行和列需要置零。
// *************************************************
type position struct {
    row int
    column int
}

func setZeroes(matrix [][]int)  {
    // 记录0位
    rowCount := len(matrix)
    columnCount := len(matrix[0])

    toSetRow := make(map[int]int)
    toSetColumn := make(map[int]int)

    // 遍历矩阵
    for row := 0; row < rowCount; row++ {
        for column := 0; column < columnCount; column++ {
            if matrix[row][column] == 0 {
                toSetRow[row] = 1
                toSetColumn[column] = 1
            }
        }
    }

    // 置零
    for row, _ := range(toSetRow) {
        // 将本行全设置为0
        for i:=0; i<len(matrix[0]); i++ {
            matrix[row][i] = 0
        }
    }
    for column, _ := range(toSetColumn) {
        // 将本列全设置为0
        for j:=0; j<len(matrix); j++ {
            matrix[j][column] = 0
        }
    }

}

// *************************************************
// 题解里的方法二
// *************************************************
func setZeroes(matrix [][]int) {
    n, m := len(matrix), len(matrix[0])
    row0, col0 := false, false
    for _, v := range matrix[0] {
        if v == 0 {
            row0 = true
            break
        }
    }
    for _, r := range matrix {
        if r[0] == 0 {
            col0 = true
            break
        }
    }
    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }
    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }
    if row0 {
        for j := 0; j < m; j++ {
            matrix[0][j] = 0
        }
    }
    if col0 {
        for _, r := range matrix {
            r[0] = 0
        }
    }
}