func maximumNumberOfStringPairs(words []string) int {
    count := 0
    n := len(words)

    for i := 0; i < n; i++ {
      for j := i + 1; j < n; j++ {
        if words[i][0] == words[j][1] && words[i][1] == words[j][0] {
          count += 1
        }
      }
    }

    return count
}
