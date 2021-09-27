type ps struct {
    Node *TreeNode
    Val int
}

func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil { return false }
    
    stack := []ps{ ps{ Node: root, Val: root.Val }}
    
    for len(stack) >= 1 {
        curr := stack[len(stack) - 1].Node
        val := stack[len(stack) - 1].Val
        stack = stack[:len(stack) - 1]
        
        if curr.Left == nil && curr.Right == nil && val == targetSum {
            return true
        }
        
        if curr.Right != nil {
            stack = append(stack, ps{ Node: curr.Right, Val: val + curr.Right.Val } )
        }
        
        if curr.Left != nil {
            stack = append(stack, ps{ Node: curr.Left, Val: val + curr.Left.Val } )
        }
    }
    
    return false
}
