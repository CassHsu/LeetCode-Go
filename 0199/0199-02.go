func rightSideView(root *TreeNode) []int {
    ans := []int{}
    if root == nil { return ans }
    
    q := []*TreeNode { root }
    for len(q) > 0 {
        count := len(q)
        
        for i := 0; i < count; i++ {
            n := q[0]
            q = q[1:]
            
            if n.Left != nil {
                q = append(q, n.Left)
            }
            
            if n.Right != nil {
                q = append(q, n.Right)
            }
            
            if i == count - 1 {
                ans = append(ans, n.Val)
            }
        }
    }
    
    return ans
}
