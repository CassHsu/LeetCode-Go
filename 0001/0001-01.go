func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    
    for i, n := range nums {
        find := target - n
        
        if v, ok := m[find]; ok {
            return []int {v, i}
        }
        
        m[n] = i
    }
    
    return nil
}
