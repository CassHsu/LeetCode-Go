import (
    "strings"
)

func capitalizeTitle(title string) string {
    words := strings.Split(strings.ToLower(title), " ")
    
    for i, w := range words {
        if len(w) > 2 {
            words[i] = strings.ToUpper(string(w[0])) + w[1:]
        }
    }
    
    ans := strings.Join(words, " ")
    return ans
}
