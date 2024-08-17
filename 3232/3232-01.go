import "strconv"

func canAliceWin(nums []int) bool {
    s := 0
    d := 0

    for _, n := range nums {
        if len(strconv.Itoa(n)) == 1 {
            s += n
        } else {
            d += n
        }
    }

    return s != d
}
