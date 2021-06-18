func averageOfLevels(root *TreeNode) []float64 {
    sums := []int{}
    counts := []int{}
    ret := []float64{}
    
    dfs(root, 0, &sums, &counts)
    for i := 0; i < len(sums); i++ {
        ret = append(ret, float64(sums[i]) / float64(counts[i]))
    }
    return ret;
}

func dfs(n *TreeNode, i int, sums *[]int, counts *[]int) {
    if n == nil {
        return
    }
    
    if i < len(*sums) {
        (*sums)[i] += n.Val
        (*counts)[i]++
        
    } else {
        *sums = append(*sums, n.Val)
        *counts = append(*counts, 1)
    }
    
    dfs(n.Left, i+1, sums, counts)
    dfs(n.Right, i+1, sums, counts)
}
