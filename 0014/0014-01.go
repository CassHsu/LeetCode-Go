import (
    "strings"
)

func longestCommonPrefix(strs []string) string {
    size := len(strs)
    if size == 0 {
        return ""
    }
    if size == 1 {
        return strs[0]
    }
    
    prefix := strs[0]
    for _, s := range strs {
        if s == "" {
            return ""
        }
        
        for strings.Index(s, prefix) != 0 {
            prefix = prefix[:len(prefix) - 1]
            if prefix == "" {
                return ""
            }
        }
        
    }
    return prefix
}