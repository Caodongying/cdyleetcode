package cdyleetcode

// *************************************************
// 看了题解后自己又对着敲了一遍，不太理解
// *************************************************
var result [][]int // 最终的返回值
var currentPath []int // 当前正在构建的路径
var currentUsed []bool // 当前路径中已经使用的值

func permute(nums []int) [][]int {
    n := len(nums)
    // 对三个全局变量初始化
    result = make([][]int, 0)
    currentPath = make([]int, n)
    currentUsed = make([]bool, n)

    // 开始回溯（递归)
    var dfs func(idx int) // 表示“确认path的第idx位”

    dfs = func(idx int) {
        if idx == n {
            // 当前path已经填满了
            // 因为后续currentPath还会被改变，需要copy一份
            temp := make([]int, len(currentPath))
            copy(temp, currentPath)
            result = append(result, temp)
            return
        }


        for i:=0; i<n; i++ {
            if currentUsed[i] == false {
                // 没使用过
                currentPath[idx] = nums[i]
                currentUsed[i] = true
                dfs(idx+1)
                // dfs(idx++)完成时，要继续寻找
                currentUsed[i] = false
            }
        }
    }

    dfs(0)

    return result

}



// *************************************************
// 用append添加path，并且进行元素退队，更好理解
// *************************************************
var result [][]int // 最终的返回值
var currentPath []int // 当前正在构建的路径
var currentUsed []bool // 当前路径中已经使用的值

func permute(nums []int) [][]int {
    n := len(nums)
    // 对三个全局变量初始化
    result = make([][]int, 0)
    currentPath = make([]int, 0, n)
    currentUsed = make([]bool, n)

    // 开始回溯（递归)
    var dfs func(idx int) // 表示“确认path的第idx位”

    dfs = func(idx int) {
        if idx == n {
            // 当前path已经填满了
            // 因为后续currentPath还会被改变，需要copy一份
            temp := make([]int, len(currentPath))
            copy(temp, currentPath)
            result = append(result, temp)
            return
        }


        for i:=0; i<n; i++ {
            if currentUsed[i] == false {
                // 没使用过
                currentPath = append(currentPath, nums[i])
                currentUsed[i] = true
                dfs(idx+1)
                // dfs(idx++)完成时，要继续寻找
                currentUsed[i] = false
                currentPath = currentPath[:len(currentPath)-1]
            }
        }
    }

    dfs(0)

    return result

}


