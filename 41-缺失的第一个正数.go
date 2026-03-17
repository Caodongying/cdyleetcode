package cdyleetcode

// *************************************************
// 自己做的暴力解法：先排序，再找到第一个非负数
// *************************************************
import "sort"

func firstMissingPositive(nums []int) int {
    sort.Ints(nums)

    firstIndex := -1

    for i := 0; i<len(nums); i++ {
        if nums[i] > 0 {
            firstIndex = i
            break
        }
    }

	// 比如测试用例[0]
    if firstIndex == -1 {
        return 1
    }

    if nums[firstIndex] > 1 {
        return 1
    }
// 0 1 1 2 2
    sequenceCount := 1
    for i := firstIndex + 1; i<len(nums); i++{
        expected := nums[firstIndex] + sequenceCount
        if nums[i] < expected {
            continue
        } else if nums[i] == expected {
			// 比如测试用例[0, 1, 1, 2, 2]
            sequenceCount++
        } else {
			// 比如测试用例[0, 1, 1, 2, 4]
            return expected
        }

    }

    return nums[len(nums)-1]+1

}

// 暴力解法：先排序，再找到第一个非负数
// 如果这个非负数大于1，则返回1
// 如果这个非负数等于1，就一直+1寻找

// *************************************************
// 原地哈希，超时，[1,1]案例没过
// *************************************************
func firstMissingPositive(nums []int) int {
    for i:=0; i<len(nums); i++ {
        for nums[i] > 0 && nums[i] <= len(nums) -1 && nums[i] != i+1{
            targetIndex := nums[i]-1
            iValue := nums[i]
            nums[i] = nums[targetIndex]
            nums[targetIndex] = iValue
        }
    }

    // 寻找结果
    for i:=0; i<len(nums); i++ {
        if nums[i] != i+1 {
            return i+1
        }
    }
    return nums[len(nums)-1]+1
}


// 使用原地哈希
// 可以确定的是，对于长度为n的数组，缺失的第一个正数的取值范围一定是[1,n+1]
// 在理想情况下，如果数组是[1,2,3,4,5,x,7]，也就是nums[i]=i+1，那么第一个不满足这个条件的下标i，就是我们要找的位置。返回i+1
// 如果遍历完数组，还没有找到，就说明答案是最后一个值+1
// 为了朝理想情况靠拢，并且保持O(1)的空间复杂度，就可以通过交换元素来实现
// 一次遍历，对于每一个下标i，如果nums[i]<=0，就不管，到下一个；如果nums[i]>len(nums)-1，就跳过
// 否则，交换nums[i]与下标为nums[i]-1处的值，然后接着对nums[i]的值进行交换

// *************************************************
// 判断影子分身的时候，错用continue
// *************************************************
func firstMissingPositive(nums []int) int {
    for i:=0; i<len(nums); i++ {
        for nums[i] > 0 && nums[i] <= len(nums) && nums[i] != i+1{
            targetIndex := nums[i]-1
            iValue := nums[i]
            if nums[targetIndex] == iValue {
                continue // [1, 1]，影子分身
            }
            nums[i] = nums[targetIndex]
            nums[targetIndex] = iValue
        }
    }

    // 寻找结果
    for i:=0; i<len(nums); i++ {
        if nums[i] != i+1 {
            return i+1
        }
    }
    return nums[len(nums)-1]+1
}


// 使用原地哈希
// 可以确定的是，对于长度为n的数组，缺失的第一个正数的取值范围一定是[1,n+1]
// 在理想情况下，如果数组是[1,2,3,4,5,x,7]，也就是nums[i]=i+1，那么第一个不满足这个条件的下标i，就是我们要找的位置。返回i+1
// 如果遍历完数组，还没有找到，就说明答案是最后一个值+1
// 为了朝理想情况靠拢，并且保持O(1)的空间复杂度，就可以通过交换元素来实现
// 一次遍历，对于每一个下标i，如果nums[i]<=0，就不管，到下一个；如果nums[i]>len(nums)-1，就跳过
// 否则，交换nums[i]与下标为nums[i]-1处的值，然后接着对nums[i]的值进行交换

// *************************************************
// 正确使用break，但是性能差
// *************************************************
func firstMissingPositive(nums []int) int {
    for i:=0; i<len(nums); i++ {
        for nums[i] > 0 && nums[i] <= len(nums) && nums[i] != i+1{
            targetIndex := nums[i]-1
            iValue := nums[i]
            if nums[targetIndex] == iValue {
                break // [1, 1]，影子分身
            }
            nums[i] = nums[targetIndex]
            nums[targetIndex] = iValue
        }
    }

    // 寻找结果
    for i:=0; i<len(nums); i++ {
        if nums[i] != i+1 {
            return i+1
        }
    }
    return nums[len(nums)-1]+1
}

// *************************************************
// 根据灵神解法，优化循环条件
// *************************************************
func firstMissingPositive(nums []int) int {
    for i:=0; i<len(nums); i++ {
        for nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i]{
            targetIndex := nums[i]-1
            iValue := nums[i]
            nums[i] = nums[targetIndex]
            nums[targetIndex] = iValue
        }
    }

    // 寻找结果
    for i:=0; i<len(nums); i++ {
        if nums[i] != i+1 {
            return i+1
        }
    }
    return len(nums)+1
}

