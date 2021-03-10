func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    start := list1
    for i := 1; i < a; i++ {
        start = start.Next
    }
    
    end := start
    for i := a; i <= b; i++ {
        end = end.Next
    }
    
    start.Next = list2
    for list2.Next != nil {
        list2 = list2.Next
    }
    
    list2.Next = end.Next
    return list1
}
