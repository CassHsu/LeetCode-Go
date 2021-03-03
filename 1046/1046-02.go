import (
    "container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int {return len(h)}
func (h IntHeap) Less(i, j int) bool {return h[i] > h[j]}
func (h IntHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *IntHeap) Push(node interface{}) {*h = append(*h, node.(int))}
func (h *IntHeap) Pop() interface{} {
    size := len(*h)
    ret := (*h)[size-1]
    *h = (*h)[0:size-1]
    return ret
}

func lastStoneWeight(stones []int) int {
    pq := IntHeap(stones)
    heap.Init(&pq)
    for pq.Len() > 1 {
        m1 := heap.Pop(&pq).(int)
        m2 := heap.Pop(&pq).(int)
        if m1 != m2 {
            heap.Push(&pq, m1 - m2)
        }
    }
    
    if pq.Len() == 0 {
        return 0
    }
    
    return heap.Pop(&pq).(int)
}
