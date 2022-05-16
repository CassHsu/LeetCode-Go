var sum int
var maxlv int

func deepestLeavesSum(root *TreeNode) int {
    sum = 0
    maxlv = 0
    
    dfs(root, 0)
    return sum
}

func dfs(root *TreeNode, lv int) {
    if root == nil {
        return
    }
    
    if lv > maxlv {
        maxlv, sum = lv, root.Val
    } else if lv == maxlv {
        sum += root.Val
    }
    
    dfs(root.Left, lv + 1)
    dfs(root.Right, lv + 1)
}
