import (
    "strconv"
    "strings"
)

func dfs(node *TreeNode) string {
    if node.Left == nil && node.Right == nil {
        return strconv.Itoa(node.Val)
    }
    
    ret := []string{}
    if node.Left != nil {
        ret = append(ret, dfs(node.Left))
    }
    
    if node.Right != nil {
        ret = append(ret, dfs(node.Right))
    }
    
    return strings.Join(ret, "-")
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    return dfs(root1) == dfs(root2)
}
