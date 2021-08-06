func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
    return getNum(firstWord) + getNum(secondWord) == getNum(targetWord)
}

func getNum(word string) (ret rune) {
    for _, c := range word {
        ret *= 10
        ret += c - rune('a')
    }
    
    return ret
}
