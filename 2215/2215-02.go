func findDifference(nums1 []int, nums2 []int) [][]int {
    m1 := map[int]bool{}
    m2 := map[int]bool{}
    
    for _, n := range nums1 {
        m1[n] = true
    }
    
    for _, n := range nums2 {
        m2[n] = true
    }
    
    ans := make([][]int, 2)
    for n, _ := range m1 {
        if _, ok := m2[n]; !ok {
            ans[0] = append(ans[0], n)
        }
    }
    
    for n, _ := range m2 {
        if _, ok := m1[n]; !ok {
            ans[1] = append(ans[1], n)
        }
    }
    
    return ans
}
