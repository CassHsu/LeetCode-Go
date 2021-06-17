func averageOfLevels(root *TreeNode) []float64 {
    ret := []float64{}
    q := []*TreeNode{ root }
    
    for len(q) > 0 {
        size := len(q)
        sum := 0.0
        
        for i := 0; i < size; i++ {
            curr := q[0]
            q = q[1:]
            sum += float64(curr.Val)
            
            if curr.Left != nil {
                q = append(q, curr.Left)
            }
            
            if curr.Right != nil {
                q = append(q, curr.Right)
            }
        }
        
        ret = append(ret, sum / float64(size))
    }
    return ret
}
