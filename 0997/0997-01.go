func findJudge(n int, trust [][]int) int {
    trusts := make([]bool, n + 1)
    trusted := make([]int, n + 1)

    for i := range trust {
        trusts[trust[i][0]] = true
        trusted[trust[i][1]]++
    }

    judge := -1
    trusts[0] = true
    for i, t := range trusts {
        if t == false {
            judge = i
            break
        }
    }

    if judge == -1 {
        return -1
    }

    if trusted[judge] == n - 1 {
        return judge
    } else {
        return -1
    }
}
