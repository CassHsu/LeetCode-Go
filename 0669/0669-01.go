func trimBST(root *TreeNode, low int, high int) *TreeNode {
    return trim(root, low, high)
}

func trim(node *TreeNode, low int, high int) *TreeNode {
    if node == nil {
        return nil
    } else if node.Val > high {
        return trim(node.Left, low, high)
    } else if node.Val < low {
        return trim(node.Right, low, high)
    } else {
        node.Left = trim(node.Left, low, high)
        node.Right = trim(node.Right, low, high)
        return node
    }
}
