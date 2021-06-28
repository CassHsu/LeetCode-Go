func minCostToMoveChips(position []int) int {
    even_cnt := 0
    odd_cnt := 0
    
    for _, v := range position {
        if v % 2 == 0 {
            even_cnt += 1
        } else {
            odd_cnt += 1
        }
    }
    
    if even_cnt > odd_cnt {
        return odd_cnt
    } else {
        return even_cnt
    }
}
