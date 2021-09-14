func flatten(root *TreeNode)  {
    if root == nil { return }
    
    l, r := root.Left, root.Right
    
    flatten(root.Left)
    flatten(root.Right)
    
    root.Left, root.Right = nil, l
    
    for root.Right != nil {
        root = root.Right
    }
    
    root.Right = r
}
