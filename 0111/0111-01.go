type nodeDepth struct {
    node *TreeNode
    depth int
}

func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    q := []nodeDepth{}
    n := nodeDepth { node: root, depth: 1 }
    q = append(q, n)
    
    for len(q) != 0 {
        n, d := q[0].node, q[0].depth
        q = q[1:]
        
        if n != nil {
            if n.Left == nil && n.Right == nil {
                return d
            } else {
                if n.Left != nil {
                    q = append(q, nodeDepth { node: n.Left, depth: d + 1 })
                }
                
                if n.Right != nil {
                    q = append(q, nodeDepth { node: n.Right, depth: d + 1 })
                }
            }
        }
    }
    
    return -1
}
