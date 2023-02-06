import (
    "strings"
    "strconv"
)

func countDigits(num int) int {
    numList := strings.Split(strconv.Itoa(num), "")
    ans := 0

    for _, n := range numList {
        ni, _ :=  strconv.Atoi(n)
        if num % ni == 0 {
            ans += 1
        }
    }

    return ans
}
