type MinStack struct {
    stack []int
    minlist []int
    count int
}

/** initialize your data structure here. */
func Constructor() MinStack {
    var minStack = new (MinStack)
    minStack.stack = []int{}
    minStack.minlist = []int{}
    minStack.count = 0
    return *minStack
}

func (this *MinStack) Push(x int)  {
    this.stack = append(this.stack, x)
    
    if this.count > 0 && x > this.minlist[this.count - 1] {
        this.minlist = append(this.minlist, this.minlist[this.count - 1])
    } else {
        this.minlist = append(this.minlist, x)
    }
    this.count++
}

func (this *MinStack) Pop()  {
    this.stack = this.stack[:this.count-1]
    this.minlist = this.minlist[:this.count-1]
    this.count--
}

func (this *MinStack) Top() int {
    return this.stack[this.count - 1]
}

func (this *MinStack) GetMin() int {
    return this.minlist[this.count - 1]
}
