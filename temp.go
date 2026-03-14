package main

// *************************************************
//
// *************************************************

import "fmt"

func minWindow(s string, t string) string {
    if len(s) < len(t) {
        return ""
    }

    left, right := 0, 0

    countT := make(map[string]int, 0)
    for _, v := range(tRune) {
        countT[string(v)]++
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
			fmt.Printf("left: %d, right: %d, countS: %v, result: %s\n", left, right, countS, result)
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

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"

	fmt.Println(minWindow(s, t))
}

// 看这个视频就理解了：https://www.bilibili.com/video/BV1m4isBBE69/?spm_id_from=333.337.search-card.all.click&vd_source=e3d59376cbbf7e58fec4949fdb3a00d4

// right指针用于扩
// left指针用于锁
// 当窗口符合条件时，left向右移动进行锁，一旦窗口又不符合了，立刻right向右扩
// left和right初始化时下标都是0

// *************************************************
//
// *************************************************

// *************************************************
//
// *************************************************

// *************************************************
//
// *************************************************
