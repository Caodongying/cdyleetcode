package cdyleetcode

// 本题问题多多，看一下Notion里的写题记录

// *************************************************
// ❌❌想用纯粹的set作为map的键，不对，因为map的键需要可以排序，set无序。而且Golang里没有set结构。
// *************************************************
// 构造map
// key: 与字符串对应的set
// value：满足此set的具体的string的数组
// 最后返回value的数组

type Set interface{}{}

func groupAnagrams(strs []string) [][]string {
    dict := make(map[set][str])
    for str := range strs {
        strSet := stringToSet(str)
        if dict[strSet] != [] {
            dict[strSet].append[str]
        } else {
            dict[strSet] = [str]
        }
    }

    // 返回dict的所有values
    result := []
    for _, v in range dict {
        result.append(v)
    }

    return result
}

func stringToSet(str string) Set {

}

// *************************************************
// ❌❌有许多基础语法错误的排序法
// *************************************************
// 用排序后的字符串作为map的键
// 值就是[string]

import (
    "sort"
)
func groupAnagrams(strs []string) [][]string {
    dict := make(map[string][]string)

    for str := range strs {
        sorted := sort.Strings(str)
        if dict[sorted] == nil {
            dict[sorted] = make([][]string)
        }
        dict[sorted].append(str)
    }

    // 返回结果
    result := make([][]string)
    for _, v := range dict {
        result.append(v)
    }
    return result
}



// *************************************************
// 修改后的排序法
// *************************************************
// 用排序后的字符串作为map的键
// 值就是[string]

import (
    "sort"
    "strings"
)
func groupAnagrams(strs []string) [][]string {
    dict := make(map[string][]string)

    for _, str := range strs { // 不要忽略_
        chars := strings.Split(str, "")
        sort.Strings(chars) // 这个函数是对slice排序，而不是对string本身
        sorted := strings.Join(chars, "")

        // 添加到map
        dict[sorted] = append(dict[sorted], str) // append是函数
    }

    // 返回结果
    result := make([][]string, 0, len(dict)) // 初始化切片要指定长度和容量
    for _, v := range dict {
        result = append(result, v)
    }
    return result
}


// *************************************************
// 使用byte数组，减少空间复杂度
// *************************************************

// 用排序后的字符串作为map的键
// 值就是[string]

import (
    "sort"
    "strings"
)
func groupAnagrams(strs []string) [][]string {
    dict := make(map[string][]string)

    for _, str := range strs { // 不要忽略_
        // 将字符串转为字节切片并排序
        bytes := []byte(str)
        sort.Slice(bytes, func(i, j int) bool {return bytes[i]<bytes[j]})
        sortedStr := string(bytes)

        // 添加到map
        dict[sortedStr] = append(dict[sortedStr], str) // append是函数
    }

    // 返回结果
    result := make([][]string, 0, len(dict)) // 初始化切片要指定长度和容量
    for _, v := range dict {
        result = append(result, v)
    }
    return result
}

// *************************************************
// 字符计数法
// *************************************************

// 用map，键是每一个字符出现的次数，用长度为26的数组保存
// 利用char-'a'表示下标
// 数组下标为i的值，表示'a'+i所表示的字符，所出现的次数

func groupAnagrams(strs []string) [][]string {
    dict := make(map[[26]int][]string)

    for _, str := range strs {
        count := [26]int{}

        for _, char := range str {
            count[char-'a']++
        }

        dict[count] = append(dict[count], str)

    }

    results := make([][]string, 0, len(dict))

    for _, v := range dict {
        results = append(results, v)
    }

    return results
}
