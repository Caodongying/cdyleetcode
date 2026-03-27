package cdyleetcode

// *************************************************
// 直接看的题解，用异或，Golang异或是^。一次循环。同零异一（位运算，比如4是100）
// *************************************************
func singleNumber(nums []int) int {
    result := 0
    for _, v := range nums {
        result = result ^ v
    }

    return result
}
