func findKthPositive(arr []int, k int) int {
    ans := 0
    for i, c := 1, 0; c < k; i++ {
        if isContains(arr, i) == false {
            ans = i
            c++
        }
    }

    return ans
}

func isContains(arr []int, c int) bool {
    for _, a := range arr {
        if a == c {
            return true
        }
    }

    return false
}
