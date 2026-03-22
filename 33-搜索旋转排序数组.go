package cdyleetcode

// *************************************************
// ❌❌ 没有考虑到旋转点为0的情况，也没有考虑到target不存在的情况
// *************************************************
func search(nums []int, target int) int {
    n := len(nums)
    if n == 0 || (n==1 && nums[0]!= target) {
        return -1
    }

    if n==1 && nums[0] == target {
        return 0
    }

    // 找到第一个<nums[0]的数组下标i
    left, right := 0, n-1
    rotateIdx := -1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] < nums[0] {
            // mid满足条件，但不可以立刻break
            right = mid-1
        } else {
            left = mid + 1
        }
    }

    // 此时left = right+1，mid = right+1, left = mid
    rotateIdx = left


    if rotateIdx == n {
        // 说明没有发生旋转，也就是旋转点即0
        left, right = 0, n-1
    } else if target >= nums[0] && target <= nums[rotateIdx-1] {
        left, right = 0, rotateIdx-1
    } else if target >= nums[rotateIdx] && target <= nums[n-1] {
        left, right = rotateIdx, n-1
    } else {
        return -1
    }

    // 二分查找值

    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return -1
}

// 先用二分法找到第一个<nums[0]的数组下标i
// i就是旋转点，nums[i]是最小值，nums[i-1]（i不等于0）是最大值
// 继续使用二分查找
// 如果target小于nums[i]，返回-1
// nums[0]到nums[i-1]闭区间，单调增。
// nums[i]到nums[len(nums)-1]闭区间，单调增

// *************************************************
// ❌❌ 没有考虑到旋转点为0的情况
// *************************************************
func search(nums []int, target int) int {
    n := len(nums)
    if n == 0 || (n==1 && nums[0]!= target) {
        return -1
    }

    if n==1 && nums[0] == target {
        return 0
    }

    // 找到第一个<nums[0]的数组下标i
    left, right := 0, n-1
    rotateIdx := -1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] < nums[0] {
            // mid满足条件，但不可以立刻break
            right = mid-1
        } else {
            left = mid + 1
        }
    }

    // 此时left = right+1，mid = right+1, left = mid
    rotateIdx = left


    if rotateIdx == n {
        // 说明没有发生旋转，也就是旋转点即0
        left, right = 0, n-1
    } else if target >= nums[0] && target <= nums[rotateIdx-1] {
        left, right = 0, rotateIdx-1
    } else if target >= nums[rotateIdx] && target <= nums[n-1] {
        left, right = rotateIdx, n-1
    } else {
        return -1
    }

    // 二分查找值

    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return -1
}

// 先用二分法找到第一个<nums[0]的数组下标i
// i就是旋转点，nums[i]是最小值，nums[i-1]（i不等于0）是最大值
// 继续使用二分查找
// 如果target小于nums[i]，返回-1
// nums[0]到nums[i-1]闭区间，单调增。
// nums[i]到nums[len(nums)-1]闭区间，单调增

// *************************************************
// 自己写的，两次二分法，第一次找旋转点，第二次找target
// *************************************************
func search(nums []int, target int) int {
    n := len(nums)
    if n == 0 || (n==1 && nums[0]!= target) {
        return -1
    }

    if n==1 && nums[0] == target {
        return 0
    }

    // 找到第一个<nums[0]的数组下标i
    left, right := 0, n-1
    rotateIdx := -1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] < nums[0] {
            // mid满足条件，但不可以立刻break
            right = mid-1
        } else {
            left = mid + 1
        }
    }

    // 此时left = right+1，mid = right+1, left = mid
    rotateIdx = left


    if rotateIdx == n {
        // 说明没有发生旋转，也就是旋转点即0
        left, right = 0, n-1
    } else if target >= nums[0] && target <= nums[rotateIdx-1] {
        left, right = 0, rotateIdx-1
    } else if target >= nums[rotateIdx] && target <= nums[n-1] {
        left, right = rotateIdx, n-1
    } else {
        return -1
    }

    // 二分查找值

    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return -1
}

// 先用二分法找到第一个<nums[0]的数组下标i
// i就是旋转点，nums[i]是最小值，nums[i-1]（i不等于0）是最大值
// 继续使用二分查找
// 如果target小于nums[i]，返回-1
// nums[0]到nums[i-1]闭区间，单调增。
// nums[i]到nums[len(nums)-1]闭区间，单调增

// *************************************************
// 不必单独判断一开始的len为0和1的情况
// *************************************************
func search(nums []int, target int) int {
    n := len(nums)

    // 找到第一个<nums[0]的数组下标i
    left, right := 0, n-1
    rotateIdx := -1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] < nums[0] {
            // mid满足条件，但不可以立刻break
            right = mid-1
        } else {
            left = mid + 1
        }
    }

    // 此时left = right+1，mid = right+1, left = mid
    rotateIdx = left


    if rotateIdx == n {
        // 说明没有发生旋转，也就是旋转点即0
        left, right = 0, n-1
    } else if target >= nums[0] && target <= nums[rotateIdx-1] {
        left, right = 0, rotateIdx-1
    } else if target >= nums[rotateIdx] && target <= nums[n-1] {
        left, right = rotateIdx, n-1
    } else {
        return -1
    }

    // 二分查找值

    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return -1
}

// 先用二分法找到第一个<nums[0]的数组下标i
// i就是旋转点，nums[i]是最小值，nums[i-1]（i不等于0）是最大值
// 继续使用二分查找
// 如果target小于nums[i]，返回-1
// nums[0]到nums[i-1]闭区间，单调增。
// nums[i]到nums[len(nums)-1]闭区间，单调增

// *************************************************
// Deepseek优化：一次二分查找
// *************************************************
func search(nums []int, target int) int {
    n := len(nums)

    left, right := 0, n-1

    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        }

        if nums[mid] >= nums[left]{ // mid在左段
            // [left, mid]单调增
            if target >= nums[left] && target <= nums[mid] { // target在[left, mid]内
                right = mid-1
            } else {
                left = mid+1
            }
        } else {
            // mid在右段，[mid, right]单调增
            if target >= nums[mid] && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid -1
            }
        }
    }

    return -1
}

// 一次二分法
// 原理：对于区间[left, right]，必定可以分为两段（可能第二段为空），这两段分别单调增，并且左段大于右段。段的分界点就是旋转点。当然如果right和left也可能处于旋转点同一侧
//  nums[mid]必然要么落在左段，要么落在右段。左段左端点已知，右段右端点已知。如果nums[mid]大于等于左端点，mid就一定在左段；如果nums[mid]小于左端点，mid就一定在右段
// 然后在这两段里找target
// 如果找不到，就返回-1

// 在二分查找中，我们利用的是排除法：既然确定 target 不在 A 区间，那么它必然在 B 区间（因为整个搜索范围是 A ∪ B）。