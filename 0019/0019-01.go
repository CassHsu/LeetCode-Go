func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{}
    dummy.Next = head
    fast := dummy
    slow := dummy
    
    for i := 0; i < n; i++ {
        fast = fast.Next
    }
    
    for fast.Next != nil {
        fast = fast.Next
        slow = slow.Next
    }
    
    slow.Next = slow.Next.Next
    return dummy.Next
}
