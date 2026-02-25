package cdyleetcode

func longestCommonSubsequence(text1 string, text2 string) int {
    // var dp [[]int]int := make([[len(text1)]int]int, len(text2))
    // dp := make([][]int, len(text1)+1, len(text2)+1) 不是这么初始化的
    dp := make([][]int, len(text1)+1)
    for i := range dp {
        dp[i] = make([]int, len(text2)+1)
    }

    // 初始化dp的第一行和第一列为0
    for i:=0; i<len(dp); i++ {
        dp[i][0] = 0
    }
    for j:=0; j<len(dp[0]); j++ {
        dp[0][j] = 0
    }

    for i:=1; i<len(dp); i++ {
        for j:=1; j<len(dp[0]); j++{
            if text1[i-1] == text2[j-1] {
                dp[i][j] = 1 + dp[i-1][j-1]
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }

    return dp[len(text1)][len(text2)]


}

// 直接看的b站波波微课的讲解
// 二维dp数组，因为有俩字符串
// dp[i][j]表示text1到第i个字符为止的子串和和text2到第j个字符为止的子串的最长公共子序列
// dp的第一行和第一列初始化为0