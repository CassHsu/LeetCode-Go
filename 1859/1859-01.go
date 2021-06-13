import (
    "strings"
    "strconv"
)

func sortSentence(s string) string {
    ss := strings.Split(s, " ")
    ret := make([]string, len(ss))
    
    for _, w := range ss {
        i, _ := strconv.Atoi(string(w[len(w) - 1]))
        ret[i - 1] = w[:len(w) - 1]
    }
    return strings.Join(ret, " ")
}
