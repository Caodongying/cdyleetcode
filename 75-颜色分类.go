package cdyleetcode

// *************************************************
// ❌❌ 思路大概正确，但是不应该用range循环，因为和2交换后的数字，可能是0可能是1，需要进一步处理，不能i++。
// 对0、1、2的处理逻辑不正确。
// *************************************************
func sortColors(nums []int)  {
    zeroPtr := 0
    twoPtr := len(nums) - 1
    for i, val := range nums {
        if val == 0 {
            if i != zeroPtr {
                // 交换
                nums[i], nums[zeroPtr] = nums[zeroPtr], val
            }
            zeroPtr++
        } else if val == 2 {
            if i > twoPtr {
                // twoPtr后面的值已经都是2了，此时循环来到twoPtr后面，说明已经排序完成
                break
            }
            if i != twoPtr {
                // 交换
                nums[i], nums[twoPtr] = nums[twoPtr], val
            }
            twoPtr--
        }
    }
}

// 一共就三个可能的数字，并且0一定是在数组最前面1组，2一定在最后面一组
// 可以设置两个指针，零指针指向“下一个可以是0的位置”，如果扫描数字是0，且零指针的已经是0，就向后移动，直到当前位置不是0或者零指针位置就是扫描数字
// 二指针同理，只不过是向前移动
// 不用单独处理1

// *************************************************
// 看了deepseek的修改
// 根据 nums[i] 的值分三种情况操作：
// 0 → 交换到前面，zeroPtr++，i++
// 2 → 交换到后面，twoPtr--，i 不动
// 1 → i++
// *************************************************
func sortColors(nums []int)  {
    zeroPtr := 0
    twoPtr := len(nums) - 1
    i := 0
    for i < len(nums) {
        if nums[i] == 0 {
            if i != zeroPtr {
                // 交换
                nums[i], nums[zeroPtr] = nums[zeroPtr], nums[i]
            }
            zeroPtr++
            i++
        } else if nums[i] == 2 {
            if i > twoPtr {
                break
            }
            if i != twoPtr {
                // 交换
                nums[i], nums[twoPtr] = nums[twoPtr], nums[i]
            }
            twoPtr--
        } else if nums[i] == 1 {
            i++
        }
    }
}

// 根据 nums[i] 的值分三种情况操作：
// 0 → 交换到前面，zeroPtr++，i++
// 2 → 交换到后面，twoPtr--，i 不动
// 1 → i++

// *************************************************
// 简化一下判断逻辑，其实就是“0和0可以交换，2和2可以交换”
// *************************************************
func sortColors(nums []int)  {
    zeroPtr := 0
    twoPtr := len(nums) - 1
    for i := 0; i <= twoPtr; {
        if nums[i] == 0 {
            nums[i], nums[zeroPtr] = nums[zeroPtr], nums[i]
            zeroPtr++
            i++
        } else if nums[i] == 2 {
            nums[i], nums[twoPtr] = nums[twoPtr], nums[i]
            twoPtr--
        } else if nums[i] == 1 {
            i++
        }
    }
}


// *************************************************
//
// *************************************************
