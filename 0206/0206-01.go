func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    curr := head
    var prev *ListNode = nil
    var next *ListNode = nil
    
    for curr != nil {
        next = curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}
