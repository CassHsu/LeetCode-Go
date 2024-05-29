func longestMonotonicSubarray(nums []int) int {
    mx := 1
    up := 1
    dn := 1

    for i := 1; i < len(nums); i++ {
        if nums[i - 1] > nums[i] {
            up = 1
            dn += 1
        } else if nums[i - 1] < nums[i] {
            up += 1
            dn = 1
        } else {
            up = 1
            dn = 1
        }

        mx = max(mx, up, dn)
    }

    return mx
}
