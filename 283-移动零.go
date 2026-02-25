package cdyleetcode

// *************************************************
// 磕磕绊绊，非0数字左移，剩下的填充0
// *************************************************
func moveZeroes(nums []int)  {
    nonZeroIndex := -1
    for i, v := range(nums) {
        if (v != 0) {
            if nonZeroIndex == -1 {
                continue
            } else {
                nums[nonZeroIndex] = v
                nonZeroIndex++
            }
        } else {
            if nonZeroIndex == -1 {
                nonZeroIndex = i
            }
        }
    }

    if nonZeroIndex == -1 {
        return
    }

    for i:=nonZeroIndex; i<len(nums); i++ {
        nums[i]=0
    }
    // firstSlot := -1
    // lastSlot := -1
    // for index, value := range(nums) {
    //     if (value == 0){
    //         slot = index
    //     } else {
    //         if (slot == -1){
    //             continue
    //         } else {
    //             nums[slot] = value
    //             slot = index
    //         }
    //     }
    // }
}