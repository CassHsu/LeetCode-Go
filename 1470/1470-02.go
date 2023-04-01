func shuffle(nums []int, n int) []int {
    size := len(nums)
    ans := []int{}

    j, k := 0, n
    for i := 0; i < size; i++ {
        v := 0
        if i % 2 == 0 {
            v = nums[j]
            j++
        } else {
            v = nums[k]
            k++
        }
        ans = append(ans, v)
    }
    return ans
}
