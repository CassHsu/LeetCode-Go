func allPathsSourceTarget(graph [][]int) [][]int {
    p := allPath { graph: graph, ans: [][]int{} }
    p.backtrack([]int{ 0 }, 0)
    return p.ans
}

type allPath struct {
    graph [][]int
    ans [][]int
}

func (p *allPath) backtrack(path []int, curr int) {
    if curr + 1 == len(p.graph) {
        tmp := make([]int, len(path))
        copy(tmp, path)
        p.ans = append(p.ans, tmp)
        return
    }
    
    for _, n := range p.graph[curr] {
        path = append(path, n)
        p.backtrack(path, n)
        path = path[:len(path) - 1]
    }
}
