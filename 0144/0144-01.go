var ret []int
func preorderTraversal(root *TreeNode) []int {
    ret = make([]int, 0)
    preorder(root)
    return ret
}

func preorder(root *TreeNode) {
    if root == nil {
        return
    }
    
    ret = append(ret, root.Val)
    preorder(root.Left)
    preorder(root.Right)
}
