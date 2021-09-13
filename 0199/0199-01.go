func rightSideView(root *TreeNode) []int {
    ans := []int{}
    lo := LevelOrder { lvls: [][]int{} }
    lo.traversal(root, 0)
    
    for _, lvl := range lo.lvls {
        ans = append(ans, lvl[len(lvl) - 1])
    }
    
    return ans
}

type LevelOrder struct {
    lvls [][]int
}

func (lo *LevelOrder) traversal(node *TreeNode, lv int) {
    if node == nil { return }
    
    if len(lo.lvls) <= lv {
        lo.lvls = append(lo.lvls, []int{ node.Val })
    } else {
        lo.lvls[lv] = append(lo.lvls[lv] , node.Val)
    }
    
    lo.traversal(node.Left, lv+1)
    lo.traversal(node.Right, lv+1)
    return
}
