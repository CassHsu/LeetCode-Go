func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
    sa := 0
    sb := 0
    ret := []int{0, 0}
    ma := make(map[int]struct{}, len(aliceSizes))
    
    for _, a := range aliceSizes {
        sa += a
        ma[a] = struct{}{}
    }
    
    for _, b := range bobSizes {
        sb += b
    }
    
    t := (sb - sa) >> 1
    for _, b := range bobSizes {
        if _, ok := ma[b - t]; ok {
            ret[0] = b - t
            ret[1] = b
            return ret
        }
    }
    
    return ret
}
