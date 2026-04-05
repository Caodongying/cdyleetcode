package cdyleetcode

// *************************************************
// 自己写的，每一层都是数组中的一个数出现几次，依次递增。错误原因：
// 1. 提前剪枝条件错误：candidates[idx] > remain 这个判断只在 idx 不变时有效，但递归中 remain 会变化，且 candidates 未排序时这个判断不准确。
// 2. 切片引用问题：result = append(result, combination) 添加的是同一个切片引用，后续修改会影响之前的结果。应该用temp!!
// 3. 循环逻辑复杂：三层循环不必要，标准 DFS 回溯写法更清晰。
// *************************************************
var result [][]int
var combination []int

func combinationSum(candidates []int, target int) [][]int {
    // 初始化
    result = make([][]int, 0)
    combination = make([]int, 0)

    dfs(candidates, 0, target)

    return result
}

func dfs(candidates []int, idx int, remain int) {
    if idx >= len(candidates) || candidates[idx] > remain {
        return
    }

    if remain == 0 {
        result = append(result, combination)
        return
    }

    for i := idx; i < len(candidates); i++ {
        // j的最大值表示本轮最多可以append的candidates[i]个数
        // [2] 8
        for j := 0; j <= remain/candidates[i]; j++ {  // 是i不是idx
            for k:=0; k<j; k++ {
                combination = append(combination, candidates[i])
            }
            dfs(candidates, i+1, remain-candidates[i]*j)
            combination = combination[:len(combination)-j]
        }
    }
}

// 树形结构构想：
// 第0层： 0到target/candidate[0]个candidate[0]
// 第1层： 0到(target-第0层的结果和)/candidate[1]个candidate[1]

// *************************************************
// 看deepseek改的，不用多重循环，而是dfs的时候用i而不是i+1
// *************************************************
var result [][]int
var combination []int

func combinationSum(candidates []int, target int) [][]int {
    // 初始化
    result = make([][]int, 0)
    combination = make([]int, 0)

    dfs(candidates, 0, target)

    return result
}

func dfs(candidates []int, idx int, remain int) {
    if remain == 0 {
        temp := make([]int, len(combination))
        copy(temp, combination)
        result = append(result, temp)
        return
    }

    for i := idx; i < len(candidates); i++ {
        if candidates[i] > remain {
            continue
        }

        combination = append(combination, candidates[i])
        dfs(candidates, i, remain-candidates[i]) // 这里i不变，因为i可以重复使用
        combination = combination[:len(combination)-1]
    }
}

// *************************************************
// 优化：先排序。在求和问题中，排序之后加剪枝是常见的套路！
// *************************************************
var result [][]int
var combination []int

func combinationSum(candidates []int, target int) [][]int {
    // 初始化
    result = make([][]int, 0)
    combination = make([]int, 0)

    sort.Ints(candidates)
    dfs(candidates, 0, target)

    return result
}

func dfs(candidates []int, idx int, remain int) {
    if remain == 0 {
        temp := make([]int, len(combination))
        copy(temp, combination)
        result = append(result, temp)
        return
    }

    for i := idx; i < len(candidates); i++ {
        if candidates[i] > remain {
            break
        }

        combination = append(combination, candidates[i])
        dfs(candidates, i, remain-candidates[i]) // 这里i不变，因为i可以重复使用
        combination = combination[:len(combination)-1]
    }
}
