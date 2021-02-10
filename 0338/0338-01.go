func countBits(num int) []int {
    ret := []int{}
    ret = append(ret, 0)
    
    for i := 1; i <= num; i++ {
        count := 0
        for offset := 30; offset >= 0; offset-- {
            mask := 1 << offset
            if (i & mask) > 0 {
                count++
            }
        }
        ret = append(ret, count)
    }
    return ret
}