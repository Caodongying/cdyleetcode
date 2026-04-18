package cdyleetcode

// *************************************************
// ❌❌ （解法正确，但是超时了）自己写的，居然过了。但是超时。我根本没用动态规划。
// *************************************************
func wordBreak(s string, wordDict []string) bool {
    result := false
    for _, word := range wordDict {
        temp := s
        if len(word) <= len(temp) && word == temp[:len(word)] {
            temp = temp[len(word):]
            if len(temp) == 0 {
                return true
            }
            result = wordBreak(temp, wordDict)
            if result == true {
                return true
            }
        }
    }
    return false
}

// 对于wordDict的每一个单词，都拿去与s比对，看是否可以匹配s的某一个前缀
// 如果可以，就继续用wordDict去匹配s的去掉上一个前缀后的内容

// *************************************************
// 动态规划法
// *************************************************
func wordBreak(s string, wordDict []string) bool {
    // 初始化字符串字典
    wordMap := make(map[string]bool, len(wordDict))
    for _, word := range wordDict {
        wordMap[word] = true
    }

    dp := make([]bool, len(s)+1)
    dp[0] = true

    for i:=1; i <= len(s); i++ {
        for j:=i-1; j>=0; j-- {
            if dp[j]==true && wordMap[s[j:i]] == true {
                dp[i] = true
            }
        }
    }

    return dp[len(s)]
}

// dp[i]表示s的前i个元素构成的子串，是否可以由wordDict中的字符串构成
// 在s[i]之前的某个位置j，如果dp[j]为true（也就是0到j-1的子串可以被构成），且s[j:i]在wordDict里，就说明dp[i]也是true
// 先用map转存wordDict，优化寻找

// *************************************************
// 找到时提前break
// *************************************************
func wordBreak(s string, wordDict []string) bool {
    // 初始化字符串字典
    wordMap := make(map[string]bool, len(wordDict))
    for _, word := range wordDict {
        wordMap[word] = true
    }

    dp := make([]bool, len(s)+1)
    dp[0] = true

    for i:=1; i <= len(s); i++ {
        for j:=i-1; j>=0; j-- {
            if dp[j]==true && wordMap[s[j:i]] == true {
                dp[i] = true
                break
            }
        }
    }

    return dp[len(s)]
}

// dp[i]表示s的前i个元素构成的子串，是否可以由wordDict中的字符串构成
// 在s[i]之前的某个位置j，如果dp[j]为true（也就是0到j-1的子串可以被构成），且s[j:i]在wordDict里，就说明dp[i]也是true
// 先用map转存wordDict，优化寻找