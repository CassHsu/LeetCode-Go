import (
    "strings"
    "fmt"
)

func canBeTypedWords(text string, brokenLetters string) int {
    texts := strings.Split(text, " ")
    ans := len(texts)
    
    for _, t := range texts {
        for i := 0; i < len(brokenLetters); i++ {
            if checkIn(t, brokenLetters[i]) {
                ans--
                break
            }
        }
    }
    
    return ans
}

func checkIn(text string, b byte) bool {
    for i := 0; i < len(text); i++ {
        if text[i] == b {
            return true
        }
    }
    
    return false
}
