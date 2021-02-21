type Node struct {
    Next *Node
    Prev *Node
    Key int
    Value int
}

type LRUCache struct {
    Capacity int
    Head *Node
    Tail *Node
    Map map[int]*Node
}

func Constructor(capacity int) LRUCache {
    m := make(map[int]*Node)
    head := &Node{Key: -1, Value: -1}
    tail := &Node{Key: -1, Value: -1}
    head.Next = tail
    tail.Prev = head
    c := LRUCache { Capacity: capacity, Map: m, Head: head, Tail: tail}
    return c
}

func (this *LRUCache) Remove(n *Node) {
    delete(this.Map, n.Key)
    n.Prev.Next = n.Next
    n.Next.Prev = n.Prev
}

func (this *LRUCache) Add(n *Node) {
    this.Map[n.Key] = n
    n.Next = this.Head.Next
    this.Head.Next.Prev = n
    n.Prev = this.Head
    this.Head.Next = n
}

func (this *LRUCache) Get(key int) int {
    n, ok := this.Map[key]
    if !ok {
        return -1
    }
    
    this.Remove(n)
    this.Add(n)
    return n.Value
}

func (this *LRUCache) Put(key int, value int)  {
    n, ok := this.Map[key]
    if ok {
        this.Remove(n)
    }
    if this.Capacity == len(this.Map) {
        this.Remove(this.Tail.Prev)
    }
    node := &Node{Key: key, Value: value}
    this.Add(node)
}
