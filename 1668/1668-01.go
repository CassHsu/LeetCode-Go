import (
    "strings"
)

func maxRepeating(sequence string, word string) int {
    repeat := word
    count := 0
    
    for strings.Contains(sequence, repeat) {
        count++
        repeat += word
    }
    
    return count
}