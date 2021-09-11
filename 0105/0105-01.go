func buildTree(preorder []int, inorder []int) *TreeNode {
    if preorder == nil || len(preorder) <= 0 || inorder == nil || len(inorder) <= 0 {
        return nil
    }
    
    val := preorder[0]
    mid := 0
    for i, v := range inorder {
        if v == val {
            mid = i
            break
        }
    }
    
    root := TreeNode { Val: val }
    root.Left  = buildTree(preorder[1:mid+1], inorder[:mid])
    root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
    return &root
}
