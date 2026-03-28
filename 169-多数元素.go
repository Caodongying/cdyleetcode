package cdyleetcode

// *************************************************
// 看的题解，摩尔排序。就是“消除”，一次遍历，设一个擂主，如果
// 当前元素和擂主一样，累计值+1，不然减一。累计值为0，更换擂主。
// 也可以这么想“帮派斗争，不同派系一命抵一命”，活到最后的那个帮派就赢了。
// *************************************************
func majorityElement(nums []int) int {
    target := nums[0]
    sum := 0

    for _, v := range nums {
        if sum == 0 {
            target = v
        }

        if v == target {
            sum++
        } else {
            sum--
        }

    }

    return target
}
