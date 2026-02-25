package cdyleetcode

// *************************************************
// 初始笨方法
// *************************************************
import (
    "sort"
)

type NumWithIndex struct {
    Value int
    Index int
}

func twoSum(nums []int, target int) []int {
    // 创建包含值和索引的结构体数组
    numsWithIndices := make([]NumWithIndex, len(nums))
    for i, num := range nums {
        numsWithIndices[i] = NumWithIndex{Value: num, Index: i}
    }

    // 按值升序排序
    sort.Slice(numsWithIndices, func(i, j int) bool {
        return numsWithIndices[i].Value < numsWithIndices[j].Value
    })

    left, right := 0, len(numsWithIndices)-1
    for left < right {
        sum := numsWithIndices[left].Value + numsWithIndices[right].Value
        if sum == target {
            // 返回原始索引，确保小的索引在前
            if numsWithIndices[left].Index < numsWithIndices[right].Index {
                return []int{numsWithIndices[left].Index, numsWithIndices[right].Index}
            } else {
                return []int{numsWithIndices[right].Index, numsWithIndices[left].Index}
            }
        } else if sum < target {
            left++
        } else {
            right--
        }
    }

    return nil
}

// *************************************************
// 暴力解法
// *************************************************
func twoSum(nums []int, target int) []int {
    for i:=0; i<len(nums); i++{
        for j:=i+1; j<len(nums); j++ {
            if (nums[i] + nums[j] == target) {
                return []int{i,j}
            }
            continue
        }
    }
    return nil
}

// *************************************************
// 将数组全部添加进map的哈希表法
// *************************************************
func twoSum(nums []int, target int) []int {
    // 构造一个哈希表（键：nums元素，值：nums下标）
    numMap := make(map[int]int)
    for i:=0; i<len(nums); i++ {
        numMap[nums[i]] = i
    }
    // 单遍历，进行查找
    for j:=0; j<len(nums);j++ {
        valueJ := nums[j]
        valueRemaining := target-valueJ
        if anotherIndex, exists := numMap[valueRemaining]; !exists {
            continue
        } else {
            if anotherIndex != j {
                return []int{j, anotherIndex}
            } else {
                continue
            }
        }
    }
    return nil
}


// *************************************************
// 改进版的哈希表法
// *************************************************
func twoSum(nums []int, target int) []int {
    // 构造一个哈希表（键：nums元素，值：nums下标）
    numMap := make(map[int]int)
    for i:=0; i<len(nums); i++ {
        // 另一个值存在于map中，直接返回
        if anotherIndex, exists := numMap[target-nums[i]]; exists {
            if i != anotherIndex {
                return []int{i, anotherIndex}
            }
        } else {
            //把当前值加入map
            numMap[nums[i]] = i
        }
    }
    return nil
}