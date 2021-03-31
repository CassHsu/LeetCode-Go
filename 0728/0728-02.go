import (
    "strconv"
    "fmt"
)

func selfDividingNumbers(left int, right int) []int {
    ret := []int{}
    for i := left; i <= right; i++ {
        if check(i) {
            ret = append(ret, i)
        }
    }
    return ret
}

func check(num int) bool {
    numStr := strconv.Itoa(num)
    for _, n := range numStr {
        d, _ := strconv.Atoi(string(n)) 
        if d == 0 || num % d != 0 {
            return false
        }
    }
    return true
}
