package cdyleetcode

// *************************************************
// ❌❌不能在数组排序后立刻删除重复元素
// *************************************************
import (
    "sort"
)

func threeSum(nums []int) [][]int {
    result := make([][]int, 0)

    // 对nums进行排序
    sort.Ints(nums) // 不是nums=sort.Ints(nums)
    uniqueNums := []int{nums[0]}
    j := 0
    // 去掉重复值
    for i:=1; i < len(nums); i++ {
        if nums[i] != uniqueNums[j] {
            j++
            uniqueNums=append(uniqueNums, nums[i])
        }
    }


    for i:=0; i < len(uniqueNums); i++ {
        left := i + 1
        right := len(uniqueNums)-1

        for left < right {
            sum := uniqueNums[i] + uniqueNums[left] + uniqueNums[right]
            if sum == 0 {
                current := []int{uniqueNums[i], uniqueNums[left], uniqueNums[right]}
                result = append(result, current)
                left++
                right--
            } else if sum > 0 {
                right--
            } else {
                left++
            }
        }
    }

    return result
}

// 先对nums进行升序排序
// 从第一个元素开始遍历(i=0)。对于每个元素，left取其后面的第一个元素下标，right取最后一个
// 在对i的遍历中，i是固定的，因此是寻找left和right。
// 如果nums[i]+nums[left]+nums[right]==0，继续寻找，移动left，直到移动的值不等于刚刚匹配的nums[left]，来避免重复
// 如果nums[i]+nums[left]+nums[right]>0，说明值要小一点，就right--
// 如果nums[i]+nums[left]+nums[right]<0，说明值要大一点，就left++

// *************************************************
// 看了代码随想录后自己写的，双指针，在移动指针的时候有较多判断
// *************************************************
import (
    "sort"
)

func threeSum(nums []int) [][]int {
    result := make([][]int, 0)

    // 对nums进行排序
    sort.Ints(nums) // 不是nums=sort.Ints(nums)


    for i:=0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        left := i + 1
        right := len(nums)-1

        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            currentLeftValue := nums[left]
            currentRightValue := nums[right]
            if sum == 0 {
                current := []int{nums[i], nums[left], nums[right]}
                result = append(result, current)
                left++
                for left < right {
                    if nums[left] == currentLeftValue {
                        left++
                    } else {
                        break
                    }
                }

                right--
                for right > left {
                    if nums[right] == currentRightValue {
                        right--
                    } else {
                        break
                    }
                }
            } else if sum > 0 {
                right--
                for right > left {
                    if nums[right] == currentRightValue {
                        right--
                    } else {
                        break
                    }
                }
            } else {
                left++
                for left < right {
                    if nums[left] == currentLeftValue {
                        left++
                    } else {
                        break
                    }
                }
            }
        }
    }

    return result
}

// *************************************************
// 看了代码随想录后的小小优化
// *************************************************
import (
    "sort"
)

func threeSum(nums []int) [][]int {
    result := make([][]int, 0)

    // 对nums进行排序
    sort.Ints(nums) // 不是nums=sort.Ints(nums)


    for i:=0; i < len(nums)-2; i++ {
        if nums[i] > 0 { // 优化一下
            break
        }
        if i > 0 && nums[i] == nums[i-1] { // 为了去重
            continue
        }

        left := i + 1
        right := len(nums)-1

        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            currentLeftValue := nums[left]
            currentRightValue := nums[right]
            if sum == 0 {
                current := []int{nums[i], nums[left], nums[right]}
                result = append(result, current)
                for left<right && nums[left]==currentLeftValue{
                    left++
                }
                for left<right && nums[right]==currentRightValue{
                    right--
                }

            } else if sum > 0 {
                right--
            } else {
                left++
            }
        }
    }

    return result
}


// *************************************************
// 看了灵神题解，再优化一次
// *************************************************
import (
    "sort"
)

func threeSum(nums []int) [][]int {
    result := make([][]int, 0)

    // 对nums进行排序
    sort.Ints(nums) // 不是nums=sort.Ints(nums)


    for i:=0; i < len(nums)-2; i++ {
        if nums[i] > 0 { // 优化一下
            break
        }
        if nums[i]+nums[i+1]+nums[i+2] > 0 { // 继续优化
            break
        }

        if i > 0 && nums[i] == nums[i-1] { // 为了去重
            continue
        }

        left := i + 1
        right := len(nums)-1

        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            currentLeftValue := nums[left]
            currentRightValue := nums[right]
            if sum == 0 {
                current := []int{nums[i], nums[left], nums[right]}
                result = append(result, current)
                for left<right && nums[left]==currentLeftValue{
                    left++
                }
                for left<right && nums[right]==currentRightValue{
                    right--
                }

            } else if sum > 0 {
                right--
            } else {
                left++
            }
        }
    }

    return result
}

