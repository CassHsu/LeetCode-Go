import (
    "strings"
)

func reversePrefix(word string, ch byte) string {
    idx := 0
    for i, w := range word {
        if byte(w) == ch {
            idx = i
            break
        }
    }
    
    arr := strings.Split(word[:idx+1], "")
    l, r := 0, idx
    for l <= r {
        arr[l], arr[r] = arr[r], arr[l]
        l++
        r--
    }
    
    word1 := strings.Join(arr, "")
    
    return word1 + word[idx+1:]
}
