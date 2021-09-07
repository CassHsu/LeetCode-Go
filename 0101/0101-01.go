func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return false
    }
    
    return same(root.Left, invert(root.Right))
}

func invert(root *TreeNode) *TreeNode {
    if root == nil { return nil }
    
    root.Left, root.Right = invert(root.Right), invert(root.Left)
    
    return root
}

func same(m *TreeNode, n *TreeNode) bool {
    if m == nil && n == nil { return true }
    if m == nil || n == nil { return false }
    if m.Val != n.Val { return false }
    
    return same(m.Left, n.Left) && same(m.Right, n.Right)
}
