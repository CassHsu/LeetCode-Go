func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
    return getNum(firstWord) + getNum(secondWord) == getNum(targetWord)
}

func getNum(word string) (ret int) {
    for _, c := range word {
        ret = ret * 10 + int(c - 'a') 
    }
    
    return ret
}
