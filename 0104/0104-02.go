func maxDepth(root *TreeNode) int {
    count := 0
    count = dfs(root)
    return count
}

func dfs(node *TreeNode) int {
    if node == nil {
        return 0
    }
    
    l := dfs(node.Left) + 1
    r := dfs(node.Right) + 1
    
    if l >= r {
        return l
    } else {
        return r
    }
}
