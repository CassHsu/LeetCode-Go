type ZzLevel struct {
    ans [][]int
}

func zigzagLevelOrder(root *TreeNode) [][]int {
    zz := ZzLevel { ans: [][]int{} }
    zz.traversal(root, 0)
    return zz.ans
}

func (z *ZzLevel) traversal(node *TreeNode, lv int) {
    if node == nil { return }
    
    if len(z.ans) <= lv {
        z.ans = append(z.ans, []int{ node.Val })
    } else {
        if lv % 2 == 0 {
            z.ans[lv] = append(z.ans[lv], node.Val)
        } else {
            z.appendLeft(node, lv)
        }
    }
    
    z.traversal(node.Left,  lv + 1)
    z.traversal(node.Right, lv + 1)
    return
}

func (z *ZzLevel) appendLeft(node *TreeNode, lv int) {
    z.ans[lv] = append(z.ans[lv], 0)
    copy(z.ans[lv][1:], z.ans[lv])
    z.ans[lv][0] = node.Val
}
