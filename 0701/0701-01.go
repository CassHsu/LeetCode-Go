func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil { return &TreeNode { Val: val } }
    
    curr := root
    for true {
        if curr.Val <= val {
            if curr.Right != nil {
                curr = curr.Right
                
            } else {
                curr.Right =  &TreeNode { Val: val }
                break
            }
            
        } else {
            if curr.Left != nil {
                curr = curr.Left
                
            } else {
                curr.Left =  &TreeNode { Val: val }
                break
            }
        }
    }
    
    return root
}
