package cdyleetcode

// *************************************************
// 直接照抄代码随想录题解，自己其实还没有深刻理解
// *************************************************

var result [][]string
var path []string

func partition(s string) [][]string {
    result = make([][]string, 0)
    path = make([]string, 0)

    dfs(s, 0)

    return result
}

func dfs(s string, start int) { // 通过递归尝试在字符串的每个位置“切一刀”，如果切下来的左半部分是回文，就去剩下的右半部分里继续重复这个过程。
    if start == len(s) {
        temp := make([]string, len(path))
        copy(temp, path)
        result = append(result, temp)
        return
    }
    for i:=start; i<len(s); i++ {
        str := s[start:i+1]
        if isPalindrome(str){
            path = append(path, str)
            dfs(s, i+1)
            path = path[:len(path)-1]
        }
    }
}

func isPalindrome(s string) bool {
    for i, j := 0, len(s)-1; i<j; i, j = i+1, j-1 { // 不可以是i!=j
        if s[i] != s[j] {
            return false
        }
    }
    return true
}
// 一开始自己想的是，每一个单独的字符都是回文子串，所以就看怎么拼接。但这样行不通
