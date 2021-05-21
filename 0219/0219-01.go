func containsNearbyDuplicate(nums []int, k int) bool {
    m := map[int]int{}
    for i, n := range nums {
        v, ok := m[n]
        if ok && i <= v {
            return true
        }
        m[n] = i + k
    }
    return false
}
