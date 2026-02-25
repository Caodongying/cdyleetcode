package cdyleetcode

func longestConsecutive(nums []int) int {
    if len(nums) == 0{
        return 0
    }
    max := 1
    longstLenMap := make(map[int]int)
    longstLenMap[nums[0]] = 1
    for i:=1; i<len(nums); i++{
        _, ok := longstLenMap[nums[i]]
        if ok {
            // nums[i]已经处理过了
            continue
        }

        // 查看左邻居是否存在于map中
        vLeft, okLeft := longstLenMap[nums[i]-1]
        if !okLeft {
            vLeft = 0
        }

        //检查右邻居
        vRight, okRight := longstLenMap[nums[i]+1]
        if !okRight {
            vRight = 0
        }

        curLen := 1 + vLeft + vRight
        longstLenMap[nums[i]] = curLen

        // 更新最大值
        if curLen > max {
            max = curLen
        }

        // 更新左右边界
        if okLeft{
            longstLenMap[nums[i]-vLeft] = curLen
        }

        if okRight{
            longstLenMap[nums[i]+vRight] = curLen
        }


    }

    return max
}

// 构建哈希表
// key是nums里的值
// value是包含key的最长连续序列的长度
// 不重复检查值
// 如果新加入的值x，其左右值都在哈希表中，长度就是1+左value+右value
//    还得更新左右边界的value，而不是左右邻居的value。（其他顺延的怎么更新呢）
// 更新边界：由于序列是连续的，只需要更新拼接后，新序列的左右边界的长度（因为中间的数字不会再被使用，无需更新）。左边界：x - left_len   右边界：x + right_len