func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    count := 0

    for _, h := range hours {
        if h >= target {
            count += 1
        }
    }

    return count
}
