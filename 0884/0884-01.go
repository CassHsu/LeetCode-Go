import (
    "fmt"
    "strings"
)

func uncommonFromSentences(s1 string, s2 string) []string {
    m := map[string]int{}
    
    as1 := strings.Split(s1, " ")
    as2 := strings.Split(s2, " ")
    
    for _, w := range as1 {
        v, ok := m[w]
        if !ok {
            m[w] = 1
        } else {
            m[w] = v + 1
        }
    }
    
    for _, w := range as2 {
        v, ok := m[w]
        if !ok {
            m[w] = 1
        } else {
            m[w] = v + 1
        }
    }
    
    ret := []string{}
    for k, v := range m {
        if v == 1 {
            ret = append(ret, k)
        }
    }
    
    return ret
}
