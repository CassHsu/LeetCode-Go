func swapNodes(head *ListNode, k int) *ListNode {
    curr := head
    var n1 *ListNode = nil
    var n2 *ListNode = nil
    
    for curr != nil {
        k--
        if n2 != nil {
            n2 = n2.Next
        }
        if k == 0 {
            n1 = curr
            n2 = head
        }
        curr = curr.Next
    }
    n1.Val, n2.Val = n2.Val, n1.Val
    
    return head
}
