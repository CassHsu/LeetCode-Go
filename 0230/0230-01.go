func kthSmallest(root *TreeNode, k int) int {
    leftSize := getSize(root.Left)
    if k <= leftSize {
        return kthSmallest(root.Left, k)
    }
    
    if k == leftSize + 1 {
        return root.Val
    }
    
    return kthSmallest(root.Right, k - (leftSize + 1))
}

func getSize(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return getSize(root.Left) + 1 + getSize(root.Right)
}