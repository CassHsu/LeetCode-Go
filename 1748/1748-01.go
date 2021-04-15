func sumOfUnique(nums []int) int {
    m := map[int]int{}
    for _, n := range nums {
        count, ok := m[n]
        if ok {
            m[n] = count + 1
        } else {
            m[n] = 1
        }
    }
    
    ans := 0
    for k, v := range m {
        if v == 1 {
            ans += k
        }
    }
    return ans
}
