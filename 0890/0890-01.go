import (
    "strconv"
)

func findAndReplacePattern(words []string, pattern string) []string {
    res := []string{}
    p := findPattern(pattern)
    
    for _, w := range words {
        if p == findPattern(w) {
            res = append(res, w)
        }   
    }
    return res
}

func findPattern(pattern string) string {
    code := ""
    pm := map[rune]int{}
    i := 0
    
    for _, p := range pattern {
        _, ok := pm[p]
        if ok == false {
            pm[p] = i
            i++
        }
        code += strconv.Itoa(pm[p])
    }
    return code
}
