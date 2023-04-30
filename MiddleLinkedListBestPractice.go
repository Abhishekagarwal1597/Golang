// This approach requires two pointers: a "slow" pointer and a "fast" pointer. The "slow" pointer moves one step at a time,
// while the "fast" pointer moves two steps at a time. When the "fast" pointer reaches the end of the linked list, 
// the "slow" pointer will be at the middle of the linked list.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    slow, fast := head, head
    
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    return slow
}

