type Lonely struct {
    Result []int
}

func (l *Lonely) DFS(node *TreeNode) {
    if node == nil {
        return
    }
    
    if node.Left != nil && node.Right == nil {
        l.Result = append(l.Result, node.Left.Val)
    }
    
    if node.Left == nil && node.Right != nil {
        l.Result = append(l.Result, node.Right.Val)
    }
    
    l.DFS(node.Left)
    l.DFS(node.Right)
}

func getLonelyNodes(root *TreeNode) []int {
    lonely := Lonely{ Result: []int{} }
    
    lonely.DFS(root)
    return lonely.Result
}
