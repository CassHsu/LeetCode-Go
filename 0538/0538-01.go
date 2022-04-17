func convertBST(root *TreeNode) *TreeNode {
    dfs(root, 0)
    return root
}

func dfs(node *TreeNode, total int) int {
    if node == nil {
        return total
    }
    
    node.Val += dfs(node.Right, total)
    return dfs(node.Left, node.Val)
}
