package main

import "fmt"
func rotate(nums []int, k int)  {
    k = k % len(nums)

    left := nums[len(nums)-k:len(nums)]

    right := nums[:len(nums)-k]

    fmt.Println("left", left)
    fmt.Println("right", right)
    nums = append(left, right...)

    fmt.Println(nums)
}
func main() {
    nums := []int{1, 2, 3, 4, 5, 6, 7}
    k := 3
    rotate(nums, k)

}
