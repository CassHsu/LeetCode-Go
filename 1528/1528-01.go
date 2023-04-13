import (
    "strings"
)

func restoreString(s string, indices []int) string {
    ans := strings.Split(s, "")
    
    for i, v := range indices {
        ans[v] = string(s[i])
    }
    return strings.Join(ans, "")
}
