package cdyleetcode

// *************************************************
// 自己一开始的解法
// *************************************************
func lengthOfLongestSubstring(s string) int {
    maxLength := 0
    lastPos := 0
    resultSet := make(map[string]int)
    runes := []rune(s)
    for left, char := range(s) {
        if left-1 >= 0 {
            delete(resultSet, string(runes[left-1]))
        }
        resultSet[string(char)] = left
        if lastPos<left {
            lastPos=left
        }
        for right:=lastPos+1; right<len(s); right++ {
            rightChar := string(runes[right])
            // 当前字符已经存在于子串里了
            if _, exists := resultSet[rightChar]; exists {
                break
            } else {
                resultSet[rightChar] = right
                lastPos = right
            }
        }
        if len(resultSet) > maxLength {
            maxLength = len(resultSet)
        }
    }
    return maxLength
}

// 关键是找每一个字符开始，找这个字符开始的最长不重复子串（最直观）
// 而不是找a开头的最长不重复子串、b开头的最长不重复子串...
// 前一个字符的最长不重复子串的末尾，肯定也包含在下一个字符的最长不重复子串中，所以要记忆这个末位
// resultSet其实可以保留，只要在移到下一位的时候，去掉上一位

// 如何处理string的每一位

// *************************************************
// chatgpt优化后的代码
// *************************************************
func lengthOfLongestSubstring(s string) int {
    charIndexMap := make(map[rune]int)
    maxLength := 0
    left := 0

    for right, char := range s {
        // 如果字符已存在并且在当前窗口内，则移动左边界
        if lastIndex, exists := charIndexMap[char]; exists && lastIndex >= left {
            left = lastIndex + 1
        }
        charIndexMap[char] = right
        if right-left+1 > maxLength {
            maxLength = right - left + 1
        }
    }

    return maxLength
}



// *************************************************
// 使用一次循环
// *************************************************
func lengthOfLongestSubstring(s string) int {
    left := 0
    maxLength := 0
    charIndexMap := make(map[rune]int)
    sRune := []rune(s)
    for right, char := range(sRune) {
        if right < len(s) && left <= right{
            if index, exists := charIndexMap[char]; exists && index >= left {
                // map中已经有当前查看的值了
                // left指针移到已存在的值的后一个
                left = index + 1
                charIndexMap[char] = right
            } else {
                // 将right对应的char添加入map
                charIndexMap[char] = right
            }
            if right-left+1 > maxLength {
                maxLength = right-left+1
            }
        } else {
            break
        }
    }
    return maxLength
}

// 借助右指针进行遍历