package main

import "fmt"
func firstMissingPositive(nums []int) int {
    for i:=0; i<len(nums); i++ {
        for nums[i] > 0 && nums[i] <= len(nums) && nums[i] != i+1{
            targetIndex := nums[i]-1
            iValue := nums[i]
            if nums[targetIndex] == iValue {
                fmt.Println("发现重复的数", iValue, "，跳过交换")
                break // [1, 1]，影子分身
            }
            nums[i] = nums[targetIndex]
            nums[targetIndex] = iValue
            fmt.Printf("交换了 %d 和 %d\n", iValue, nums[targetIndex])
        }
    }

    fmt.Println("交换后的数组是", nums)

    // 寻找结果
    for i:=0; i<len(nums); i++ {
        if nums[i] != i+1 {
            fmt.Println("第一个缺失的正数是", i+1)
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

func main() {
    nums := []int{1,1}
    firstMissingPositive(nums)
}
