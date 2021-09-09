type InOrder struct {
    ans []int
}

func inorderTraversal(root *TreeNode) []int {
    inorder := InOrder { ans: []int{}}
    inorder.traversal(root)
    return inorder.ans
}

func (in *InOrder) traversal(node *TreeNode) {
    if node == nil {
        return
    }
    
    in.traversal(node.Left)
    in.ans = append(in.ans, node.Val)
    in.traversal(node.Right)
    return
}
