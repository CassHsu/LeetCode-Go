func isBalanced(root *TreeNode) bool {
    return check(root) != -1
}

func check(node *TreeNode) int {
    if node == nil {
        return 0
    }
    
    l := check(node.Left)
    r := check(node.Right)
    
    if l == -1 || r == -1 || (l - r) > 1 || (r - l) > 1 {
        return -1
    }
    
    if l >= r {
        return l + 1
    } else {
        return r + 1
    }
}
