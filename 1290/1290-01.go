func getDecimalValue(head *ListNode) int {
    num := head.Val
    for head.Next != nil {
        num = num * 2 + head.Next.Val
        head = head.Next
    }
    return num
}
