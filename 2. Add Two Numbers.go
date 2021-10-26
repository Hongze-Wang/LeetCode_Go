/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    rear := head
    carry := 0
    for l1 != nil || l2 != nil {
        sum := carry
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        carry = sum / 10
        val := sum % 10
        rear.Next = &ListNode{val, nil}
        rear = rear.Next
    }
    if carry != 0 {
        rear.Next = &ListNode{carry, nil}
    }
    return head.Next
}
