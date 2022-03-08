func mostFrequent(nums []int, key int) int {
    m := map[int]int{}
    
    for i := 1; i < len(nums); i++ {
        if nums[i - 1] == key {
            m[nums[i]]++
        }
    }
    
    ans := nums[0]
    max := 0
    
    for k, v := range m {
        if v > max {
            max = v
            ans = k
        }
    }
    
    return ans
}
