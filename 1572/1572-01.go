func diagonalSum(mat [][]int) int {
    ans := 0
    l := 0
    r := len(mat) - 1

    for _, row := range mat {
        if l == r {
            ans += row[r]
        } else {
            ans += row[l]
            ans += row[r]
        }

        l++
        r--
    }

    return ans
}
