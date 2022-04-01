func findDifference(nums1 []int, nums2 []int) [][]int {
    m1 := map[int]bool{}
    m2 := map[int]bool{}
    
    for _, n := range nums1 {
        m1[n] = true
    }
    
    for _, n := range nums2 {
        m2[n] = true
    }
    
    a1 := []int{}
    for n, _ := range m1 {
        if _, ok := m2[n]; !ok {
            a1 = append(a1, n)
        }
    }
    
    a2 := []int{}
    for n, _ := range m2 {
        if _, ok := m1[n]; !ok {
            a2 = append(a2, n)
        }
    }
    
    ans := [][]int{a1, a2}
    return ans
}
