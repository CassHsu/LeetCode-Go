func maximizeSum(nums []int, k int) int {
    ans := 0
    mx := 0
    for _, n := range nums {
        if n > mx {
            mx = n
        }
    }

    for i := 0; i < k; i++ {
        ans += mx
        ans += i
    }

    return ans
}
