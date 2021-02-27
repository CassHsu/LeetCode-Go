import (
    "sort"
    "fmt"
    "strings"
)

func checkInclusion(s1 string, s2 string) bool {
    target := sortIt(s1)
    curr := ""
    
    for _, c := range s2 {
        curr += string(c)
        if len(curr) == len(s1) {
            if sortIt(curr) == target {
                return true
            }
            curr = curr[1:]
        }
    }
    return false
}

func sortIt(s string) string {
    t := strings.Split(s, "")
    sort.Strings(t)
    return strings.Join(t, "")
}
