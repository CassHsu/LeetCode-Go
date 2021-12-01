func anagramMappings(nums1 []int, nums2 []int) []int {
    m := map[int]int{}
    ans := make([]int, len(nums1))
    
    for i, n := range nums2 {
        m[n] = i
    }
    
    for i, n := range nums1 {
        ans[i], _ = m[n]
    }
    
    return ans
}
