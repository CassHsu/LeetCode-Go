type RecentCounter struct {
    requests []int
}

func Constructor() RecentCounter {
    rc := RecentCounter{}
    rc.requests = []int{}
    return rc
}

func (this *RecentCounter) Ping(t int) int {
    this.requests = append(this.requests, t)
    l := t - 3000
    for _, r := range this.requests {
        if r < l {
            this.requests = this.requests[1:]
        }
    }
    return len(this.requests)
}