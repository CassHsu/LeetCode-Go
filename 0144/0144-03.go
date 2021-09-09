type PreOrder struct {
    ans []int
}

func preorderTraversal(root *TreeNode) []int {
    preorder := PreOrder{ ans: []int{} }
    preorder.traversal(root)
    return preorder.ans
}

func (p *PreOrder) traversal(node *TreeNode) {
    if node == nil {
        return
    }
    
    p.ans = append(p.ans, node.Val)
    p.traversal(node.Left)
    p.traversal(node.Right)
    return
}
