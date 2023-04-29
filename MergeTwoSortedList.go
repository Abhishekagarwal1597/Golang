
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    // Initialize a dummy node as the head of the result list
    dummy := &ListNode{Val: 0, Next: nil}
    // Initialize a pointer to the current tail of the result list
    tail := dummy
    
    // Traverse both input lists until at least one of them is exhausted
    for list1 != nil && list2 != nil {
        // Compare the heads of the input lists and append the smaller node to the result list
        if list1.Val <= list2.Val {
            tail.Next = list1
            list1 = list1.Next
        } else {
            tail.Next = list2
            list2 = list2.Next
        }
        // Move the tail pointer to the newly appended node
        tail = tail.Next
    }
    // Append any remaining nodes from the non-exhausted input list to the result list
    if list1 != nil {
        tail.Next = list1
    } else if list2 != nil {
        tail.Next = list2
    }
    // Return the head of the result list (after the dummy node)
    return dummy.Next
}
