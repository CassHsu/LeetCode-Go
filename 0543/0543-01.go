type diameter struct {
    count int
}

func diameterOfBinaryTree(root *TreeNode) int {
    d := diameter { count: 0 }
    d.dfs(root)
    return d.count
}

func (d *diameter) dfs(node *TreeNode) int {
    if node == nil {
        return 0
    }
    
    l := d.dfs(node.Left)
    r := d.dfs(node.Right)
    d.count = max(d.count, l + r)
    return max(l, r) + 1
}

func max(m int, n int) int {
    if m >= n {
        return m
    } else {
        return n
    }
}
