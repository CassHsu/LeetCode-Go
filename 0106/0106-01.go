func buildTree(inorder []int, postorder []int) *TreeNode {
    if inorder == nil || postorder == nil  {
        return nil
    }
    
    postlen := len(postorder)
    if len(inorder) == 0 || postlen == 0  {
        return nil
    }
    
    root := TreeNode { Val: postorder[postlen - 1] }
    
    mid := 0
    for i, v := range inorder {
        if v == root.Val {
            mid = i
            break
        }
    }
    
    root.Left  = buildTree(inorder[:mid],   postorder[:mid])
    root.Right = buildTree(inorder[mid+1:], postorder[mid: postlen - 1])
    
    return &root
}
