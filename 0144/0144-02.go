func preorderTraversal(root *TreeNode) []int {
    ans := []int{}
    stack := []*TreeNode{}
    
    if root != nil {
        stack = append(stack, root)
    }
    
    for len(stack) > 0 {
        node := stack[len(stack) - 1]
        stack = stack[:len(stack)- 1]
        
        ans = append(ans, node.Val)
        
        if node.Right != nil {
            stack = append(stack, node.Right)
        }
        
        if node.Left != nil {
            stack = append(stack, node.Left)
        }
    }
    
    return ans
}
