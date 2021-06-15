import (
    "fmt"
    "strings"
)

func isalpha(c byte) bool {
    if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')  {
        return true
    }
    return false
}

func reverseOnlyLetters(s string) string {
    ret := []string{}
    j := len(s) - 1
    
    for i := range s {
        if isalpha(s[i]) {
            for isalpha(s[j]) == false {
                j--
            }
            
            ret = append(ret, string(s[j]))
            j--
            
        } else {
            ret = append(ret, string(s[i]))
        }
    }
    return strings.Join(ret, "")
}
