package cdyleetcode

func findMin(nums []int) int {
    l := 0
    r := len(nums) - 1

    for l < r {
        mid := l + (r-l)/2

        if nums[mid] > nums[r] {
            l = mid + 1
        } else {
            r = mid
        }
    }

    return nums[l]
}

// // 下面是不会写的
// func findMin(nums []int) int {
//     l := 0
//     r := len(nums) - 1
//     var mid int
//     for {
//         if l>=r {
//             break
//         }
//         mid = (r+l)/2

//         if nums[mid] > nums[r] {
//             l = mid + 1
//         } else{
//             r = mid
//         }
//     }

//     return mid
// }

// 找旋转点，就是最小值
// 二分法，设置左右指针和中间指针
// 如果中间点大于右指针值，说明答案在右边，那左指针右移
// 如果中间点小于等于右指针值，说明中间点在左边，右指针左移