var ret []int
func postorderTraversal(root *TreeNode) []int {
    ret = make([]int, 0)
    postorder(root)
    return ret
}

func postorder(root *TreeNode) {
    if root == nil {
        return
    }
    
    postorder(root.Left)
    postorder(root.Right)
    ret = append(ret, root.Val)
}
