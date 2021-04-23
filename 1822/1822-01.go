func arraySign(nums []int) int {
    var p float64 = 1
    for _, n := range nums {
        p *= float64(n)
    }
    return signFunc(p)
}

func signFunc(p float64) int {
    if p > 0 {
        return 1
    } else if (p < 0) {
        return -1
    } else {
        return 0
    }
}
