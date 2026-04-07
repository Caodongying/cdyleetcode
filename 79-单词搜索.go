package cdyleetcode

// *************************************************
// 忽略了一个条件：不可以重复使用同一个单元格
// *************************************************
func exist(board [][]byte, word string) bool {
    m := len(board) // 行数
    n := len(board[0]) // 列数

    var dfs func(row int, column int, cur int) bool
    dfs = func(row int, column int, cur int) bool {
        if cur == len(word) {
            return true
        }

        if row < 0 || row > m-1 || column < 0 || column > n-1 {
            return false
        }

        if word[cur] != board[row][column] {
            return false
        }



        canFindNext := dfs(row-1, column, cur+1) || dfs(row+1, column, cur+1) || dfs(row, column-1, cur+1) || dfs(row, column+1, cur+1)

        return canFindNext
    }

    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            if board[i][j] == word[0] {
                // 找到了第一个字符
                find := dfs(i, j, 0)
                if find == true {
                    return true
                }
            }
        }
    }

    // 没有找到第一个字符
    return false
}

// *************************************************
// visited数组使用错误
// *************************************************
func exist(board [][]byte, word string) bool {
    m := len(board) // 行数
    n := len(board[0]) // 列数
    visited := make([][]bool, m)
    for i:=0; i<m; i++ {
        visited[i] = make([]bool, n)
    }

    var dfs func(row int, column int, cur int) bool // 表示从[row][column]开始，能不能找到word中，从cur下标开始的之后的部分
    dfs = func(row int, column int, cur int) bool {
        if cur == len(word) {
            return true
        }

        if row < 0 || row > m-1 || column < 0 || column > n-1 {
            return false
        }

        if word[cur] != board[row][column] {
            return false
        }

        visited[row][column] = true

        canFindNext := dfs(row-1, column, cur+1) || dfs(row+1, column, cur+1) || dfs(row, column-1, cur+1) || dfs(row, column+1, cur+1)

        if !canFindNext {
            visited[row][column] = false
        }

        return canFindNext
    }

    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            if board[i][j] == word[0] {
                // 找到了第一个字符
                find := dfs(i, j, 0)
                if find == true {
                    return true
                }
            }
        }
    }

    // 没有找到第一个字符
    return false
}

// *************************************************
// deepseek给的修改，注意判断是否访问过，是就返回false
// *************************************************
func exist(board [][]byte, word string) bool {
    m := len(board) // 行数
    n := len(board[0]) // 列数
    visited := make([][]bool, m)
    for i:=0; i<m; i++ {
        visited[i] = make([]bool, n)
    }

    var dfs func(row int, column int, cur int) bool // 表示从[row][column]开始，能不能找到word中，从cur下标开始的之后的部分
    dfs = func(row int, column int, cur int) bool {
        if cur == len(word) {
            return true
        }

        if row < 0 || row > m-1 || column < 0 || column > n-1 {
            return false
        }

        if visited[row][column] || word[cur] != board[row][column] {
            return false
        }

        visited[row][column] = true

        canFindNext := dfs(row-1, column, cur+1) || dfs(row+1, column, cur+1) || dfs(row, column-1, cur+1) || dfs(row, column+1, cur+1)

        visited[row][column] = false

        return canFindNext
    }

    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            if board[i][j] == word[0] {
                // 找到了第一个字符
                find := dfs(i, j, 0)
                if find == true {
                    return true
                }
            }
        }
    }

    // 没有找到第一个字符
    return false
}
