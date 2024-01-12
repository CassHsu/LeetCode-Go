import "sort"

func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    
    count := 0
    si := len(s) - 1
    gi := len(g) - 1

    for si >= 0 && gi >= 0 {
        if s[si] >= g[gi] {
            count++
            si--
            gi--
        } else {
            gi--
        }
    }
    return count
}
