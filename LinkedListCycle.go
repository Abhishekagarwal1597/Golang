/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func detectCycle(head *ListNode) *ListNode {
    // Check if the input linked list is nil or has only one node
    if head == nil || head.Next == nil {
        return nil
    }
    
    // Initialize two pointers to the head of the linked list
    slow, fast := head, head
    
    // Use "tortoise and hare" algorithm to find cycle in linked list
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        
        // If slow and fast pointers meet at some node, there is a cycle in the linked list
        if slow == fast {
            // Initialize another pointer to the head of the linked list
            cycleStart := head
            
            // Move both cycleStart and slow pointers one node at a time until they meet at start of cycle
            for cycleStart != slow {
                cycleStart = cycleStart.Next
                slow = slow.Next
            }
            
            // Return the node where the cycle starts
            return cycleStart
        }
    }
    
    // If there is no cycle in the linked list, return nil
    return nil
}

