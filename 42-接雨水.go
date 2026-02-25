package cdyleetcode

// *************************************************
// 大神分享的思路：将实例图中的蓝色也看成实心的，算出每一层的面积，加起来，然后减去height的总和，剩下的就是水量。
// 定义一个双指针，以及层数h=1，双指针在两头往中间移动，只要指针大于等于h，就停下来
// 目的：当两边指针都停下来的时候，计算第一层的面积（直接左指针减右指针+1），然后h++计算第二层的面积，以此类推计算每一层的面积，然后用这个面积减去height的和，剩下的就是水量了
// *************************************************

func trap(height []int) int {
    if len(height) <= 2 {
        return 0
    }

    // 找出最大高度，确定要计算多少层
    maxHeight := 0
    totalHeightSum := 0
    for _, h := range height {
        totalHeightSum += h
        if h > maxHeight {
            maxHeight = h
        }
    }

    // 如果所有柱子高度都为0，直接返回0
    if maxHeight == 0 {
        return 0
    }

    totalArea := 0
    h := 1 // 从第1层开始

    // 逐层计算面积
    for h <= maxHeight {
        left, right := 0, len(height)-1

        // 找到左边第一个高度 >= h 的柱子
        for left <= right && height[left] < h {
            left++
        }

        // 找到右边第一个高度 >= h 的柱子
        for left <= right && height[right] < h {
            right--
        }

        // 如果左右指针有效，计算当前层的面积
        if left <= right {
            // 当前层的面积 = 右指针 - 左指针 + 1
            totalArea += (right - left + 1)
        }

        h++
    }

    // 总水量 = 分层总面积 - 柱子占据的总面积
    return totalArea - totalHeightSum
}