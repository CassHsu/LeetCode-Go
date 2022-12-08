func rangeSumBST(root *TreeNode, low int, high int) int {
    if root == nil {
        return 0
    }

    v := 0
    if low <= root.Val && root.Val <= high {
        v = root.Val
    }

    l := rangeSumBST(root.Left, low, high)
    r := rangeSumBST(root.Right, low, high)

    return l + r + v
}
