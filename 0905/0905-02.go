func sortArrayByParity(nums []int) []int {
    ans := make([]int, len(nums))
    c := 0
    
    for _, n := range nums {
        if n % 2 == 0 {
            ans[c] = n
            c++
        }
    }
    
    for _, n := range nums {
        if n % 2 == 1 {
            ans[c] = n
            c++
        }
    }
    
    return ans
}
