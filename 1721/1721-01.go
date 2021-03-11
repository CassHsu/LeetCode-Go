func swapNodes(head *ListNode, k int) *ListNode {
    values := []int{}
    curr := head
    for curr != nil {
        values = append(values, curr.Val)
        curr = curr.Next
    }
    
    il := k - 1
    ir := len(values) - k
    
    curr = head
    i := 0
    for curr != nil {
        if i == il {
            curr.Val = values[ir]
        }
        if i == ir {
            curr.Val = values[il]
        }
        i += 1
        curr = curr.Next
    }
    return head
}
