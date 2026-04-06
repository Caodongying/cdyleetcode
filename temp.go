package main

import "fmt"

// import "fmt"
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
    // 如果无效，依然返回
    if len(path)!= 0 &&!isValid(path, targetSum/2) {
        return
    }

    if len(path) == targetSum {
        temp := string(path)
        fmt.Printf("temp是%v", temp)
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

func main() {
   generateParenthesis(3)
}
