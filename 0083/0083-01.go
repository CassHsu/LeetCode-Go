func deleteDuplicates(head *ListNode) *ListNode {
    ite := head
    for ite != nil {
        tmp := ite.Next
        for tmp != nil && tmp.Val == ite.Val {
            tmp = tmp.Next
        }
        ite.Next = tmp
        ite = ite.Next
    }
    return head
}
