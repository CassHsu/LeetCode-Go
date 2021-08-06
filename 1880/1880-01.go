import (
    "fmt"
    "strings"
    "strconv"
)

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
    firstNum, _ := strconv.Atoi(getNum(firstWord));
    secondNum, _ := strconv.Atoi(getNum(secondWord));
    targetNum, _ := strconv.Atoi(getNum(targetWord));
    
    return firstNum + secondNum == targetNum
}

func getNum(word string) string {
    ls := strings.Split(word, "")
    
    for i, c := range ls {
        ls[i] = strconv.Itoa(int(c[0]) - int('a'))
    }
    
    return strings.Join(ls, "")
}
