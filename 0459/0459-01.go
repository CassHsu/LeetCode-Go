import (
    "fmt"
    "strings"
)

func repeatedSubstringPattern(s string) bool {
    size := len(s)
    if size == 0 {
        return false
    }
    ss := (s+s)[1: size*2-1]
    return strings.Contains(ss, s)
}
