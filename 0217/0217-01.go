func containsDuplicate(nums []int) bool {
    m := map[int]int{}
    for _, n := range nums {
        _, ok := m[n]
        if ok {
            return true
        }
        m[n] = 1
    }
    return false
}