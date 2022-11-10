import (
    "fmt"
    "strconv"
    "strings"
)

func maximum69Number (num int) int {
    s := strconv.Itoa(num)
    res := []string{}
    find := false

    for _, r := range s {
        c := string(r)
        if find == false && c == "6" {
            res = append(res, "9")
            find = true
        } else {
            res = append(res, c)
        }
    }

    ans, _ := strconv.Atoi(strings.Join(res, ""))
    return ans
}
