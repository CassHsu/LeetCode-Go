import "sort"

func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    
    count := 0
    for len(g) > 0 && len(s) > 0 {
        if s[0] >= g[0] {
            count++
            g = g[1:]
        }
        s = s[1:]
    }
    return count
}
