import (
    "fmt"
    "strconv"
)

func reformat(s string) string {
    alpha := []string{}
    digit := []string{}
    
    for _, c := range s {
        sc := string(c)
        if _, err := strconv.Atoi(sc); err == nil {
            digit = append(digit, sc)
        } else {
            alpha = append(alpha, sc)
        }
    }
    
    if len(alpha) - len(digit) > 1 || len(digit) - len(alpha) > 1 {
        return ""
    }
    
    isAlphaFirst := false
    if len(alpha) > len(digit) {
        isAlphaFirst = true
    }
    
    ans := ""
    a := 0
    d := 0
    
    for len(alpha) > a && len(digit) > d {
        if isAlphaFirst {
            ans += (alpha[a] + digit[d])
        } else {
            ans += (digit[d] + alpha[a])
        }
        
        a++
        d++
    }
    
    if len(alpha) > a {
        ans += alpha[a]
    }
    if len(digit) > d {
        ans += digit[d]
    }
    
    return ans
}
