func strStr(haystack string, needle string) int {
    if needle == "" || haystack == needle {
        return 0
    }
    
    p := 0
    hSize := len(haystack)
    nSize := len(needle)
    
    for p <= (hSize - nSize) {
        if haystack[p: p + nSize] == needle {
            return p
        } else {
            p++
        }
    }
    return -1
}