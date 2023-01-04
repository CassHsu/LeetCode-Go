import "strings"

func detectCapitalUse(word string) bool {
    if word == strings.ToLower(word) {
        return true
    }

    if word == strings.ToUpper(word) {
        return true
    }

    if word == strings.ToUpper(string(word[0])) + strings.ToLower(word[1: len(word)]) {
        return true
    }

    return false
}
