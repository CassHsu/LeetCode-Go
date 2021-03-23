func isPalindrome(head *ListNode) bool {
    fast := head
    slow := head
    stack := []int{}
    
    for fast != nil && fast.Next != nil {
        stack = append(stack, slow.Val)
        
        fast = fast.Next.Next
        slow = slow.Next
    }
    
    if fast != nil {
        slow = slow.Next
    }
    
    for slow != nil {
        size := len(stack) - 1
        if slow.Val != stack[size] {
            return false
        }
        stack = stack[:size]
        slow = slow.Next
    }
    return true
}
