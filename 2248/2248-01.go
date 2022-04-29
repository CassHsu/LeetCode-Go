func intersection(nums [][]int) []int {
    res := []int{}
    count := make([]int, 1001)
    
    for _, n := range nums {
        for _, v := range n {
            count[v]++
        }
    }
    
    for i := 0; i < 1001; i++ {
        if count[i] == len(nums) {
            res = append(res, i)
        }
    }
    
    return res
}
