package cdyleetcode

// *************************************************
// 自己写的二分法，最后return不是简单的return mid，而是要分类讨论
// *************************************************
func searchInsert(nums []int, target int) int {
    n := len(nums)
    left := 0
    right := n-1
    var mid int
    for left <= right {
        mid = (right+left+1)/2
        if nums[mid] == target {
            return mid
        }

        if nums[mid] > target {
            // 说明target在mid的左边
            right = mid-1
        } else {
            // 说明target在mid的右边
            left = mid+1
        }
    }

    if left > right {
        // 最后一次查找时，target在mid右边
        return left
    } else {
        // 最后一次查找时，target在mid左边
        return right
    }
}

// 二分查找
// 一个for循环：先确认left, mid, right

// *************************************************
// deepseek的优化，更简洁
// *************************************************
func searchInsert(nums []int, target int) int {
    n := len(nums)
    left := 0
    right := n-1
    var mid int
    for left <= right {
        mid = left + (right-left)/2
        if nums[mid] == target {
            return mid
        }

        if nums[mid] > target {
            // 说明target在mid的左边
            right = mid-1
        } else {
            // 说明target在mid的右边
            left = mid+1
        }
    }

    return left
}

// 二分查找
// 一个for循环：先确认left, mid, right
