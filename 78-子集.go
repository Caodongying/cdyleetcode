package cdyleetcode

// *************************************************
// ❌❌ 我的做法是全排列，也就是次序不一样也是一个结果
// *************************************************
func subsets(nums []int) [][]int {
    n := len(nums)
    result := make([][]int, 0)
    path := make([]int, 0)
    used := make([]bool, n)

    // 空子集
    result = append(result, []int{}) // 空数组不是[]或者[]int

    var dfs func(length int, idx int) // 找到总长度为length的数组的第idx位
    dfs = func(length int, idx int){
        if len(path) == length {
            // 当前路径满足要求
            temp := make([]int, length)
            copy(temp, path)
            result = append(result, temp)
            // 还原path
            path = []int{}
            return
        }

        for i:=0; i<n; i++ {
            if !used[i] { // nums[i]还没有使用过
                path = append(path, nums[i])
                used[i] = true
                dfs(length, idx+1)
                used[i] = false
            }
        }

    }

    // 依次罗列长度为i的子集
    for i := 1; i <= n; i++ {
        dfs(i, 0)
    }

    return result
}

// *************************************************
// deepseek改的。不需要used数组；dfs递归调用是i+1，不是idx+1；补上了节点弹出
// *************************************************
func subsets(nums []int) [][]int {
    n := len(nums)
    result := make([][]int, 0)
    path := make([]int, 0)

    // 空子集
    result = append(result, []int{}) // 空数组不是[]或者[]int

    var dfs func(length int, idx int) // 找到总长度为length的数组的第idx位
    dfs = func(length int, idx int){
        if len(path) == length {
            // 当前路径满足要求
            temp := make([]int, length)
            copy(temp, path)
            result = append(result, temp)
            return
        }

        for i:=idx; i<n; i++ {
            path = append(path, nums[i])
            dfs(length, i+1)
            path = path[:len(path)-1]
        }

    }

    // 依次罗列长度为i的子集
    for i := 1; i <= n; i++ {
        dfs(i, 0)
    }

    return result
}

// *************************************************
// deepseek优化，不需要数组length。不按照“子集的大小”分组，
// 而是按照“从哪个下标出发”。因为子集无序，取过的元素不会重复取，
// 写回溯算法的时候，for就要从startIndex开始，而不是从0开始
// *************************************************
func subsets(nums []int) [][]int {
    n := len(nums)
    result := make([][]int, 0)
    path := make([]int, 0)

    // 空子集
    result = append(result, []int{}) // 空数组不是[]或者[]int

    var dfs func(start int) // 以当前 path 为基础，尝试在 nums[start:] 范围内添加元素，生成所有可能的子集。
    dfs = func(start int){ // start：表示接下来可以选择的元素在 nums 中的起始索引
        // 不需要等长度固定，而是每层都记录当前路径
        for i:=start; i<n; i++ {
            path = append(path, nums[i])

            temp := make([]int, len(path))
            copy(temp, path)
            result = append(result, temp)

            dfs(i+1)

            path = path[:len(path)-1]
        }
    }

    dfs(0)

    return result
}

// *************************************************
// 优化，不单独处理空集合
// *************************************************
func subsets(nums []int) [][]int {
    n := len(nums)
    result := make([][]int, 0)
    path := make([]int, 0)

    var dfs func(start int) // 以当前 path 为基础，尝试在 nums[start:] 范围内添加元素，生成所有可能的子集。
    dfs = func(start int){ // start：表示接下来可以选择的元素在 nums 中的起始索引
        temp := make([]int, len(path))
        copy(temp, path)
        result = append(result, temp)
        // 不需要等长度固定，而是每层都记录当前路径
        for i:=start; i<n; i++ {
            path = append(path, nums[i])

            dfs(i+1)

            path = path[:len(path)-1]
        }
    }

    dfs(0)

    return result
}