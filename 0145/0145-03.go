type PostOrder struct {
    ans []int
}

func postorderTraversal(root *TreeNode) []int {
    postorder := PostOrder { ans: []int{} }
    postorder.traversal(root)
    return postorder.ans
}

func (post *PostOrder) traversal(node *TreeNode) {
    if node == nil {
        return
    }
    
    post.traversal(node.Left)
    post.traversal(node.Right)
    post.ans = append(post.ans, node.Val)
    return
}
