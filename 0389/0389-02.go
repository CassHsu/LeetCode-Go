func findTheDifference(s string, t string) byte {
    ms := map[rune]int{}
    mt := map[rune]int{}
    
    for _, c := range s {
        v, ok := ms[c]
        if !ok {
            ms[c] = 1
        } else {
            ms[c] = v + 1
        }
    }
    
    for _, c := range t {
        v, ok := mt[c]
        if !ok {
            mt[c] = 1
        } else {
            mt[c] = v + 1
        }
    }
    
    for k, vt := range mt {
        vs, ok := ms[k]
        if !ok || vs != vt {
            return byte(k)
        }
    }
    return byte(0)
}
