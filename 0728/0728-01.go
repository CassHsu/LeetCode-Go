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
    x := num
    for x > 0 {
        d := x % 10
        x /= 10
        if d == 0 || (num % d) > 0 {
            return false
        }
    }
    return true
}
