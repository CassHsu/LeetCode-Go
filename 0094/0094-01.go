type InOrder struct {
    ans []int
}

func inorderTraversal(root *TreeNode) []int {
    inorder := InOrder { ans: []int{}}
    inorder.dfs(root)
    return inorder.ans
}

func (in *InOrder) dfs(node *TreeNode) {
    if node == nil {
        return
    }
    
    in.dfs(node.Left)
    in.ans = append(in.ans, node.Val)
    in.dfs(node.Right)
}
