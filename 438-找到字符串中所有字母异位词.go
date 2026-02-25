package cdyleetcode

// *************************************************
// 暴力解法
// *************************************************
func findAnagrams(s string, p string) []int {
    runeS := []rune(s)
    lengthS := len(runeS)
    lengthP := len(p)

    left := 0
    right := left + lengthP - 1

    result := make([]int, 0)

    runeP := []rune(p) // 这里直接string可以吗

    // var mapP map[rune]int
    // var mapS map[rune]int
    mapP := make(map[rune]int)


    for _, v := range(runeP) {
        mapP[v] += 1
    }

    for right < lengthS {
        // 构建mapS
        mapS := make(map[rune]int)
        for j:=left; j<=right; j++{
            mapS[runeS[j]] += 1
        }

        // 如果二者完全一致，左指针移到下一位
        if reflect.DeepEqual(mapP, mapS) == true {
            result = append(result, left)
        }

        left = left + 1
        right = left + lengthP - 1
    }

    return result
}

// 直接记住：定长的滑动窗口，长度和p长度一样
// 构造两个map，看是否一样

// 这个太复杂，暂时不做：在窗口内部，如果某个字符不在p中，那么左指针直接移动到这个字符后面的第一个在p里的字符处

// *************************************************
// deepseek的优化
// *************************************************
func findAnagrams(s string, p string) []int {
    if len(s) < len(p) {
        return []int{}
    }

    result := []int{}
    lengthP := len(p)
    lengthS := len(s)

    // 使用数组代替 map 进行频率计数（假设只包含小写字母）
    // 如果包含其他字符，可以用大小为 256 的数组（ASCII）或继续使用 map
    freqP := [26]int{}
    freqWindow := [26]int{}

    // 初始化频率计数
    for i := 0; i < lengthP; i++ {
        freqP[p[i]-'a']++
        freqWindow[s[i]-'a']++
    }

    // 比较初始窗口
    if freqP == freqWindow {
        result = append(result, 0)
    }

    // 滑动窗口
    for i := lengthP; i < lengthS; i++ {
        // 移除左边界的字符
        freqWindow[s[i-lengthP]-'a']--
        // 添加右边界的字符
        freqWindow[s[i]-'a']++

        // 比较频率数组
        if freqP == freqWindow {
            result = append(result, i-lengthP+1)
        }
    }

    return result
}

// *************************************************
// 不用map，改用array，但是arrayS还是每次重建的
// *************************************************
func findAnagrams(s string, p string) []int {
    lengthS := len(s)
    lengthP := len(p)

    left := 0
    right := left + lengthP - 1

    result := make([]int, 0)

    arrayP := [26]int{}


    for i:=0; i<lengthP; i++ {
        arrayP[p[i]-'a'] += 1
    }

    for right < lengthS {
        // 构建arrayS
        arrayS := [26]int{}
        for j:=left; j<=right; j++{
            arrayS[s[j]-'a'] += 1
        }

        // 如果二者完全一致，左指针移到下一位
        if arrayP == arrayS { // 数组是值类型，不是引用类型
            result = append(result, left)
        }

        left = left + 1
        right = left + lengthP - 1
    }

    return result
}

// *************************************************
// 自己的优化，充分利用左右边界，只更新旧的左边界和新的右边界
// *************************************************
func findAnagrams(s string, p string) []int {
    lengthS := len(s)
    lengthP := len(p)
    result := []int{}

    if lengthS < lengthP {
        return result
    }

    left := 0
    right := left + lengthP - 1


    arrayP := [26]int{}
    arrayWindow := [26]int{}


    // 初始化
    for i:=0; i<lengthP; i++ {
        arrayP[p[i]-'a'] += 1
        arrayWindow[s[i]-'a'] += 1
    }

    if arrayP == arrayWindow { // 数组是值类型，不是引用类型
        result = append(result, left)
    }

    //  根据右边界更新
    for right=right+1; right < lengthS; right++ {
        left = right - lengthP + 1
        // 更新arrayWindow
        // 将左边界的前一个值，也就是上一个左边值，从当前arrayWindow的对应镜子里减一
        arrayWindow[s[left-1]-'a'] -= 1
        // 只更新右边界
        arrayWindow[s[right]-'a'] += 1

        if arrayP == arrayWindow { // 数组是值类型，不是引用类型
            result = append(result, left)
        }
    }

    return result
}

// *************************************************
// 移除对左边界的赋值计算
// *************************************************
func findAnagrams(s string, p string) []int {
    lengthS := len(s)
    lengthP := len(p)

    if lengthS < lengthP {
        return []int{}
    }

    result := []int{}

    arrayP := [26]int{}
    arrayWindow := [26]int{}


    // 初始化
    for i:=0; i<lengthP; i++ {
        arrayP[p[i]-'a'] += 1
        arrayWindow[s[i]-'a'] += 1
    }

    if arrayP == arrayWindow { // 数组是值类型，不是引用类型
        result = append(result, 0)
    }

    //  根据右边界更新
    for right:=lengthP; right < lengthS; right++ {
        // 更新arrayWindow
        // 将左边界的前一个值，也就是上一个左边值，从当前arrayWindow的对应镜子里减一
        arrayWindow[s[right - lengthP]-'a'] -= 1
        // 只更新右边界
        arrayWindow[s[right]-'a'] += 1

        if arrayP == arrayWindow { // 数组是值类型，不是引用类型
            result = append(result, right - lengthP +1)
        }
    }

    return result
}

// *************************************************
// 官方题解
// *************************************************


func findAnagrams(s, p string) (ans []int) {
    sLen, pLen := len(s), len(p)
    if sLen < pLen {
        return
    }

    count := [26]int{}
    for i, ch := range p {
        count[s[i]-'a']++
        count[ch-'a']--
    }

    differ := 0
    for _, c := range count {
        if c != 0 {
            differ++
        }
    }
    if differ == 0 {
        ans = append(ans, 0)
    }

    for i, ch := range s[:sLen-pLen] {
        if count[ch-'a'] == 1 { // 窗口中字母 s[i] 的数量与字符串 p 中的数量从不同变得相同
            differ--
        } else if count[ch-'a'] == 0 { // 窗口中字母 s[i] 的数量与字符串 p 中的数量从相同变得不同
            differ++
        }
        count[ch-'a']--

        if count[s[i+pLen]-'a'] == -1 { // 窗口中字母 s[i+pLen] 的数量与字符串 p 中的数量从不同变得相同
            differ--
        } else if count[s[i+pLen]-'a'] == 0 { // 窗口中字母 s[i+pLen] 的数量与字符串 p 中的数量从相同变得不同
            differ++
        }
        count[s[i+pLen]-'a']++

        if differ == 0 {
            ans = append(ans, i+1)
        }
    }
    return
}
