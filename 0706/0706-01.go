type MyHashMap struct {
    buckets []int
}

func Constructor() MyHashMap {
    buckets := make([]int, 1000001)
    
    for i, _ := range buckets {
        buckets[i] = -1
    }
    
    return MyHashMap { buckets: buckets }
}

func (this *MyHashMap) Put(key int, value int)  {
    this.buckets[key] = value
}

func (this *MyHashMap) Get(key int) int {
    return this.buckets[key]
}

func (this *MyHashMap) Remove(key int)  {
    this.buckets[key] = -1
}
