func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    m := map[string]int{ "type": 0, "color": 1, "name": 2 }
    count := 0
    for _, item := range items {
        if item[m[ruleKey]] == ruleValue {
            count++
        }
    }
    return count
}
