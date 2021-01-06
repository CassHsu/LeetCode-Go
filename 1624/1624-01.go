func maxLengthBetweenEqualCharacters(s string) int {
    m := map[rune]int{}
    ans := -1
    
    for i, c := range s {
        v, ok := m[c]
        if ok {
            d := i - v - 1
            if d > ans {
                ans = d
            }
            
        } else {
            m[c] = i
        }
    }
    return ans
}