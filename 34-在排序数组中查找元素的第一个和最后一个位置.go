package cdyleetcode

// *************************************************
// 自己写的，找第一个比target小和大的数，但是出错很多：
// 1. 当 nums[mid] < target 时，应该继续向右找，而不是立即判断并返回
// 2. 在找左右边界时，应该用标准的二分查找模板，而不是提前 break。
// 3. 边界条件处理不当：当 target 在数组最左边时，找不到比 target 小的数。当 target 在数组最右边时，找不到比 target 大的数
// *************************************************
func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }

    left := 0
    right := len(nums)-1

    targetLeft := -1

    // 找第一个比target小的数
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] >= target {
            // 当前值太大了，要向左找
            right = mid -1
        } else {
            // 找到了第一个比target小的数
            if mid == len(nums)-1 || nums[mid+1] != target {
                return []int{-1, -1}
            }
            targetLeft = mid+1
            break
        }
    }

    // 没找到比target小的数
    if left > right && targetLeft == -1{
        return []int{-1, -1}
    } else if left > right && nums[0] == target {
        targetLeft = -1
    }

    // 找第一个比target大的数
    left = targetLeft + 1
    targetRight := left

    right = len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] <= target {
            left = mid + 1
        } else {
            // 找到了第一个比target大的数
            if mid == 0 || nums[mid-1] != target {
                return []int{-1, -1}
            }
            targetRight = mid - 1
            break
        }
    }

    // 没找到比target大的数
    if left > right && targetRight == targetLeft + 1{
        return []int{targetLeft+1, targetRight}
    }

    return []int{targetLeft+1, targetRight-1}

}

// 3 4 6 7 9 9 9 9 9 9 10

// 找到第一个比target小的位置i，看这个位置后一位是不是target
//   - 如果不是，说明不存在，返回[-1, -1]
//   - 如果是，就找第一个比target大的位置j，然后返回[i+1, j-1]

// *************************************************
// deepseek的修改
// *************************************************
func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }

    left := 0
    right := len(nums)-1
    targetLeft := -1

    // 找第一个比target小的数
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] >= target {
            // 当前值 >= target，要向左找
            right = mid - 1
        } else {
            // nums[mid] < target，记录这个位置
            targetLeft = mid
            left = mid + 1  // 继续向右找更大的小于target的数
        }
    }

    // 检查target是否存在
    if targetLeft+1 >= len(nums) || nums[targetLeft+1] != target {
        return []int{-1, -1}
    }

    // 找第一个比target大的数
    left = targetLeft + 1
    right = len(nums)-1
    targetRight := targetLeft + 1  // 先初始化为target的左边界

    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] <= target {
            // 当前值 <= target，要向右找
            left = mid + 1
            targetRight = mid  // 记录最后一个等于target的位置
        } else {
            // 找到了第一个比target大的数
            right = mid - 1
        }
    }

    // targetRight是最后一个等于target的位置
    return []int{targetLeft + 1, targetRight}
}

// *************************************************
// 直接看灵神的解法，找第一个满足nums[i]>=target的下标
// *************************************************

// 返回第一个满足nums[i]>=target的下标
func lowerBound(nums []int, target int) int {
    left := 0
    right := len(nums)-1

    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] >= target { // mid满足条件，找更小的
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    // right是在mid满足条件时更新的，所以right+1就是>=target的下标
    // 而循环退出时，right+1正是left

    return left
}

func searchRange(nums []int, target int) []int {
    start := lowerBound(nums, target)
    if start == len(nums) || nums[start] != target {
        return []int{-1, -1}
    }

    end := lowerBound(nums, target+1) - 1

    return []int{start, end}

}

// 灵神的解法：lowerBound函数表示找到第一个满足nums[i]>=target的下标i
// 如果数组为空，或者所有元素都小于target，就返回len(nums)
// 先调用lowerBound(target)，如果下标>len(nums)-1，或者这个nums[i]不等于target，说明target不存在于nums中，返回[-1,-1]。记录返回值为start。
// 再调用lowerBound(target+1)。如果下标>len(nums)-1（其实就是下标为len(nums))，说明大于等于target+1的数字不存在，而lowerBound(target)存在，那么end就是len(nums)-1，也就是lowerBound(target+1)-1
// 而如果lowerBound(target+1)的下标i合法，说明i-1就是target序列的end。这个值也就是lowerBound(target+1)-1
