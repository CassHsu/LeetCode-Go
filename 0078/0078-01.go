func subsets(nums []int) [][]int {
    ret := [][]int{}
    ret = append(ret, []int{})
    
    for _, n := range nums {
        size := len(ret)
        for i := 0; i < size; i++ {
            ss := []int{}
            ss = append(ss, ret[i]...)
            ss = append(ss, n)
            ret = append(ret, ss)
        }
    }
    return ret
}
