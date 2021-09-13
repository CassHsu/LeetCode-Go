import (
    "strings"
)

func reversePrefix(word string, ch byte) string {
    idx := strings.Index(word, string(ch))
    
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
