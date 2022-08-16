func romanToInt(s string) int {
    m := map[string]int {
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }
    
    count := 0
    for i := 1; i < len(s); i++ {
        curr := m[s[i - 1 : i]]
        next := m[s[i : i + 1]]
        
        
        if curr < next {
            count = count - curr
        } else {
            count = count + curr
        }
        
    }
    
    count = count + m[s[len(s)-1:len(s)]]
    
    return count
}
