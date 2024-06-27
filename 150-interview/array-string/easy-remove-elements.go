package main
import "fmt"

func main() {
    nums := []int{0 , 1, 2, 2, 3, 0, 4, 2}
    res := remove(nums, 2)
    fmt.Println(res, nums)
}

func remove(nums []int, val int) int {
    k, x := 0, len(nums)-1
    for i := 0; i <= x; i++ {
        if nums[i] == val {
            shift(nums, i)
            x--
            k++
        }
    }
    return (len(nums)-1 - k)
}

func shift(s []int, i int) {
    if i < 0 {
        return
    }
    if i >= len(s)-1 {
        return
    }
    temp := s[i]
    s = append(s[:i], s[i+1:]...)
    s = append(s, temp)
}
