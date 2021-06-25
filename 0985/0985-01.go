func sumEvenAfterQueries(nums []int, queries [][]int) []int {
    ans := []int{}
    s := 0
    for _, n := range nums {
        if n % 2 == 0 {
            s += n
        }
    }
    
    for _, q := range queries {
        if nums[q[1]] % 2 == 0 {
            s -= nums[q[1]]
        }
        
        nums[q[1]] += q[0]
        
        if nums[q[1]] % 2 == 0 {
            s += nums[q[1]]
        }
        
        ans = append(ans, s)
    }
    
    return ans
}
