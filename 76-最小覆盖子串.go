package main

// *************************************************
// 看了b站的讲解，自己写的，有点不太顺利，但是还行，用了vscode自己调试
// *************************************************
func minWindow(s string, t string) string {
    if len(s) < len(t) {
        return ""
    }

    left, right := 0, 0

    countT := make(map[string]int, 0)
    for _, char := range(t) {
        countT[string(char)]++
    }

    countS := make(map[string]int, 0)

    if _, exists := countT[string(s[0])]; exists {
        countS[string(s[0])]++
    }


    minLen := 0
    result := ""

    for left <= right && right < len(s) {
        for cover(countS, countT) {
            // 记录当前结果
            if (right-left+1 < minLen) || (right-left+1 > 0 && minLen == 0) {
                minLen = right-left+1
                result = s[left:right+1]
            }

            // 开始左指针收缩
            left++
            if left > right {
                break
            }
            // 更新窗口
            if _, exists := countT[string(s[left-1])]; exists {
                countS[string(s[left-1])]--
            }
            // 下面的不能加，是导致错误的原因。left负责的是缩减，下面的代码会导致元素被重复添加
            // if _, exists := countT[string(s[left])]; exists {
            //     countS[string(s[left])]++
            // }
        }

        // 窗口不包含t，右指针向右扩展
        right++
        if right == len(s) {
            break
        }
        if _, exists := countT[string(s[right])]; exists { // string[i]分中英文
            countS[string(s[right])]++
        }

    }

    return result

}

// 判断m1是否覆盖m2
func cover(m1, m2 map[string]int) bool {
    if len(m1) != len(m2) {
        return false
    }

    for k, v := range m2 {
        if m1[k] < v {
            return false
        }
    }

    return true
}

// 看这个视频就理解了：https://www.bilibili.com/video/BV1m4isBBE69/?spm_id_from=333.337.search-card.all.click&vd_source=e3d59376cbbf7e58fec4949fdb3a00d4

// right指针用于扩
// left指针用于缩
// 当窗口符合条件时，left向右移动进行锁，一旦窗口又不符合了，立刻right向右扩
// left和right初始化时下标都是0

// *************************************************
// 看了Gemini优化后自己写的，先判断是否包含再扩右指针，有点容易错
// *************************************************
func minWindow(s string, t string) string {
    if len(s) < len(t) {
        return ""
    }

    // 使用数组而不是map
    need := [128]int{}
    window := [128]int{}

    uniqueChars := 0

    for i:=0; i<len(t); i++ {
        if need[t[i]] == 0 {
            uniqueChars++
        }
        need[t[i]]++
    }

    left, right := 0, 0
    validCount := 0 //记录当前窗口满足了多少种字符的要求
    start := 0
    minLen := len(s)+1

    if need[s[0]] > 0 {
        window[s[0]]++

        if window[s[0]] == need[s[0]] {
            validCount++
        }
    }

    for left <= right && right < len(s) {
        // 判断是否满足子串要求
        for validCount == uniqueChars {
            // 判断是否为当前最短子串
            if right-left+1 < minLen {
                start = left
                minLen = right-left+1
            }
            // 缩
            charLeft := s[left]
            left++
            // 缩出去的值是t中的
            if need[charLeft] > 0 {
                window[charLeft]--
                if window[charLeft] < need[charLeft] {
                    validCount--
                }
            }
        }

        // 右指针扩
        right++
        if right == len(s) {
            break
        }
        if need[s[right]] > 0 {
            window[s[right]]++
            if window[s[right]] == need[s[right]]{
                validCount++
            }
        }

    }

    if minLen == len(s)+1 {
        return ""
    }
    return s[start:start+minLen]

}

// right指针用于扩
// left指针用于锁
// 当窗口符合条件时，left向右移动进行锁，一旦窗口又不符合了，立刻right向右扩
// left和right初始化时下标都是0
