package cdyleetcode

// *************************************************
// ❌❌ 没注意返回类型是[]string，而不是[][]string
// *************************************************
var result []string
var path []string

func generateParenthesis(n int) []string {
    result = make([][]string, 0)
    path = make([]string, 0, 2*n)

    brackets := []string{"(", "(", "(", ")", ")", ")"}

    dfs(brackets, 0, 2*n)

    return result
}

func dfs(brackets []string, start int, targetSum int) {
    if len(path) == targetSum {
        temp := make([]string, targetSum)
        copy(temp, path)
        result = append(result, temp)
        return
    }

    // 如果无效，依然返回
    if !isValid(path, targetSum/2) {
        return
    }

    for i:=start; i<2*n; i++ {
        path = append(path, brackets[i])
        dfs(brackets, start+1, targetSum)
        path = path[:len(path)-1]
    }
}

func isValid(path []string, pair int) bool {
    // 遇到左括号就压栈
    // 遇到右括号就弹栈。如果栈空就返回bool
    stack := make([]string, 0, pair)

    for _, v := range path {
        if v == "(" {
            stack = append(stack, v)
        } else if v == ")" {
            if len(stack) == 0 {
                return false
            }
        }
    }

    return true

}

// 遍历所有的组合
// 如果无效就回溯

// *************************************************
// ❌❌ 错误地当成无序的组合做了。其实应该是有序的
// *************************************************
var result []string
var path []rune

func generateParenthesis(n int) []string {
    result = make([]string, 0)
    path = make([]rune, 0, 2*n)

    brackets := make([]rune, 0, n*2)
    for i:=0; i<n; i++ {
        brackets = append(brackets, '(')
        brackets = append(brackets, ')')
    }

    dfs(brackets, 0, 2*n)

    return result
}

func dfs(brackets []rune, start int, targetSum int) {
    if len(path) == targetSum {
        temp := string(path)
        result = append(result, temp)
        return
    }

    // 如果无效，依然返回
    if !isValid(path, targetSum/2) {
        return
    }

    for i:=start; i<targetSum; i++ {
        path = append(path, brackets[i])
        dfs(brackets, start+1, targetSum)
        path = path[:len(path)-1]
    }
}

func isValid(path []rune, pair int) bool {
    // 遇到左括号就压栈
    // 遇到右括号就弹栈。如果栈空就返回bool
    stack := make([]rune, 0, pair)

    for _, v := range path {
        if v == '('{
            stack = append(stack, v)
        } else if v == ')' {
            if len(stack) == 0 {
                return false
            }
            // 别忘了弹栈
            stack = stack[:len(stack)-1]
        }
    }

    if len(stack) != 0 { // 千万不能漏了
        return false
    }

    return true

}

// 遍历所有的组合
// 如果无效就回溯

// *************************************************
// ❌❌ 写的排列不对。
// *************************************************
var result []string
var path []rune
var visited []bool

func generateParenthesis(n int) []string {
    result = make([]string, 0)
    path = make([]rune, 0, 2*n)
    visited = make([]bool, 2*n)

    brackets := make([]rune, 0, n*2)
    for i:=0; i<n; i++ {
        brackets = append(brackets, '(')
        brackets = append(brackets, ')')
    }

    dfs(brackets, 2*n)

    return result
}

func dfs(brackets []rune, targetSum int) {
    // 如果无效，返回
    if len(path)!=0 && !isValid(path, targetSum/2) {
        return
    }

    if len(path) == targetSum {
        temp := string(path)
        result = append(result, temp)
        return
    }


    for i:=0; i<targetSum; i++ {
        if !visited[i] {
            path = append(path, brackets[i])
            visited[i] = true
        }
        dfs(brackets, targetSum)
        path = path[:len(path)-1]
        visited[i] = false
    }
}

func isValid(path []rune, pair int) bool {
    // 遇到左括号就压栈
    // 遇到右括号就弹栈。如果栈空就返回bool
    stack := make([]rune, 0, pair)

    for _, v := range path {
        if v == '('{
            stack = append(stack, v)
        } else if v == ')' {
            if len(stack) == 0 {
                return false
            }
            // 别忘了弹栈
            stack = stack[:len(stack)-1]
        }
    }

    if len(stack) != 0 { // 千万不能漏了
        return false
    }

    return true

}

