func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    count := 0
    m := map[int]int {}
    
    for _, n1 := range nums1 {
        for _, n2 := range nums2 {
            if n, ok := m[n1 + n2]; ok {
                m[n1 + n2] = n + 1
            } else {
                m[n1 + n2] = 1
            }
        }
    }
    
    for _, n3 := range nums3 {
        for _, n4 := range nums4 {
            if n, ok := m[-(n3 + n4)]; ok {
                count += n
            }
        }
    }
    
    return count
}
