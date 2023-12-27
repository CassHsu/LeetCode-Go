func buyChoco(prices []int, money int) int {
    ans := money
    step := 0
    sort.Ints(prices)

    for _, p := range prices {
        if p <= ans {
            ans -= p
            step += 1

            if step == 2 {
                return ans
            }

        } else {
            return money
        }
    }

    return money
}