// 遍历所有的排列
// 如果无效就回溯
// 但是如何得到所有排列？

// *************************************************
// ❌❌ 没有进行去重！
// *************************************************
var result []string
var path []rune
var visited []bool

func generateParenthesis(n int) []string {
    result = make([]string, 0)
    path = make([]rune, 0, 2*n)
    visited = make([]bool, 2*n)

    brackets := make([]rune, 0, n*2)
    for i:=0; i<n; i++ {
        brackets = append(brackets, '(')
        brackets = append(brackets, ')')
    }

    dfs(brackets, 2*n)

    return result
}

func dfs(brackets []rune, targetSum int) {
    // 如果无效，返回
    if len(path)!=0 && !isValid(path, targetSum/2) {
        return
    }

    if len(path) == targetSum {
        temp := string(path)
        result = append(result, temp)
        return
    }


    for i:=0; i<targetSum; i++ {
        if !visited[i] {
            path = append(path, brackets[i])
            visited[i] = true
            dfs(brackets, targetSum)
            path = path[:len(path)-1]
            visited[i] = false
        }
    }
}

func isValid(path []rune, pair int) bool {
    // 遇到左括号就压栈
    // 遇到右括号就弹栈。如果栈空就返回bool
    stack := make([]rune, 0, pair)

    for _, v := range path {
        if v == '('{
            stack = append(stack, v)
        } else if v == ')' {
            if len(stack) == 0 {
                return false
            }
            // 别忘了弹栈
            stack = stack[:len(stack)-1]
        }
    }

    // if len(stack) != 0 { // 千万不能加，因为这是中间状态
    //     return false
    // }

    return true

}

// 遍历所有的排列
// 如果无效就回溯
// 但是如何得到所有排列？

// *************************************************
// ❌❌ 看了Gemini，加了去重，但是去重的前提是数组有序
// *************************************************
var result []string
var path []rune
var visited []bool

func generateParenthesis(n int) []string {
    result = make([]string, 0)
    path = make([]rune, 0, 2*n)
    visited = make([]bool, 2*n)

    brackets := make([]rune, 0, n*2)
    for i:=0; i<n; i++ {
        brackets = append(brackets, '(')
        brackets = append(brackets, ')')
    }

    dfs(brackets, 2*n)

    return result
}

func dfs(brackets []rune, targetSum int) {
    // 如果无效，返回
    if len(path)!=0 && !isValid(path, targetSum/2) {
        return
    }

    if len(path) == targetSum {
        temp := string(path)
        result = append(result, temp)
        return
    }


    for i:=0; i<targetSum; i++ {
        if i>0 && brackets[i] == brackets[i-1] && !visited[i-1] { // 最为关键的去重
            continue
        }
        if !visited[i] {
            path = append(path, brackets[i])
            visited[i] = true
            dfs(brackets, targetSum)
            path = path[:len(path)-1]
            visited[i] = false
        }
    }
}

func isValid(path []rune, pair int) bool {
    // 遇到左括号就压栈
    // 遇到右括号就弹栈。如果栈空就返回bool
    stack := make([]rune, 0, pair)

    for _, v := range path {
        if v == '('{
            stack = append(stack, v)
        } else if v == ')' {
            if len(stack) == 0 {
                return false
            }
            // 别忘了弹栈
            stack = stack[:len(stack)-1]
        }
    }

    // if len(stack) != 0 { // 千万不能加，因为这是中间状态
    //     return false
    // }

    return true

}

// 遍历所有的排列
// 如果无效就回溯
// 但是如何得到所有排列？

// *************************************************
// 重新构造brackets数组
// *************************************************
var result []string
var path []rune
var visited []bool

func generateParenthesis(n int) []string {
    result = make([]string, 0)
    path = make([]rune, 0, 2*n)
    visited = make([]bool, 2*n)

    brackets := make([]rune, 0, n*2)
    for i := 0; i < n; i++ {
        brackets = append(brackets, '(')
    }
    for i := 0; i < n; i++ {
        brackets = append(brackets, ')')
    }

    dfs(brackets, 2*n)

    return result
}

