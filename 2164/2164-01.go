func sortEvenOdd(nums []int) []int {
    odds := []int{}
    evens := []int{}
    
    for i, n := range nums {
        if i % 2 == 0 {
            evens = append(evens, n)
        } else {
            odds = append(odds, n)
        }
    }
    
    sort.Ints(evens)
    
    sort.Slice(odds, func(i, j int) bool {
        return odds[i] > odds[j]
    })
    
    i := 0
    j := 0
    oddSize := len(odds)
    evenSize := len(evens)
    
    ans := []int{}
    for i < evenSize || j < oddSize {
        if i < evenSize {
            ans = append(ans, evens[i])
            i++
        }
        
        if j < oddSize {
            ans = append(ans, odds[j])
            j++
        }
    }
    
    return ans
}
