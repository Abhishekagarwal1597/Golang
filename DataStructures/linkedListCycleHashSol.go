

func detectCycle(head *ListNode) *ListNode {
    // Step 1: Create an empty hash table to store visited nodes
    visited := make(map[*ListNode]bool)
    
    // Step 2: Set current node to head of linked list
    curr := head
    
    // Step 3: Loop through linked list until end is reached or cycle is found
    for curr != nil {
        // If current node is already in hash table, there is a cycle
        if visited[curr] {
            return curr
        }
        
        // Mark current node as visited and move to next node
        visited[curr] = true
        curr = curr.Next
    }
    
    // Step 5: If we reach the end of the linked list, there is no cycle
    return nil
}
