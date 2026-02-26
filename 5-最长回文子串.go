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


// *************************************************
// 中心扩散法
// *************************************************

// 中心扩散法
// 一个循环，针对每一个下标，向两侧寻找回文子串，并且记录最长值并更新（可以根据中心点和长度计算得出起末）
// 奇数长度子串的中心点是一个字符，偶数长度子串的中心点是两个字符
// 因此在从每一个下标向两侧扩散的时候，要比较单中心店和双中心点的情况

func longestPalindrome(s string) string {
    if len(s) <= 1 {
        return s
    }

    start := 0
    end := 0

    for i:=0; i<len(s); i++ {
        len1 := extendPalindromString(s, i, i) // 奇数
        len2 := 0
        if i != len(s)-1 {
            len2 = extendPalindromString(s, i, i+1) // 偶数
        }

        if len1 > len2 && len1 > end-start+1 {
            // 新的奇数长度回文子串成为最长子串
            // 更新end和start
            start = i-(len1-1)/2
            end = i+(len1-1)/2
        } else if len2 > len1 && len2 > end-start+1 {
            // 新的偶数长度回文子串成为最长子串
            // 更新end和start
            start = i-(len2/2-1)
            end = i+1+(len2/2-1) // 这里要注意不是i+(len2/2-1)
        }
    }

    return s[start:end+1]
}

// 从给定中心点向两边找回文子串
func extendPalindromString(s string, c1 int, c2 int) int {
    for  ; c1>=0 && c2 < len(s) && s[c1] == s[c2] ; {
        c1 -= 1
        c2 += 1
    }

    // 结束循环时，回文子串如果有，就是c1+1到c2-1，长度是c2-1-(c1+1)+1=c2-c1-1
    if s[c1+1] == s[c2-1]{
        return c2-c1-1
    } else {
        return 0
    }
}

// *************************************************
// 参考deepseek的优化，简化一些不必要的判断
// *************************************************

// 中心扩散法
// 一个循环，针对每一个下标，向两侧寻找回文子串，并且记录最长值并更新（可以根据中心点和长度计算得出起末）
// 奇数长度子串的中心点是一个字符，偶数长度子串的中心点是两个字符
// 因此在从每一个下标向两侧扩散的时候，要比较单中心店和双中心点的情况

func longestPalindrome(s string) string {
    if len(s) <= 1 {
        return s
    }

    start := 0
    end := 0

    for i:=0; i<len(s); i++ {
        len1 := expandAroundCenter(s, i, i) // 奇数
        len2 := 0
        if i != len(s)-1 {
            len2 = expandAroundCenter(s, i, i+1) // 偶数
        }

        if len1 > len2 && len1 > end-start+1 {
            // 新的奇数长度回文子串成为最长子串
            // 更新end和start
            start = i-(len1-1)/2
            end = i+(len1-1)/2
        } else if len2 > len1 && len2 > end-start+1 {
            // 新的偶数长度回文子串成为最长子串
            // 更新end和start
            start = i-(len2/2-1)
            end = i+1+(len2/2-1) // 这里要注意不是i+(len2/2-1)
        }
    }

    return s[start:end+1]
}

// 从给定中心点向两边找回文子串
func expandAroundCenter(s string, left int, right int) int {
    for left>=0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }

    // 思考一下，其实可以直接返回right-left-1
    return right-left-1
}