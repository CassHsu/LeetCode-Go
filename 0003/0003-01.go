func lengthOfLongestSubstring(s string) int {
    m := map[rune]int{}
    count := 0
    start := -1
    
    for i, k := range s {
        if v, exist := m[k]; exist && v > start {
            start = v
            m[k] = i
        } else {
            m[k] = i
            if i - start > count {
                count = i - start
            }
        }
    }
    
    return count
}
