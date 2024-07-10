func minimumAverage(nums []int) float64 {
    sort.Ints(nums)
    size := len(nums)
    l := 0
    r := size - 1
    ans := (float64)(nums[r])
    avg := 0.0

    for l <= r {
        avg = (float64)(nums[l] + nums[r]) / 2.0
        if avg < ans {
            ans = avg
        }
        l += 1
        r -= 1
    }

    return ans
}
