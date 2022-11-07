import (
    "strings"
    "fmt"
)

func reverseVowels(s string) string {
    vowels := []string{ "a", "e", "i", "o", "u", "A", "E", "I", "O", "U" }
    m := map[string]bool{}
    for _, v := range vowels {
        m[v] = true
    }

    res := []string{}
    for _, c := range s {
        res = append(res, string(c))
    }

    l := 0
    r := len(s) - 1
    for l < r {
        _, ok1 := m[res[l]]
        _, ok2 := m[res[r]]
        
        if ok1 && ok2 {
            res[l], res[r] = res[r], res[l]
            l++
            r--
        } else if ok2 {
            l++
        } else {
            r--
        }
    }
    
    return strings.Join(res, "")
}
