import (
    "strconv"
)

func countAndSay(n int) string {
    if n == 1 {
        return "1"
    }
    
    cs := countAndSay(n - 1)
    init := cs[0]
    count := 0
    res := ""
    
    for i, _ := range cs {
        if cs[i] == init {
            count++
        } else {
            res += strconv.Itoa(count) + string(init)
            count = 1
            init = cs[i]
        }
    }
    
    res += strconv.Itoa(count) + string(init)
    return res
}