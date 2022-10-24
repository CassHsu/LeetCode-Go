func findErrorNums(nums []int) []int {
    m := map[int]int{}
    ans := []int{-1, -1}
    
    for _, n := range nums {
        if v, ok := m[n]; !ok {
            m[n] = 1
        } else {
            m[n] = v + 1
        }
    }
    
    for i := 1; i <= len(nums) ; i++ {
        if v, ok := m[i]; ok {
            if v == 2 {
                ans[0] = i
            }
        } else {
            ans[1] = i
        }
    }
    
    return ans
}
