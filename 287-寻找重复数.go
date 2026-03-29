package cdyleetcode

// *************************************************
// ❌❌ 自己写的。以为如果当前异或后的结果，可以和上上个结果一样，就说明找到了。但无法从异或结果中直接定位哪个数字重复了
// *************************************************
func findDuplicate(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }

    result := nums[0]

    var target int

    for i:=1; i<len(nums); i++ {
        newResult := result ^ nums[i]
        if result == newResult {
            target = nums[i]
            break
        }
        result = newResult
    }

    return target
}

// 异或运算
// 如果异或一个新值后，结果和异或前一样，新值就是结果

// *************************************************
// ❌❌ 错误理解题意，本题并不是1到n每个数都出现，因此重复数可能出现不止两次
// *************************************************
func findDuplicate(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    result := 0

    for _, v := range nums {
        result = result ^ v
    }

    for i:=1; i<=len(nums)-1; i++ {
        result = result ^ i
    }

    return result
}


// *************************************************
// 看了题解后自己写的
// *************************************************
func findDuplicate(nums []int) int {
    slow, fast := 0, 0

    for {
        slow = nums[slow]
        fast = nums[nums[fast]]
        if slow == fast {
            break
        }
    }

    head := 0
    for {
        slow = nums[slow]
        head = nums[head]
        if slow == head {
            break
        }
    }

    return slow

}
// 下标 0 1 2 3 4
// 数值 1 3 4 2 2
// i是下标，把nums[i]看做是从i出发的next指针（nums[i]的取址是1到len(nums)-1
// 构造一个链表，链表内容表示从nums[0]出发，根据next指针，一路遍历到的数值。0作为第一个元素
// 也就是 0 -> 1 -> 3 -> 2 -> 4 -> 2
// 2 -> 4 -> 2就构成一个环
// 然后用Floyd判圈算法，用快慢指针slow和fast （正确性证明：https://www.bilibili.com/video/BV1gJbnzwEqr/?spm_id_from=333.337.search-card.all.click&vd_source=e3d59376cbbf7e58fec4949fdb3a00d4） 不用想为什么，用方程式证明的
// slow一次走一步，fast一次走两步，第一次相遇是在环里。这时候说明有环，入还口就是重复元素
// slow位置不变，fast回到原点，这下两个都一次走一步，相遇的地点就是重复数，也就是入还口

