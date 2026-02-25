package cdyleetcode

// *************************************************
// 动态规划法
// *************************************************

// ①确认dp数组维数和含义dp[i][j]: 从下标i到j的子串是否为回文子串
// ②递推公式
// 如果s[i] == s[j]，那么需要分长度。如果子串长度小于等于3，就是true；否则看里面的i+1到j-1子串，也就是dp[i+1][j-1]是否为true
// 如果s[i] != s[j], dp[i][j]就肯定是false
// ③初始化。因为在看外子串的回文属性时，依赖的是内子串，所以按照长度进行初始化（死记硬背）。长度1的子串，都是yes，长度2或3的子串，看端点是否相等
// ④长度大于3，就用递推公式了


func longestPalindrome(s string) string {
    length := len(s)

    if length <= 1 {
        return s
    }

    // 初始化dp[][]并分配
    dp := make([][]bool, length)
    for i := range dp {
        dp[i] = make([]bool, length)
    }

    start := 0
    maxLen := 1

    // 先初始化长度为1的子串的回文属性
    for i:=0; i<length; i++ {
        dp[i][i] = true
    }

    // 初始化之后的长度
    for subLen := 2; subLen <= length; subLen++{
        for i:=0; i+subLen-1 < length; i++ {
            j := i+subLen-1
            if s[i]==s[j] {
                if subLen <= 3 {
                    dp[i][j] = true
                } else {
                    dp[i][j] = dp[i+1][j-1]
                }

                if dp[i][j] && subLen > maxLen {
                    start = i
                    maxLen = subLen
                }
            }
        }
    }

    return s[start : start+maxLen]
}
