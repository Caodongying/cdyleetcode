package cdyleetcode

// *************************************************
// 自己写的，先排序，然后一次遍历。算是暴力写法， 有点磕绊，用vscode debug了
// *************************************************
func merge(intervals [][]int) [][]int {
    if len(intervals) == 1 {
        return intervals
    }

    result := [][]int{} // append在这种情况下怎么用的

    // 先对intervals进行排序
    intervalSort(intervals)

    // 先处理前两个intervals
    result = append(result, mergeTwoIntervals(intervals[0], intervals[1])...)
    // 一次循环，进行合并
    for i:=2; i<len(intervals); i++ {
        end := result[len(result)-1]
        temp := mergeTwoIntervals(end, intervals[i])
        result = result[:len(result)-1]
        result = append(result, temp...)
    }

    return result

}

func mergeTwoIntervals(left []int, right []int) [][]int {
    result := [][]int{}
    if left[0] == right[0] {
        maxRight := max(left[1], right[1])
        result = append(result, []int{left[0], maxRight})
    } else if left[1] < right[1] && left[1] >= right[0] {
        result = append(result, []int{left[0], right[1]})
    } else if left[1] < right[0]{
        result = append(result, left, right)
    } else {
        result = append(result, left)
    }
    return result
}

func intervalSort(intervals [][]int) {
    // 冒泡排序
    for i:=0; i<len(intervals); i++ {
        for j:=0; j<len(intervals)-1-i; j++ {
            if intervals[j][0] > intervals[j+1][0] {
                temp := intervals[j+1]
                intervals[j+1] = intervals[j]
                intervals[j] = temp
            }
        }
    }
}
// 先按照intervals[i][0]进行升序排序
// 再用一次循环进行合并
// 每次合并只需要考虑当前interval和结果集

// *************************************************
// 自己看着Gemini优化写的
// *************************************************
import (
    "sort"
)

func merge(intervals [][]int) [][]int {
    if len(intervals) == 1 {
        return intervals
    }

    // 使用标准库进行排序
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    // 初始化结果集，并放入第一个区间
    result := [][]int{intervals[0]}

    // 一次循环，进行合并
    for i:=1; i<len(intervals); i++ {
        last := result[len(result)-1]
        current := intervals[i]

        // 如果当前两区间重叠，就替换last的右端点
        // 重叠：左区间的right端点，大于等于右区间的left端点
        if last[1] >= current[0] {
            last[1] = max(last[1], current[1])
        } else {
            // 如果当前两区间不重叠，current直接加入结果集
            result = append(result, current)
        }

    }

    return result

}


// *************************************************
// 进一步优化，预先指定容量
// *************************************************
import (
    "sort"
)

func merge(intervals [][]int) [][]int {
    if len(intervals) == 1 {
        return intervals
    }

    // 使用标准库进行排序
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    // 初始化结果集，并放入第一个区间
    result := make([][]int, 0, len(intervals))
    result = append(result, intervals[0])

    // 一次循环，进行合并
    for i:=1; i<len(intervals); i++ {
        last := result[len(result)-1]
        current := intervals[i]

        // 如果当前两区间重叠，就替换last的右端点
        // 重叠：左区间的right端点，大于等于右区间的left端点
        if last[1] >= current[0] {
            last[1] = max(last[1], current[1])
        } else {
            // 如果当前两区间不重叠，current直接加入结果集
            result = append(result, current)
        }

    }

    return result

}
