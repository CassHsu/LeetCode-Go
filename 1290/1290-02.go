func getDecimalValue(head *ListNode) int {
    num := head.Val
    for head.Next != nil {
        num = num << 1 | head.Next.Val
        head = head.Next
    }
    return num
}
