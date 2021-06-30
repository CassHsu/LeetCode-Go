import (
    "math"
)

func shortestToChar(s string, c byte) []int {
    size := len(s)
    ans := make([]int, size)
    prev := math.MinInt32 / 2
    
    for i, ch := range s {
        if byte(ch) == c {
            prev = i
        }
        ans[i] = i - prev
    }
    
    prev = math.MaxInt32 / 2
    for i := size - 1; i >= 0; i-- {
        if s[i] == c {
            prev = i
        }
        
        if prev - i < ans[i] {
            ans[i] = prev - i
        }
    }
    
    return ans
}
