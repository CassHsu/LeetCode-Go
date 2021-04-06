func findDuplicates(nums []int) []int {
    ret := []int{}
    m := map[int]int{}
    for _, n := range nums {
        v, ok := m[n]
        if !ok {
            m[n] = 1
        } else {
            if v == 1 {
                ret = append(ret, n)
            }
            m[n]++
        }
    }
    return ret
}
