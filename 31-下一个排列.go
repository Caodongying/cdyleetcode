package cdyleetcode

// *************************************************
// ❌❌ 第一次写，简单地认为交换就是交换prev和last
// *************************************************
func nextPermutation(nums []int)  {
    if len(nums) == 1 {
        return
    }

    last := len(nums) - 1
    prev := len(nums) - 2

    for prev != -1 {
        if nums[prev] < nums[last] {
            break
        }
        prev--
        last--
    }

    if prev != -1 {
        nums[prev], nums[last] = nums[last], nums[prev]
    }

    sortStart := last + 1
    if prev == -1 {
        sortStart = last
    }


    // 对nums[sortStart:]之后的元素排序
    for i:=sortStart; i<len(nums)-1; i++ {
        for j:=i+1; j<len(nums); j++ {
            if nums[i] > nums[j] {
                nums[i], nums[j] = nums[j], nums[i]
            }
        }
    }
}

// 从末尾开始，比较后一位和前一位。用循环，直到后一位比前一位大
// 如果此时前一位不为-1，就交换后一位和前一位。然后把新的后一位向后的数组排序
// 如果此时前一位已经是-1，说明此时是最大序列，直接对整个数组排序
// 上面两步可以合并，就是如果前一位不为-1就交换。然后对后面的排序


// *************************************************
// 找交换点，确定和哪个数交换，然后排序
// *************************************************
func nextPermutation(nums []int)  {
    if len(nums) == 1 {
        return
    }

    last := len(nums) - 1
    prev := len(nums) - 2

    for prev != -1 {
        if nums[prev] < nums[last] {
            break
        }
        prev--
        last--
    }

    if prev != -1 {
        min := nums[last]
        minIdx := last // minIdx表示比nums[prev]大的元素的最小下标
        // 寻找minIdx
        for i:= last+1; i<len(nums); i++ {
            if nums[i]>nums[prev] && nums[i]<min {
                min = nums[i]
                minIdx = i
            }
        }

        nums[prev], nums[minIdx] = nums[minIdx], nums[prev]

    }


    // 对nums[sortStart:]之后的元素排序
    for i:=last; i<len(nums)-1; i++ {
        for j:=i+1; j<len(nums); j++ {
            if nums[i] > nums[j] {
                nums[i], nums[j] = nums[j], nums[i]
            }
        }
    }
}

// 从末尾开始，比较后一位和前一位。用循环，直到后一位比前一位大
// 如果此时前一位不为-1，说明prev就是交换点。但是要和后续数组中的大于nums[prev]的最小数交换
// 如果此时前一位已经是-1，说明此时是最大序列，直接对整个数组排序

// *************************************************
// 看了题解，根据数组降序性，优化
// *************************************************

func nextPermutation(nums []int)  {
    if len(nums) == 1 {
        return
    }

    last := len(nums) - 1
    prev := len(nums) - 2

    for prev != -1 && nums[prev] >= nums[last] {
        prev--
        last--
    }

    if prev != -1 {
        // nums[last:]已经是降序的
        // 找要交换的数
        i := len(nums) - 1
        for i>=last && nums[i] <= nums[prev] {
            i--
        }

        nums[prev], nums[i] = nums[i], nums[prev]
    }

    // 把nums[last:]翻转，因为nums[last:]依然是降序的
    slices.Reverse(nums[last:])
}

// 从末尾开始，比较后一位和前一位。用循环，直到后一位比前一位大
// 如果此时前一位不为-1，说明prev就是交换点。但是要和后续数组中的大于nums[prev]的最小数交换。因为后续数组已经是降序了，只要从后往前找到第一个比nums[prev]大的数即可。接着翻转旋转点之后的数组
// 如果此时前一位已经是-1，说明此时是最大序列，直接对整个数组排序


