# here u need to maintain two node, prev and next to keep track of the previous and next node of the linked list.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
// Define a function to reverse a singly-linked list
func reverseList(head *ListNode) *ListNode {
   // Declare two pointers to hold previous and next nodes in the list 
   var prev *ListNode = nil   // Set prev to nil to start with an empty list
    var next *ListNode 
   
  // Check if the input list is empty
    if head == nil {
      // If so, return the empty list
        return head
    } else{
      
      // Otherwise, loop through the list
        for head != nil {
            // Set next to the current node's next node. Initialized next variable to store the next node in the list before modifying the head.Next
            next = head.Next
            // Set the current node's next node to the previous node
            head.Next = prev
            // Set prev to the current node
            prev = head
            // Set head to the next node
            head = next
        }
      //return the reverse list
   return prev 
    }
}
