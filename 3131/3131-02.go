func addedInteger(nums1 []int, nums2 []int) int {
    m1, m2 := math.MaxInt, math.MaxInt

    for _, n := range nums1 {
        m1 = min(m1, n)
    }

    for _, n := range nums2 {
        m2 = min(m2, n)
    }

    return m2 - m1
}
