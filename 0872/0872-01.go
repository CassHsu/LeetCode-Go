import (
    "strconv"
)

func dfs(node *TreeNode) string {
    if node.Left == nil && node.Right == nil {
        return strconv.Itoa(node.Val)
    }
    
    ret := ""
    if node.Left != nil {
        v := dfs(node.Left)
        if ret != "" {
            ret = ret + "-" + v
        } else {
            ret = v
        }
    }
    
    if node.Right != nil {
        v := dfs(node.Right)
        if ret != "" {
            ret = ret + "-" + v
        } else {
            ret = v
        }
    }
    
    return ret
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    return dfs(root1) == dfs(root2)
}
