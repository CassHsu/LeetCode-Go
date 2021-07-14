func postorderTraversal(root *TreeNode) []int {
    ans := []int{}
    stack := []*TreeNode{}
    
    if root != nil {
        stack = append(stack, root)
    }
    
    for len(stack) > 0 {
        node := stack[len(stack) - 1]
        stack = stack[:len(stack)- 1]
        
        if node.Left != nil {
            stack = append(stack, node.Left)
        }
        
        if node.Right != nil {
            stack = append(stack, node.Right)
        }
        
        ans = append(ans, node.Val)
    }
    
    for i, j := 0, len(ans) - 1; i < len(ans) / 2; i, j = i+1, j-1 {
        ans[i], ans[j] = ans[j], ans[i]
    }
    return ans
}
