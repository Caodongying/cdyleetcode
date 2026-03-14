package main

import "fmt"

func merge(intervals [][]int) [][]int {
    if len(intervals) == 1 {
        return intervals
    }

    result := [][]int{} // append在这种情况下怎么用的

    // 先对intervals进行排序
    intervalSort(intervals)

    // 先处理前两个intervals
    result = append(result, mergeTwoIntervals(intervals[0], intervals[1])...)
    fmt.Printf("前两个: %v %v 合并的结果是%v\n", intervals[0], intervals[1], result)
    // 一次循环，进行合并
    for i:=2; i<len(intervals); i++ {
        end := result[len(result)-1]
        temp := mergeTwoIntervals(end, intervals[i])
        fmt.Println(end, "和", intervals[i], "合并的结果是", temp)
        result = result[:len(result)-1]
        result = append(result, temp...)
    }

    fmt.Println("最终的结果是", result)
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
    for i:=0; i<len(intervals)-1; i++ {
        for j:=0; j<len(intervals)-1-i; j++ {
            if intervals[j][0] > intervals[j+1][0] {
                temp := intervals[j+1]
                intervals[j+1] = intervals[j]
                intervals[j] = temp
            }
        }
    }
    fmt.Println("排序后的intervals是", intervals)
}

func main() {
    intervals := [][]int{{1,3}, {2,6}, {8,10}, {15,18}}
    merge(intervals)

}
// 先按照intervals[i][0]进行升序排序
// 再用一次循环进行合并
// 每次合并只需要考虑当前interval和结果集