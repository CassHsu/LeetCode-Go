func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
    sa := 0
    sb := 0
    ret := []int{0, 0}
    
    for _, a := range aliceSizes {
        sa += a
    }
    
    for _, b := range bobSizes {
        sb += b
    }
    
    t := (sa - sb) / 2
    for _, a := range aliceSizes {
        for _, b := range bobSizes {
            if (a == b + t) {
                ret[0] = a
                ret[1] = b
                return ret
            }
        
        }
    }
    
    return ret
}
