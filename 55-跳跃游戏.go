package cdyleetcode

// *************************************************
// 自己写的，用到max，单独考虑了0，并且只遍历到n-1位
// *************************************************
func canJump(nums []int) bool {
    maxPos := 0

    for i:=0; i<len(nums)-1; i++ { // 这里注意不可以遍历到最后一个
        if nums[i] == 0 {
            if maxPos <= i {
                // 靠前面的位能够走到当前0之后，覆盖0的影响
                return false
            }
        }
        maxPos = max(maxPos, i+nums[i])
    }

    if maxPos < len(nums) - 1 {
        return false
    } else {
        return true
    }

}

// *************************************************
// 改变循环条件为maxPos
// *************************************************
func canJump(nums []int) bool {
    maxPos := 0

    for i:=0; i<=maxPos; i++ { // 改变循环条件
        maxPos = max(maxPos, i+nums[i])
        if maxPos >= len(nums)-1 {
            return true
        }
    }

    return false

}
// *************************************************
// 还是常规循环条件，但是在里面加一个可达性判断
// *************************************************
func canJump(nums []int) bool {
    maxPos := 0

    for i, v := range nums {
        if i > maxPos { // 无法到达i
            return false
        }
        maxPos = max(maxPos, i+v)

        if maxPos >= len(nums) - 1{ // 可以跳出循环了
            break
        }
    }

    return true
}
