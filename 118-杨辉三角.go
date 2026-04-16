package cdyleetcode

// *************************************************
// ❌❌ 不是叫dfs，按顺序一层一层生成，不用两个参数。
// *************************************************
func generate(numRows int) [][]int {
    result := make([][]int, numRows)
    for i:=0; i<numRows; i++ {
        result[i] = make([]int, i+1)
    }

    var dfs func(level int, index int)
    dfs = func(level int, index int){
        if level == 0 || index == 0 || index == level {
            result[level] = append(result[level], 1)
            return
        }

        current := dfs(level-1, index-1) + dfs(level-1, index)
        result[level] = append(result[level], current)
    }

    dfs(numRows-1, numRows-1)

    return result
}

// dfs(level, index) = dfs(level-1, index-1) + dfs(level-1, index)

// *************************************************
// 改名为generateRow。数组初始化要小心，初始长度为0。level是从0开始，numRows从1开始，循环条件要小心
// *************************************************
func generate(numRows int) [][]int {
    result := make([][]int, numRows)
    for i:=0; i<numRows; i++ {
        // result[i] = make([]int, i+1)
        result[i] = make([]int, 0, i+1)
    }

    var generateRow func(level int) // 生成第i层, level从0开始
    generateRow = func(level int){
        if level == 0 {
            result[level] = append(result[level], 1)
            return
        }

        for i:=0; i<=level; i++ {
            current := 1
            if i != 0 && i != level {
                current = result[level-1][i-1] + result[level-1][i]
            }
            result[level] = append(result[level], current)
        }
    }

    for i:=0; i<numRows; i++ {
        generateRow(i)
    }

    return result
}

// *************************************************
// 不需要额外的函数，直接循环。而且，每一行直接分配数组大小。
// *************************************************
func generate(numRows int) [][]int {
    result := make([][]int, numRows)

    for i:=0; i<numRows; i++ {
        result[i] = make([]int, i+1)

        result[i][0], result[i][i] = 1, 1

        for j:=1; j<i; j++ {
            result[i][j] = result[i-1][j-1] + result[i-1][j]
        }
    }

    return result
}

// 关键是把每一排左对齐