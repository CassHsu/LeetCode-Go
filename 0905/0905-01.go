func sortArrayByParity(nums []int) []int {
    odd := []int{}
    even := []int{}
    
    for _, n := range nums {
        if n % 2 == 0 {
            even = append(even, n)
        } else {
            odd = append(odd, n)
        }
    }
    
    return append(even, odd...)
}
