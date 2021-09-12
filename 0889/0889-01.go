func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
    if preorder == nil || postorder == nil { return nil }
    postlen := len(postorder)
    if len(preorder) == 0 || postlen == 0 { return nil }
    
    root := TreeNode { Val: preorder[0] }
    if len(preorder) == 1 { return &root }
    
    left_root_idx := 0
    for i, v := range postorder {
        if v == preorder[1] {
            left_root_idx = i
            break
        }
    }
    
    root.Left  = constructFromPrePost(preorder[1: left_root_idx + 2], postorder[: left_root_idx+1])
    root.Right = constructFromPrePost(preorder[left_root_idx + 2:],   postorder[left_root_idx + 1: postlen - 1])
    return &root
}
