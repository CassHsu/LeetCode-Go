import (
    "strings"
)

func mostWordsFound(sentences []string) int {
    max := 0
    
    for _, s := range sentences {
        count := len(strings.Split(s, " "))
        if count > max {
            max = count
        }
    }
    
    return max
}
