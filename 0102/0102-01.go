type LevelOrder struct {
    ans [][]int
}

func levelOrder(root *TreeNode) [][]int {
    lo := LevelOrder { ans: [][]int{} }
    lo.traversal(root, 0)
    return lo.ans
}

func (o *LevelOrder) traversal(node *TreeNode, lvl int) {
    if node == nil {
        return
    }
    
    if len(o.ans) <= lvl {
        o.ans = append(o.ans, []int { node.Val })
    } else {
        o.ans[lvl] = append(o.ans[lvl], node.Val)
    }
    
    o.traversal(node.Left, lvl + 1)
    o.traversal(node.Right, lvl + 1)
    
    return
}
