func majorityElement(nums []int) int {
    size := len(nums) / 2
    m := map[int]int{}
    
    for _, n := range nums {
        count, ok := m[n]
        if !ok {
            m[n] = 1
        } else {
            m[n] = count + 1
        }
        
        if m[n] > size {
            return n
        }
    }
    return -1
}
