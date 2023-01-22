func differenceOfSum(nums []int) int {
    es, ds := 0, 0
    digits := ""

    for _, n := range nums {
        digits += strconv.Itoa(n)
        es += n
    }

    for _, d := range digits {
        s, _ := strconv.Atoi(string(d))
        ds += s
    }

    ans := 0
    if es > ds {
        ans = es - ds
    } else {
        ans = ds - es
    }
    return ans
}