func dfs(brackets []rune, targetSum int) {
    // 如果无效，返回
    if len(path)!=0 && !isValid(path, targetSum/2) {
        return
    }

    if len(path) == targetSum {
        temp := string(path)
        result = append(result, temp)
        return
    }


    for i:=0; i<targetSum; i++ {
        if i>0 && brackets[i] == brackets[i-1] && !visited[i-1] { // 最为关键的去重
            continue
        }
        if !visited[i] {
            path = append(path, brackets[i])
            visited[i] = true
            dfs(brackets, targetSum)
            path = path[:len(path)-1]
            visited[i] = false
        }
    }
}

func isValid(path []rune, pair int) bool {
    // 遇到左括号就压栈
    // 遇到右括号就弹栈。如果栈空就返回bool
    stack := make([]rune, 0, pair)

    for _, v := range path {
        if v == '('{
            stack = append(stack, v)
        } else if v == ')' {
            if len(stack) == 0 {
                return false
            }
            // 别忘了弹栈
            stack = stack[:len(stack)-1]
        }
    }

    // if len(stack) != 0 { // 千万不能加，因为这是中间状态
    //     return false
    // }

    return true

}

// 遍历所有的排列
// 如果无效就回溯
// 但是如何得到所有排列？

// *************************************************
// 不用栈，直接用left个数判断括号序列是否有效。注意，只要保证left大于等于right括号。如果left偏多是可以的，因为括号序列还没有构建完毕。
// *************************************************
var result []string
var path []rune
var visited []bool

func generateParenthesis(n int) []string {
    result = make([]string, 0)
    path = make([]rune, 0, 2*n)
    visited = make([]bool, 2*n)

    brackets := make([]rune, 0, n*2)
    for i := 0; i < n; i++ {
        brackets = append(brackets, '(')
    }
    for i := 0; i < n; i++ {
        brackets = append(brackets, ')')
    }

    dfs(brackets, 2*n)

    return result
}

func dfs(brackets []rune, targetSum int) {
    // 如果无效，返回
    if len(path)!=0 && !isValid(path, targetSum/2) {
        return
    }

    if len(path) == targetSum {
        temp := string(path)
        result = append(result, temp)
        return
    }


    for i:=0; i<targetSum; i++ {
        if i>0 && brackets[i] == brackets[i-1] && !visited[i-1] { // 最为关键的去重
            continue
        }
        if !visited[i] {
            path = append(path, brackets[i])
            visited[i] = true
            dfs(brackets, targetSum)
            path = path[:len(path)-1]
            visited[i] = false
        }
    }
}

func isValid(path []rune, pair int) bool {
    left := 0

    for _, v := range path {
        if v == '('{
            left++
        } else if v == ')' {
            left--
            if left < 0 {
                return false
            }
        }
    }

    return true

}

// 遍历所有的排列
// 如果无效就回溯
// 但是如何得到所有排列？

// *************************************************
// ❌❌ dfs调用大错误
// *************************************************
var result []string

func generateParenthesis(n int) []string {
    result = make([]string, 0)

    dfs(n, n, "")

    return result
}

func dfs(left int, right int, path string) {
    if left == 0 && right == 0 {
        result = append(result, path)
        return
    }

    if left > 0 {
        path = path + "("
        left--
        dfs(left, right, path)
    }

    if right > left {
        path = path + ")"
        right--
        dfs(left, right, path)
    }

}

// 对于一个有效的括号序列的任意前缀，都满足左括号个数大于等于右括号个数
// 所以这个回溯法比较特别
// left表示还剩几个左括号，right表示还剩几个右括号
// 如果left大于0，加入(；如果right比left多，说明右括号剩下的多，而当前序列里(是多的，需要right

// *************************************************
// dfs传值更新.不要在当前层修改变量的值，而是直接把运算结果传递给下一层
// *************************************************
var result []string

func generateParenthesis(n int) []string {
    result = make([]string, 0)

    dfs(n, n, "")

    return result
}

func dfs(left int, right int, path string) {
    if left == 0 && right == 0 {
        result = append(result, path)
        return
    }

    if left > 0 {
        dfs(left-1, right, path + "(")
    }

    if right > left {
        dfs(left, right-1, path + ")")
    }

}

// 对于一个有效的括号序列的任意前缀，都满足左括号个数大于等于右括号个数
// 所以这个回溯法比较特别
// left表示还剩几个左括号，right表示还剩几个右括号
// 如果left大于0，加入(；如果right比left多，说明右括号剩下的多，而当前序列里(是多的，需要right