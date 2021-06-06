import (
    "fmt"
    "strings"
)

func wordPattern(pattern string, s string) bool {
    words := strings.Split(s, " ")
    if len(pattern) != len(words) {
        return false
    }
    
    mp := map[byte]string{}
    ms := map[string]byte{}
    
    for i := 0; i < len(pattern); i++ {
        p := pattern[i]
        w := words[i]
        
        if vw, ok := mp[p]; !ok {
            mp[p] = w
        } else {
            if (vw != w) { return false }
        }
        
        if vp, ok := ms[w]; !ok {
            ms[w] = p
            if (vp != p) { return false }
        }
    }
    return true
}
